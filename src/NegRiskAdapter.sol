// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {ERC1155TokenReceiver} from "lib/solmate/src/tokens/ERC1155.sol";
import {ERC20} from "lib/solmate/src/tokens/ERC20.sol";
import {SafeTransferLib} from "lib/solmate/src/utils/SafeTransferLib.sol";
import {WrappedCollateral, ISeasonScopedCollateral, ISeasonScopedTransfer} from "./WrappedCollateral.sol";
import {MarketData, MarketStateManager, IMarketStateManagerEE} from "./modules/MarketDataManager.sol";
import {CTHelpers} from "./libraries/CTHelpers.sol";
import {Helpers} from "./libraries/Helpers.sol";
import {NegRiskIdLib} from "./libraries/NegRiskIdLib.sol";
import {IConditionalTokens} from "./interfaces/IConditionalTokens.sol";
import {Auth} from "./modules/Auth.sol";
import {IAuthEE} from "./modules/interfaces/IAuth.sol";

/// @title INegRiskAdapterEE
/// @notice NegRiskAdapter Errors and Events
interface INegRiskAdapterEE is IMarketStateManagerEE, IAuthEE {
    error InvalidIndexSet();
    error LengthMismatch();
    error UnexpectedCollateralToken();
    error NoConvertiblePositions();
    error NotApprovedForAll();
    /// @dev A conditionId that this adapter never prepared was used on a season-scoped
    ///      deployment. Without the question behind it the market — and so the season
    ///      backing it — cannot be resolved, and paying out would have to guess.
    error UnknownCondition(bytes32 conditionId);
    /// @dev The underlying reported a season id of zero, which {marketSessionId} uses to
    ///      mean "not pinned yet". Accepting it would leave the market re-pinning on every
    ///      intake and paying out of whichever season happened to be live.
    error InvalidCollateralSessionId();
    /// @dev A conversion tried to pay out of a market that has never taken collateral, so
    ///      there is no season to draw the released collateral from.
    error MarketNotFunded(bytes32 marketId);
    /// @dev Collateral was offered to a market whose season has closed. Taking it in would
    ///      escrow live-season units under a pin the payout path can never hand back.
    error MarketSeasonClosed(bytes32 marketId, uint256 pinnedSessionId, uint256 liveSessionId);

    event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data);
    event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data);
    event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome);
    event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount);
    event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount);
    event PositionsConverted(
        address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount
    );
    event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout);
    event BulkOutcomeReported(bytes32 indexed marketId, uint256 winnerIndex, uint256 questionCount);
}

