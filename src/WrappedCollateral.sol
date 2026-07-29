// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {ERC20} from "lib/solmate/src/tokens/ERC20.sol";
import {ERC1155TokenReceiver} from "lib/solmate/src/tokens/ERC1155.sol";
import {SafeTransferLib} from "lib/solmate/src/utils/SafeTransferLib.sol";

/// @title IWrappedCollateralEE
/// @notice WrappedCollateral Errors and Events
interface IWrappedCollateralEE {
    error OnlyOwner();
    /// @dev A plain-ERC20 entrypoint was used on a wrapper whose underlying is
    ///      season-scoped. Those entrypoints move the underlying through its ERC-20
    ///      facade, which addresses only the live season, so they would pay out of
    ///      whichever season happens to be live rather than the one that backs the
    ///      market being exited. Use the `*Season` entrypoints instead.
    error SeasonScopedCollateral();
    /// @dev A `*Season` entrypoint was used on a wrapper whose underlying is a plain
    ///      ERC20 and has no seasons to name.
    error NotSeasonScopedCollateral();
    /// @dev The named season holds less than the caller is trying to take out. Guards
    ///      the invariant that a market can only ever be paid out of the season that
    ///      backs it, so a wrong season id fails loudly instead of draining another
    ///      market's backing.
    error InsufficientSeasonEscrow(uint256 seasonId, uint256 requested, uint256 available);
    /// @dev {unwrapSeasonFor} was called by someone other than the owner or the owner's
    ///      settlement contract. Both of those only ever unwrap against positions they
    ///      have just settled for that condition; anyone else naming a condition would be
    ///      choosing which market's backing to draw on.
    error NotSettlementCaller();
    /// @dev The named condition has no season behind it — its market never took
    ///      collateral, so there is nothing to unwrap out of.
    error ConditionNotFunded(bytes32 conditionId);

    /// @notice `amount` of the underlying entered the wrapper under `seasonId`.
    event SeasonEscrowCredited(uint256 indexed seasonId, uint256 amount);
    /// @notice `amount` of the underlying left the wrapper from `seasonId`.
    event SeasonEscrowDebited(uint256 indexed seasonId, uint256 amount);
}

/// @title ISeasonScopedCollateral
/// @notice Capability interface for a season-scoped ERC-1155 collateral token: balances
///         are kept per season id and one season at a time is "live".
/// @dev Structurally identical to `IPunchPoint` in the prediction-contracts repo, and
///      deliberately so — `type(...).interfaceId` is the XOR of the selectors declared in
///      the interface itself, excluding inherited ones, so both types resolve to the same
///      id (`0x473e92af`). Declared here rather than imported because this repository does
///      not depend on prediction-contracts.
interface ISeasonScopedCollateral {
    /// @notice Contract whose `SENDER_ROLE` gates transfers of this token.
    function ctf() external view returns (address);

    /// @notice The season that the plain ERC-20 facade currently addresses.
    function currentSessionId() external view returns (uint256);
}

/// @title ISeasonScopedTransfer
/// @notice The subset of ERC-1155 this wrapper needs to move a named season.
/// @title ISeasonPinRegistry
/// @notice The part of the owning adapter this wrapper has to consult: which season backs
///         a given condition, and which settlement contract is entitled to say so.
/// @dev Read at call time rather than cached at construction. The wrapper is created from
///      inside the adapter's constructor, so the adapter has no deployed code yet and
///      neither value can be read back then.
interface ISeasonPinRegistry {
    function conditionSessionId(bytes32 conditionId) external view returns (uint256);
    function ctf() external view returns (address);
}

interface ISeasonScopedTransfer {
    function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external;
    function balanceOf(address account, uint256 id) external view returns (uint256);
    function setApprovalForAll(address operator, bool approved) external;
}

string constant NAME = "Wrapped Collateral";
string constant SYMBOL = "WCOL";