/// @title NegRiskAdapter
/// @notice Adapter for the CTF enabling the linking of a set binary markets where only one can resolve true
/// @notice The adapter prevents more than one question in the same multi-outcome market from resolving true
/// @notice And the adapter allows for the conversion of a set of no positions, to collateral plus the set of
/// complementary yes positions
/// @author Mike Shrieve (mike@polymarket.com)
contract NegRiskAdapter is ERC1155TokenReceiver, MarketStateManager, INegRiskAdapterEE, Auth {
    using SafeTransferLib for ERC20;

    /*//////////////////////////////////////////////////////////////
                                 STATE
    //////////////////////////////////////////////////////////////*/

    IConditionalTokens public immutable ctf;
    ERC20 public immutable col;
    WrappedCollateral public immutable wcol;
    address public immutable vault;

    address public constant NO_TOKEN_BURN_ADDRESS = address(bytes20(bytes32(keccak256("NO_TOKEN_BURN_ADDRESS"))));
    uint256 public constant FEE_DENOMINATOR = 10_000;

    /// @notice Whether the collateral keeps balances per season.
    /// @dev Probed once, at construction, and never re-probed — see {WrappedCollateral}
    ///      for why a mid-life flip would split the wrapper across both modes.
    bool public immutable seasonScoped;

    /// @notice The question behind each condition this adapter prepared.
    /// @dev `conditionId` is a hash of the question, so it cannot be walked backwards.
    ///      Recording the pair at preparation time is what lets {splitPosition} — which is
    ///      handed only a conditionId — find the market, and through it the season backing
    ///      that market. Written in both modes; only read in season-scoped mode.
    mapping(bytes32 conditionId => bytes32 questionId) public conditionQuestionId;

    /// @notice The season each market's collateral is escrowed under; `0` means not pinned.
    /// @dev Pinned at the market's first collateral intake and never moved, mirroring the
    ///      CTF's own `conditionSessionId`. Pinning per market rather than per question is
    ///      what makes {convertPositions} sound: conversion pays out of collateral pooled
    ///      across every question in the market, so those questions have to agree on a
    ///      season by construction rather than by check.
    mapping(bytes32 marketId => uint256 seasonId) public marketSessionId;

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    /// @param _ctf        - ConditionalTokens address
    /// @param _collateral - collateral address
    constructor(address _ctf, address _collateral, address _vault) {
        ctf = IConditionalTokens(_ctf);
        col = ERC20(_collateral);
        vault = _vault;

        wcol = new WrappedCollateral(_collateral, col.decimals());
        // approve the ctf to transfer wcol on our behalf
        wcol.approve(_ctf, type(uint256).max);

        // The wrapper probes the collateral on the way up; reuse its answer rather than
        // probing twice and risking the two disagreeing.
        bool isSeasonScoped = wcol.seasonScoped();
        seasonScoped = isSeasonScoped;

        if (isSeasonScoped) {
            // Season-scoped collateral leaves this contract through ERC-1155
            // `safeTransferFrom`, which consults operator approval and never the ERC-20
            // allowance. It HAS to be an operator approval: an allowance on such a token is
            // recorded against the live season alone, so the `approve` below would fall to
            // zero at the first season boundary and every wrap after that would revert —
            // permanently, since this contract is immutable and exposes no way to renew it.
            ISeasonScopedTransfer(_collateral).setApprovalForAll(address(wcol), true);
        } else {
            // approve wcol to transfer collateral on our behalf
            col.approve(address(wcol), type(uint256).max);
        }
    }

    /*//////////////////////////////////////////////////////////////
                                 SEASONS
    //////////////////////////////////////////////////////////////*/

    /// @notice The season a condition's collateral is escrowed under; `0` when the market
    ///         has taken no collateral yet, and always `0` on a plain-ERC20 deployment.
    function conditionSessionId(bytes32 _conditionId) public view returns (uint256) {
        bytes32 questionId = conditionQuestionId[_conditionId];
        if (questionId == bytes32(0)) return 0;
        return marketSessionId[NegRiskIdLib.getMarketId(questionId)];
    }

    /// @dev Resolves the season to escrow under for `_conditionId`, pinning the market to
    ///      the live season on its first intake. Every later intake, and every payout,
    ///      obeys the record — re-reading the live season instead would let a market take
    ///      in one season and pay out of another.
    function _pinSession(bytes32 _conditionId) internal returns (uint256) {
        bytes32 questionId = conditionQuestionId[_conditionId];
        if (questionId == bytes32(0)) revert UnknownCondition(_conditionId);

        bytes32 marketId = NegRiskIdLib.getMarketId(questionId);
        uint256 sessionId = marketSessionId[marketId];
        if (sessionId != 0) return sessionId;

        sessionId = ISeasonScopedCollateral(address(col)).currentSessionId();
        if (sessionId == 0) revert InvalidCollateralSessionId();

        marketSessionId[marketId] = sessionId;
        return sessionId;
    }

    /// @dev The season a payout must draw on. Unlike {_pinSession} this never pins: a
    ///      market that never took collateral has nothing to pay out, and silently pinning
    ///      it here would let a later intake land in a season the payout already used.
    function _payoutSession(bytes32 _conditionId) internal view returns (uint256) {
        uint256 sessionId = conditionSessionId(_conditionId);
        if (sessionId == 0) revert UnknownCondition(_conditionId);
        return sessionId;
    }

    /*//////////////////////////////////////////////////////////////
                                  IDS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the conditionId for a given questionId
    /// @param _questionId  - the questionId
    /// @return conditionId - the corresponding conditionId
    function getConditionId(bytes32 _questionId) public view returns (bytes32) {
        return CTHelpers.getConditionId(
            address(this), // oracle
            _questionId,
            2 // outcomeCount
        );
    }

    /// @notice Returns the positionId for a given questionId and outcome
    /// @param _questionId  - the questionId
    /// @param _outcome     - the boolean outcome
    /// @return positionId  - the corresponding positionId
    function getPositionId(bytes32 _questionId, bool _outcome) public view returns (uint256) {
        bytes32 collectionId = CTHelpers.getCollectionId(
            bytes32(0),
            getConditionId(_questionId),
            _outcome ? 1 : 2 // 1 (0b01) is yes, 2 (0b10) is no
        );

        uint256 positionId = CTHelpers.getPositionId(address(wcol), collectionId);
        return positionId;
    }

    /*//////////////////////////////////////////////////////////////
                             SPLIT POSITION
    //////////////////////////////////////////////////////////////*/

    /// @notice Splits collateral to a complete set of conditional tokens for a single question
    /// @notice This function signature is the same as the CTF's splitPosition
    /// @param _collateralToken - the collateral token, must be the same as the adapter's collateral token
    /// @param _conditionId - the conditionId for the question
    /// @param _amount - the amount of collateral to split
    function splitPosition(address _collateralToken, bytes32, bytes32 _conditionId, uint256[] calldata, uint256 _amount)
        external
    {
        if (_collateralToken != address(col)) revert UnexpectedCollateralToken();
        splitPosition(_conditionId, _amount);
    }

    /// @notice Splits collateral to a complete set of conditional tokens for a single question
    /// @param _conditionId - the conditionId for the question
    /// @param _amount      - the amount of collateral to split
    function splitPosition(bytes32 _conditionId, uint256 _amount) public {
        if (seasonScoped) {
            uint256 sessionId = _pinSession(_conditionId);

            // A market can only ever be funded while its own season is live, so the
            // intake only ever needs the live season and the ordinary ERC-20 pull below
            // reaches it. Checked explicitly rather than left to fail on a balance,
            // because the facade would otherwise take LIVE-season units into a market
            // pinned to an older one — escrowing collateral the payout path, bound to the
            // pin, could never hand back.
            //
            // Pulling through the facade rather than by naming the season is deliberate:
            // an ERC-20 allowance is all the funder has to grant, and because such an
            // allowance is itself season-keyed it expires on its own at the boundary. An
            // operator approval would have to span every season to be usable at all, and
            // would then outlive both the market and any revocation of this adapter.
            uint256 liveSessionId = ISeasonScopedCollateral(address(col)).currentSessionId();
            if (sessionId != liveSessionId) {
                revert MarketSeasonClosed(
                    NegRiskIdLib.getMarketId(conditionQuestionId[_conditionId]), sessionId, liveSessionId
                );
            }

            col.safeTransferFrom(msg.sender, address(this), _amount);
            wcol.wrapSeason(address(this), _amount, sessionId);
        } else {
            col.safeTransferFrom(msg.sender, address(this), _amount);
            wcol.wrap(address(this), _amount);
        }

        ctf.splitPosition(address(wcol), bytes32(0), _conditionId, Helpers.partition(), _amount);
        ctf.safeBatchTransferFrom(
            address(this), msg.sender, Helpers.positionIds(address(wcol), _conditionId), Helpers.values(2, _amount), ""
        );

        emit PositionSplit(msg.sender, _conditionId, _amount);
    }

    /*//////////////////////////////////////////////////////////////
                            MERGE POSITIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Merges a complete set of conditional tokens for a single question to collateral
    /// @notice This function signature is the same as the CTF's mergePositions
    /// @param _collateralToken - the collateral token, must be the same as the adapter's collateral token
    /// @param _conditionId     - the conditionId for the question
    /// @param _amount          - the amount of collateral to merge
    function mergePositions(
        address _collateralToken,
        bytes32,
        bytes32 _conditionId,
        uint256[] calldata,
        uint256 _amount
    ) external {
        if (_collateralToken != address(col)) revert UnexpectedCollateralToken();
        mergePositions(_conditionId, _amount);
    }

    /// @notice Merges a complete set of conditional tokens for a single question to collateral
    /// @param _conditionId - the conditionId for the question
    /// @param _amount      - the amount of collateral to merge
    function mergePositions(bytes32 _conditionId, uint256 _amount) public {
        uint256[] memory positionIds = Helpers.positionIds(address(wcol), _conditionId);

        // get conditional tokens from sender
        ctf.safeBatchTransferFrom(msg.sender, address(this), positionIds, Helpers.values(2, _amount), "");
        ctf.mergePositions(address(wcol), bytes32(0), _conditionId, Helpers.partition(), _amount);

        // Pays out of the season this market was pinned to, not the live one, so a market
        // that outlived its season still hands back exactly what it took in.
        if (seasonScoped) {
            wcol.unwrapSeason(msg.sender, _amount, _payoutSession(_conditionId));
        } else {
            wcol.unwrap(msg.sender, _amount);
        }

        emit PositionsMerge(msg.sender, _conditionId, _amount);
    }

    /*//////////////////////////////////////////////////////////////
                           ERC1155 OPERATIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Proxies ERC1155 balanceOf to the CTF
    /// @param _owner   - the owner of the tokens
    /// @param _id      - the positionId
    /// @return balance - the owner's balance
    function balanceOf(address _owner, uint256 _id) external view returns (uint256) {
        return ctf.balanceOf(_owner, _id);
    }

    /// @notice Proxies ERC1155 balanceOfBatch to the CTF
    /// @param _owners   - the owners of the tokens
    /// @param _ids      - the positionIds
    /// @return balances - the owners' balances
    function balanceOfBatch(address[] memory _owners, uint256[] memory _ids) external view returns (uint256[] memory) {
        return ctf.balanceOfBatch(_owners, _ids);
    }

    /// @notice Proxies ERC1155 safeTransferFrom to the CTF
    /// @notice Can only be called by an admin
    /// @notice Requires this contract to be approved for all
    /// @notice Requires the sender to be approved for all
    /// @param _from  - the owner of the tokens
    /// @param _to    - the recipient of the tokens
    /// @param _id    - the positionId
    /// @param _value - the amount of tokens to transfer
    /// @param _data  - the data to pass to the recipient
    function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes calldata _data)
        external
        onlyAdmin
    {
        if (!ctf.isApprovedForAll(_from, msg.sender)) {
            revert NotApprovedForAll();
        }

        return ctf.safeTransferFrom(_from, _to, _id, _value, _data);
    }

    /*//////////////////////////////////////////////////////////////
                            REDEEM POSITION
    //////////////////////////////////////////////////////////////*/

    /// @notice Redeem a set of conditional tokens for collateral
    /// @param _conditionId - conditionId of the conditional tokens to redeem
    /// @param _amounts     - amounts of conditional tokens to redeem
    /// _amounts should always have length 2, with the first element being the amount of yes tokens to redeem and the
    /// second element being the amount of no tokens to redeem
    function redeemPositions(bytes32 _conditionId, uint256[] calldata _amounts) public {
        uint256[] memory positionIds = Helpers.positionIds(address(wcol), _conditionId);

        // get conditional tokens from sender
        ctf.safeBatchTransferFrom(msg.sender, address(this), positionIds, _amounts, "");
        ctf.redeemPositions(address(wcol), bytes32(0), _conditionId, Helpers.partition());

        uint256 payout = wcol.balanceOf(address(this));
        if (payout > 0) {
            // Redemption is the path a market pinned to a closed season still has to
            // work on — the whole reason the season is recorded rather than re-read.
            if (seasonScoped) {
                wcol.unwrapSeason(msg.sender, payout, _payoutSession(_conditionId));
            } else {
                wcol.unwrap(msg.sender, payout);
            }
        }

        emit PayoutRedemption(msg.sender, _conditionId, _amounts, payout);
    }

    /*//////////////////////////////////////////////////////////////
                            CONVERT POSITIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Convert a set of no positions to the complementary set of yes positions plus collateral proportional to
    /// (# of no positions - 1)
    /// @notice If the market has a fee, the fee is taken from both collateral and the yes positions
    /// @param _marketId - the marketId
    /// @param _indexSet - the set of positions to convert, expressed as an index set where the least significant bit is
    /// the first question (index zero)
    /// @param _amount   - the amount of tokens to convert
    function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) external {
        MarketData md = getMarketData(_marketId);
        uint256 questionCount = md.questionCount();

        if (md.oracle() == address(0)) revert MarketNotPrepared();
        if (questionCount <= 1) revert NoConvertiblePositions();
        if (_indexSet == 0) revert InvalidIndexSet();
        if ((_indexSet >> questionCount) > 0) revert InvalidIndexSet();

        // if _amount is 0, return early
        if (_amount == 0) {
            return;
        }

        uint256 index = 0;
        uint256 noPositionCount;

        // count number of no positions
        while (index < questionCount) {
            unchecked {
                if ((_indexSet & (1 << index)) > 0) {
                    ++noPositionCount;
                }
                ++index;
            }
        }

        uint256 yesPositionCount = questionCount - noPositionCount;
        uint256[] memory noPositionIds = new uint256[](noPositionCount);
        uint256[] memory yesPositionIds = new uint256[](yesPositionCount);
        uint256[] memory accumulatedNoPositionIds = new uint256[](yesPositionCount);

        // mint the amount of wcol required
        wcol.mint(yesPositionCount * _amount);

        // populate noPositionIds and yesPositionIds
        // split yes positions
        {
            uint256 noIndex;
            uint256 yesIndex;
            index = 0;

            while (index < questionCount) {
                bytes32 questionId = NegRiskIdLib.getQuestionId(_marketId, uint8(index));

                if ((_indexSet & (1 << index)) > 0) {
                    // NO
                    noPositionIds[noIndex] = getPositionId(questionId, false);

                    unchecked {
                        ++noIndex;
                    }
                } else {
                    // YES
                    yesPositionIds[yesIndex] = getPositionId(questionId, true);
                    accumulatedNoPositionIds[yesIndex] = getPositionId(questionId, false);

                    // split position to get yes and no tokens
                    // the no tokens will be discarded
                    _splitPosition(getConditionId(questionId), _amount);

                    unchecked {
                        ++yesIndex;
                    }
                }
                unchecked {
                    ++index;
                }
            }
        }

        // transfer the caller's no tokens _and_ accumulated no tokens to the burn address
        // these must never be redeemed
        {
            ctf.safeBatchTransferFrom(
                msg.sender, NO_TOKEN_BURN_ADDRESS, noPositionIds, Helpers.values(noPositionIds.length, _amount), ""
            );
            ctf.safeBatchTransferFrom(
                address(this),
                NO_TOKEN_BURN_ADDRESS,
                accumulatedNoPositionIds,
                Helpers.values(yesPositionCount, _amount),
                ""
            );
        }

        uint256 feeAmount = (_amount * md.feeBips()) / FEE_DENOMINATOR;
        uint256 amountOut = _amount - feeAmount;

        if (noPositionIds.length > 1) {
            // collateral out is always proportional to the number of no positions minus 1
            uint256 multiplier = noPositionIds.length - 1;
            // transfer collateral fees to vault
            _releaseConverted(_marketId, vault, multiplier * feeAmount);
            // transfer collateral to sender
            _releaseConverted(_marketId, msg.sender, multiplier * amountOut);
        }

        if (yesPositionIds.length > 0) {
            if (feeAmount > 0) {
                // transfer yes token fees to vault
                ctf.safeBatchTransferFrom(
                    address(this), vault, yesPositionIds, Helpers.values(yesPositionIds.length, feeAmount), ""
                );
            }

            // transfer yes tokens to sender
            ctf.safeBatchTransferFrom(
                address(this), msg.sender, yesPositionIds, Helpers.values(yesPositionIds.length, amountOut), ""
            );
        }

        emit PositionsConverted(msg.sender, _marketId, _indexSet, _amount);
    }

    /*//////////////////////////////////////////////////////////////
                             PREPARE MARKET
    //////////////////////////////////////////////////////////////*/

    /// @notice Prepare a multi-outcome market
    /// @param _feeBips  - the fee for the market, out of 10_000
    /// @param _metadata     - metadata for the market
    /// @return marketId - the marketId
    function prepareMarket(uint256 _feeBips, bytes calldata _metadata) external returns (bytes32) {
        bytes32 marketId = _prepareMarket(_feeBips, _metadata);

        emit MarketPrepared(marketId, msg.sender, _feeBips, _metadata);

        return marketId;
    }

    /*//////////////////////////////////////////////////////////////
                            PREPARE QUESTION
    //////////////////////////////////////////////////////////////*/

    /// @notice Prepare a question for a given market
    /// @param _marketId   - the id of the market for which to prepare the question
    /// @param _metadata   - the question metadata
    /// @return questionId - the id of the resulting question
    function prepareQuestion(bytes32 _marketId, bytes calldata _metadata) external returns (bytes32) {
        (bytes32 questionId, uint256 questionIndex) = _prepareQuestion(_marketId);
        bytes32 conditionId = getConditionId(questionId);

        // Recorded unconditionally, even when the condition already exists on the CTF:
        // this pair is what lets a later `splitPosition` — handed only a conditionId —
        // find the market whose season backs it.
        conditionQuestionId[conditionId] = questionId;

        // check to see if the condition has already been prepared on the ctf
        if (ctf.getOutcomeSlotCount(conditionId) == 0) {
            ctf.prepareCondition(address(this), questionId, 2);
        }

        emit QuestionPrepared(_marketId, questionId, questionIndex, _metadata);

        return questionId;
    }

    /*//////////////////////////////////////////////////////////////
                             REPORT OUTCOME
    //////////////////////////////////////////////////////////////*/

    /// @notice Report the outcome of a question
    /// @param _questionId - the questionId to report
    /// @param _outcome    - the outcome of the question
    function reportOutcome(bytes32 _questionId, bool _outcome) external {
        _reportOutcome(_questionId, _outcome);

        ctf.reportPayouts(_questionId, Helpers.payouts(_outcome));

        emit OutcomeReported(NegRiskIdLib.getMarketId(_questionId), _questionId, _outcome);
    }

    /*//////////////////////////////////////////////////////////////
                          BULK REPORT OUTCOME
    //////////////////////////////////////////////////////////////*/

    /// @notice Resolve all questions in a neg-risk market atomically
    /// @dev Exactly one question (at _winnerIndex) resolves true, all others false
    /// @param _marketId The neg-risk market ID
    /// @param _winnerIndex The question index that resolves true (0-based)
    function bulkReportOutcome(bytes32 _marketId, uint256 _winnerIndex) external {
        MarketData md = getMarketData(_marketId);
        address oracle = md.oracle();
        uint256 questionCount = md.questionCount();

        if (oracle == address(0)) revert MarketNotPrepared();
        if (oracle != msg.sender) revert OnlyOracle();
        if (questionCount < 2) revert NoConvertiblePositions();
        if (_winnerIndex >= questionCount) revert IndexOutOfBounds();
        if (md.determined()) revert MarketAlreadyDetermined();

        for (uint256 i = 0; i < questionCount;) {
            bytes32 questionId = NegRiskIdLib.getQuestionId(_marketId, uint8(i));
            bool outcome = (i == _winnerIndex);

            _reportOutcome(questionId, outcome);
            ctf.reportPayouts(questionId, Helpers.payouts(outcome));

            emit OutcomeReported(_marketId, questionId, outcome);

            unchecked {
                ++i;
            }
        }

        emit BulkOutcomeReported(_marketId, _winnerIndex, questionCount);
    }

    /*//////////////////////////////////////////////////////////////
                                INTERNAL
    //////////////////////////////////////////////////////////////*/

    /// @dev internal function to avoid stack too deep in convertPositions
    function _splitPosition(bytes32 _conditionId, uint256 _amount) internal {
        ctf.splitPosition(address(wcol), bytes32(0), _conditionId, Helpers.partition(), _amount);
    }

    /// @dev Releases the collateral a conversion frees up, out of the season backing the
    ///      market. Conversion pools collateral across every question in the market, which
    ///      is exactly why the season is pinned per market and not per question: there is
    ///      one season to name here, by construction rather than by check.
    ///
    ///      Also internal to avoid stack too deep in {convertPositions}, and it absorbs the
    ///      zero-amount guard the two call sites would otherwise each need.
    function _releaseConverted(bytes32 _marketId, address _to, uint256 _amount) internal {
        if (_amount == 0) return;

        if (seasonScoped) {
            uint256 sessionId = marketSessionId[_marketId];
            if (sessionId == 0) revert MarketNotFunded(_marketId);
            wcol.releaseSeason(_to, _amount, sessionId);
        } else {
            wcol.release(_to, _amount);
        }
    }
}