/// @title WrappedCollateral
/// @author Mike Shrieve (mike@polymarket.com)
/// @notice Wraps an ERC20 token to be used as collateral in the CTF
/// @dev Two modes, decided once at construction by probing the underlying for
///      {ISeasonScopedCollateral} through ERC-165:
///
///        plain          — the original behaviour, untouched. `wrap`/`unwrap`/`release`
///                         move the underlying with ordinary ERC-20 transfers.
///        season-scoped  — the underlying keeps balances per season id and its ERC-20
///                         facade only ever addresses the live one. Wrapping and
///                         unwrapping must therefore name a season, and the plain
///                         entrypoints are disabled.
///
///      The season-scoped mode exists because WCOL is fungible and seasons are not. A
///      wrapper serving two markets pinned to different seasons holds both seasons at
///      once; unwrapping through the ERC-20 facade would hand the exiting market
///      whichever season is live, drawn out of the other market's backing. Naming the
///      season on the way out — and checking it against a per-season ledger — is what
///      keeps each market's backing its own.
///
///      Which season backs which market is not knowable here: this contract sees amounts,
///      not conditions. The owner (the adapter) holds that mapping and is the only caller
///      allowed to name a season on the way out.
contract WrappedCollateral is IWrappedCollateralEE, ERC20, ERC1155TokenReceiver {
    using SafeTransferLib for ERC20;

    /*//////////////////////////////////////////////////////////////
                                 STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable owner;
    address public immutable underlying;

    /// @notice Whether the underlying keeps balances per season.
    /// @dev Probed once, at construction, and never re-probed. An underlying that changed
    ///      its answer mid-life would otherwise split this wrapper across both modes: part
    ///      of its holdings entered under a season and tracked, part moved as plain ERC-20
    ///      and untracked, with no way to reconcile the two on the way out.
    bool public immutable seasonScoped;

    /// @notice How much of the underlying this wrapper holds under each season.
    /// @dev Credited on {wrapSeason}, debited on {unwrapSeason} and {releaseSeason}. Only
    ///      populated in season-scoped mode.
    mapping(uint256 seasonId => uint256 amount) public seasonEscrow;

    /*//////////////////////////////////////////////////////////////
                               MODIFIERS
    //////////////////////////////////////////////////////////////*/

    modifier onlyOwner() {
        if (msg.sender != owner) revert OnlyOwner();
        _;
    }

    /// @dev Guards the entrypoints that move the underlying as a plain ERC-20.
    modifier onlyPlain() {
        if (seasonScoped) revert SeasonScopedCollateral();
        _;
    }

    /// @dev Guards the entrypoints that name a season.
    modifier onlySeasonScoped() {
        if (!seasonScoped) revert NotSeasonScopedCollateral();
        _;
    }

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    /// @param _underlying The address of the underlying ERC20 token
    /// @param _decimals The number of decimals of the underlying ERC20 token
    constructor(address _underlying, uint8 _decimals) ERC20(NAME, SYMBOL, _decimals) {
        owner = msg.sender;
        underlying = _underlying;
        seasonScoped = _probeSeasonScoped(_underlying);
    }

    /*//////////////////////////////////////////////////////////////
                                 UNWRAP
    //////////////////////////////////////////////////////////////*/

    /// @notice Unwraps the specified amount of tokens
    /// @param _to The address to send the unwrapped tokens to
    /// @param _amount The amount of tokens to unwrap
    function unwrap(address _to, uint256 _amount) external onlyPlain {
        _burn(msg.sender, _amount);
        ERC20(underlying).safeTransfer(_to, _amount);
    }

    /// @notice Unwraps `_amount` and pays it out of `_seasonId`.
    /// @dev Owner-only, unlike {unwrap}: the season a given amount of WCOL is backed by is
    ///      a property of the market it came from, which only the owner knows. Letting any
    ///      holder name a season would reintroduce exactly the cross-market draw this mode
    ///      exists to prevent.
    /// @param _to The address to send the unwrapped tokens to
    /// @param _amount The amount of tokens to unwrap
    /// @param _seasonId The season backing this amount
    function unwrapSeason(address _to, uint256 _amount, uint256 _seasonId) external onlyOwner onlySeasonScoped {
        _burn(msg.sender, _amount);
        _payOutOfSeason(_to, _amount, _seasonId);
    }

    /// @notice Unwraps `_amount` against `_conditionId`, paying out of the season backing
    ///         that condition's market.
    /// @dev The entrypoint for a settlement contract that redeems positions itself and so
    ///      never learns the season: it knows the condition, and the owner maps that to a
    ///      season. The CTF's `batchRedeem` is the caller this exists for — it burns the
    ///      redeemer's positions and unwraps the proceeds in one pass, with no adapter
    ///      call in between to name the season.
    ///
    ///      Restricted to the owner and the owner's settlement contract, and NOT open the
    ///      way {unwrap} is. Both of those only reach here holding WCOL they just settled
    ///      for this very condition. An arbitrary holder naming a condition would be
    ///      picking which market's backing to draw on — the wrapper is fungible, the
    ///      seasons behind it are not, and that choice is exactly what the season ledger
    ///      exists to take away.
    /// @param _conditionId The condition whose market's season backs this amount.
    /// @param _to The address to send the unwrapped tokens to.
    /// @param _amount The amount of tokens to unwrap.
    function unwrapSeasonFor(bytes32 _conditionId, address _to, uint256 _amount) external onlySeasonScoped {
        ISeasonPinRegistry registry = ISeasonPinRegistry(owner);
        if (msg.sender != owner && msg.sender != registry.ctf()) revert NotSettlementCaller();

        uint256 seasonId = registry.conditionSessionId(_conditionId);
        if (seasonId == 0) revert ConditionNotFunded(_conditionId);

        _burn(msg.sender, _amount);
        _payOutOfSeason(_to, _amount, seasonId);
    }

    /*//////////////////////////////////////////////////////////////
                                 ADMIN
    //////////////////////////////////////////////////////////////*/

    /// @notice Wraps the specified amount of tokens
    /// @notice Can only be called by the owner
    /// @param _to     - the address to send the wrapped tokens to
    /// @param _amount - the amount of tokens to wrap
    function wrap(address _to, uint256 _amount) external onlyOwner onlyPlain {
        ERC20(underlying).safeTransferFrom(msg.sender, address(this), _amount);
        _mint(_to, _amount);
    }

    /// @notice Wraps `_amount` of season `_seasonId`, taken from the owner.
    /// @dev The pull is an ERC-1155 `safeTransferFrom`, which consults operator approval
    ///      and never the ERC-20 allowance — the owner has to have granted this contract
    ///      `setApprovalForAll` on the underlying.
    /// @param _to       - the address to send the wrapped tokens to
    /// @param _amount   - the amount of tokens to wrap
    /// @param _seasonId - the season to take in
    function wrapSeason(address _to, uint256 _amount, uint256 _seasonId) external onlyOwner onlySeasonScoped {
        ISeasonScopedTransfer(underlying).safeTransferFrom(msg.sender, address(this), _seasonId, _amount, "");

        seasonEscrow[_seasonId] += _amount;
        emit SeasonEscrowCredited(_seasonId, _amount);

        _mint(_to, _amount);
    }

    /// @notice Burns the specified amount of tokens
    /// @notice Can only be called by the owner
    /// @param _amount - the amount of tokens to burn
    function burn(uint256 _amount) external onlyOwner {
        _burn(msg.sender, _amount);
    }

    /// @notice Mints the specified amount of tokens
    /// @notice Can only be called by the owner
    /// @param _amount - the amount of tokens to mint
    /// @dev Mints unbacked WCOL — no underlying moves — so the season ledger is untouched
    ///      in both modes. The neg-risk conversion math is what keeps supply and holdings
    ///      reconciled.
    function mint(uint256 _amount) external onlyOwner {
        _mint(msg.sender, _amount);
    }

    /// @notice Releases the specified amount of the underlying token
    /// @notice Can only be called by the owner
    /// @param _to     - the address to send the released tokens to
    /// @param _amount - the amount of tokens to release
    function release(address _to, uint256 _amount) external onlyOwner onlyPlain {
        ERC20(underlying).safeTransfer(_to, _amount);
    }

    /// @notice Releases `_amount` of the underlying out of `_seasonId`, without burning WCOL.
    /// @param _to       - the address to send the released tokens to
    /// @param _amount   - the amount of tokens to release
    /// @param _seasonId - the season to draw on
    function releaseSeason(address _to, uint256 _amount, uint256 _seasonId) external onlyOwner onlySeasonScoped {
        _payOutOfSeason(_to, _amount, _seasonId);
    }

    /*//////////////////////////////////////////////////////////////
                                INTERNAL
    //////////////////////////////////////////////////////////////*/

    /// @dev Debits the season ledger and sends the named season out. Shared by
    ///      {unwrapSeason} and {releaseSeason}, which differ only in whether WCOL burns.
    function _payOutOfSeason(address _to, uint256 _amount, uint256 _seasonId) internal {
        uint256 escrowed = seasonEscrow[_seasonId];
        if (escrowed < _amount) revert InsufficientSeasonEscrow(_seasonId, _amount, escrowed);

        unchecked {
            seasonEscrow[_seasonId] = escrowed - _amount;
        }
        emit SeasonEscrowDebited(_seasonId, _amount);

        ISeasonScopedTransfer(underlying).safeTransferFrom(address(this), _to, _seasonId, _amount, "");
    }

    /// @dev True when `_underlying` advertises {ISeasonScopedCollateral} through ERC-165.
    ///      Raw staticcall rather than `try`: an underlying whose fallback returns nothing
    ///      (or a non-canonical bool) would make return-data decoding revert *outside* a
    ///      try/catch, so probing must never depend on decoding succeeding. A plain ERC20
    ///      either reverts or returns nothing here, and reads as false.
    function _probeSeasonScoped(address _underlying) internal view returns (bool) {
        (bool ok, bytes memory data) = _underlying.staticcall(
            abi.encodeWithSignature("supportsInterface(bytes4)", type(ISeasonScopedCollateral).interfaceId)
        );

        // Byte 31 is the low byte of the first returned word — non-zero means `true`.
        return ok && data.length >= 32 && data[31] != 0;
    }
}
