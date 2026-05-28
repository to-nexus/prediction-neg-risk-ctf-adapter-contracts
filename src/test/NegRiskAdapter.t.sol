// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {console2 as console} from "lib/forge-std/src/Test.sol";
import {TestHelper} from "../dev/TestHelper.sol";
import {NegRiskAdapter} from "../NegRiskAdapter.sol";
import {Vault} from "../Vault.sol";
import {WrappedCollateral} from "../WrappedCollateral.sol";
import {IConditionalTokens, IERC1155} from "../interfaces/IConditionalTokens.sol";
import {NegRiskIdLib} from "../libraries/NegRiskIdLib.sol";
import {USDC} from "./mock/USDC.sol";
import {CtfDeployer} from "./helpers/CtfDeployer.sol";

/// @notice Integration tests for NegRiskAdapter.convertPositions over a 4-question market.
///         Snapshots full state before/after each conversion so the input → output
///         relationship is fully observable from the logs (run with -vv).
/// @dev    Requires `forge build` to have been run in `prediction-contracts/` so the
///         ConditionalTokens + ERC1967Proxy artifacts are present on disk.
contract NegRiskAdapterConvertPositionsTest is TestHelper {
    uint8 internal constant QUESTION_COUNT = 4;
    uint256 internal constant AMOUNT = 1_000e6;

    USDC internal col;
    IConditionalTokens internal ctf;
    Vault internal vault;
    NegRiskAdapter internal adapter;
    WrappedCollateral internal wcol;

    bytes32 internal marketId;

    address internal ctfAdmin;
    address internal oracle;
    address internal user;
    address internal burnAddr;

    /// @dev Aggregate state across user / burn / vault / wcol used for delta assertions.
    struct State {
        uint256[] userYes;
        uint256[] userNo;
        uint256[] burnYes;
        uint256[] burnNo;
        uint256 userCol;
        uint256 vaultCol;
        uint256 wcolSupply;
        uint256 wcolUnderlying;
        uint256 adapterWcol;
    }

    function setUp() public {
        ctfAdmin = carly;
        oracle = alice;
        user = brian;

        col = new USDC();
        ctf = CtfDeployer.deploy(vm, ctfAdmin);

        vault = new Vault();
        adapter = new NegRiskAdapter(address(ctf), address(col), address(vault));
        wcol = adapter.wcol();
        burnAddr = adapter.NO_TOKEN_BURN_ADDRESS();

        CtfDeployer.grantSenderRole(vm, ctf, ctfAdmin, address(adapter));

        vm.startPrank(oracle);
        marketId = adapter.prepareMarket(0, bytes("neg-risk-convertPositions-test"));
        for (uint8 i = 0; i < QUESTION_COUNT; ++i) {
            adapter.prepareQuestion(marketId, abi.encodePacked("Q", i));
        }
        vm.stopPrank();

        col.mint(user, 100 * AMOUNT);

        vm.startPrank(user);
        col.approve(address(adapter), type(uint256).max);
        IERC1155(address(ctf)).setApprovalForAll(address(adapter), true);
        vm.stopPrank();
    }

    // ------------------------------------------------------------
    // convertPositions — 1 NO position (indexSet = 0b0001)
    //
    // User surrenders 1 NO_0. Expected:
    //   - NO_0:  -AMOUNT  (transferred to burn)
    //   - YES_1, YES_2, YES_3: +AMOUNT each (received from adapter-side splits)
    //   - YES_0: unchanged (user retains the YES from their own split)
    //   - collateral: unchanged  (multiplier = 1 - 1 = 0)
    // ------------------------------------------------------------

    function test_convertPositions_oneNo() public {
        _splitForUser(user, _u8(0));

        State memory before = _snapshot("BEFORE (1 NO)");

        vm.prank(user);
        adapter.convertPositions(marketId, 1, AMOUNT);

        State memory afterS = _snapshot("AFTER  (1 NO)");

        _assertConversion({
            before: before,
            afterS: afterS,
            indexSet: 1,
            noCount: 1,
            yesCount: 3,
            expectedColPayout: 0
        });
    }

    // ------------------------------------------------------------
    // convertPositions — 2 NO positions (indexSet = 0b0011)
    //
    // User surrenders NO_0 + NO_1. Expected:
    //   - NO_0, NO_1:  -AMOUNT each
    //   - YES_2, YES_3: +AMOUNT each
    //   - collateral:  +AMOUNT (multiplier = 2 - 1 = 1)
    // ------------------------------------------------------------

    function test_convertPositions_twoNos() public {
        _splitForUser(user, _u8(0, 1));

        State memory before = _snapshot("BEFORE (2 NOs)");

        vm.prank(user);
        adapter.convertPositions(marketId, 3, AMOUNT);

        State memory afterS = _snapshot("AFTER  (2 NOs)");

        _assertConversion({
            before: before,
            afterS: afterS,
            indexSet: 3,
            noCount: 2,
            yesCount: 2,
            expectedColPayout: AMOUNT
        });
    }

    // ------------------------------------------------------------
    // convertPositions — 3 NO positions (indexSet = 0b0111)
    //
    // User surrenders NO_0 + NO_1 + NO_2. Expected:
    //   - NO_0, NO_1, NO_2: -AMOUNT each
    //   - YES_3: +AMOUNT
    //   - collateral: +2*AMOUNT (multiplier = 3 - 1 = 2)
    // ------------------------------------------------------------

    function test_convertPositions_threeNos() public {
        _splitForUser(user, _u8(0, 1, 2));

        State memory before = _snapshot("BEFORE (3 NOs)");

        vm.prank(user);
        adapter.convertPositions(marketId, 7, AMOUNT);

        State memory afterS = _snapshot("AFTER  (3 NOs)");

        _assertConversion({
            before: before,
            afterS: afterS,
            indexSet: 7,
            noCount: 3,
            yesCount: 1,
            expectedColPayout: 2 * AMOUNT
        });
    }

    // ------------------------------------------------------------
    // convertPositions — 2 NOs at non-contiguous indices (indexSet = 0b0101)
    //
    // User surrenders NO_0 + NO_2. Expected:
    //   - NO_0, NO_2: -AMOUNT each
    //   - YES_1, YES_3: +AMOUNT each
    //   - collateral: +AMOUNT
    // ------------------------------------------------------------

    function test_convertPositions_twoNos_nonContiguous() public {
        _splitForUser(user, _u8(0, 2));

        State memory before = _snapshot("BEFORE (2 NOs non-contig)");

        vm.prank(user);
        adapter.convertPositions(marketId, 5, AMOUNT);

        State memory afterS = _snapshot("AFTER  (2 NOs non-contig)");

        _assertConversion({
            before: before,
            afterS: afterS,
            indexSet: 5,
            noCount: 2,
            yesCount: 2,
            expectedColPayout: AMOUNT
        });
    }

    // ------------------------------------------------------------
    // unified diff assertions
    // ------------------------------------------------------------

    /// @dev Asserts the *delta* between two snapshots matches the formal convertPositions
    ///      semantics:
    ///        for i in indexSet     : user.NO_i -= AMOUNT,  burn.NO_i += AMOUNT
    ///        for j NOT in indexSet : user.YES_j += AMOUNT, burn.NO_j  += AMOUNT (adapter-side split)
    ///        user.collateral      += (noCount - 1) * AMOUNT
    ///        vault.collateral     += 0   (fee = 0 in setUp)
    ///        wcol.totalSupply     += yesCount * AMOUNT
    ///        wcol.underlying      -= (noCount - 1) * AMOUNT  (released as collateral payout)
    function _assertConversion(
        State memory before,
        State memory afterS,
        uint256 indexSet,
        uint256 noCount,
        uint256 yesCount,
        uint256 expectedColPayout
    ) internal pure {
        assertEq(noCount + yesCount, QUESTION_COUNT, "noCount + yesCount must equal QUESTION_COUNT");

        for (uint8 i = 0; i < QUESTION_COUNT; ++i) {
            bool isNoSurrender = (indexSet & (1 << i)) > 0;

            if (isNoSurrender) {
                // user surrenders 1 NO at this index
                assertEq(before.userNo[i] - afterS.userNo[i], AMOUNT, "user NO delta (surrender)");
                // user's YES at this index is untouched by convertPositions
                assertEq(afterS.userYes[i], before.userYes[i], "user YES delta (untouched on surrender index)");
                // burn address receives 1 NO at this index (the user's NO)
                assertEq(afterS.burnNo[i] - before.burnNo[i], AMOUNT, "burn NO delta (user surrender)");
            } else {
                // user receives 1 YES at this index
                assertEq(afterS.userYes[i] - before.userYes[i], AMOUNT, "user YES delta (received)");
                // user's NO at this index is unchanged
                assertEq(afterS.userNo[i], before.userNo[i], "user NO delta (not in indexSet)");
                // burn address receives 1 NO at this index (from adapter's internal split)
                assertEq(afterS.burnNo[i] - before.burnNo[i], AMOUNT, "burn NO delta (adapter split)");
            }
            // No YES is ever burned
            assertEq(afterS.burnYes[i], before.burnYes[i], "burn YES delta must be zero");
        }

        // collateral payout to user
        assertEq(afterS.userCol - before.userCol, expectedColPayout, "user collateral delta");
        // no fee → vault unchanged
        assertEq(afterS.vaultCol, before.vaultCol, "vault collateral delta must be zero (fee=0)");
        // wcol total supply rose by exactly yesCount * AMOUNT (the new mint inside convertPositions)
        assertEq(afterS.wcolSupply - before.wcolSupply, yesCount * AMOUNT, "wcol totalSupply delta");
        // wcol underlying fell by exactly the released collateral (since no fee, all went to user)
        assertEq(before.wcolUnderlying - afterS.wcolUnderlying, expectedColPayout, "wcol underlying delta");
        // adapter holds no wcol after the conversion settles
        assertEq(afterS.adapterWcol, 0, "adapter must hold zero wcol after conversion");
    }

    // ------------------------------------------------------------
    // helpers
    // ------------------------------------------------------------

    function _splitForUser(address _user, uint8[] memory _questionIndices) internal {
        for (uint256 i = 0; i < _questionIndices.length; ++i) {
            bytes32 questionId = NegRiskIdLib.getQuestionId(marketId, _questionIndices[i]);
            bytes32 conditionId = adapter.getConditionId(questionId);
            vm.prank(_user);
            adapter.splitPosition(conditionId, AMOUNT);
        }
    }

    function _snapshot(string memory label) internal returns (State memory s) {
        s.userYes = new uint256[](QUESTION_COUNT);
        s.userNo = new uint256[](QUESTION_COUNT);
        s.burnYes = new uint256[](QUESTION_COUNT);
        s.burnNo = new uint256[](QUESTION_COUNT);
        for (uint8 i = 0; i < QUESTION_COUNT; ++i) {
            s.userYes[i] = _balanceYes(user, i);
            s.userNo[i] = _balanceNo(user, i);
            s.burnYes[i] = _balanceYes(burnAddr, i);
            s.burnNo[i] = _balanceNo(burnAddr, i);
        }
        s.userCol = col.balanceOf(user);
        s.vaultCol = col.balanceOf(address(vault));
        s.wcolSupply = wcol.totalSupply();
        s.wcolUnderlying = col.balanceOf(address(wcol));
        s.adapterWcol = wcol.balanceOf(address(adapter));

        console.log("--- %s ---", label);
        for (uint8 i = 0; i < QUESTION_COUNT; ++i) {
            console.log("  Q%d  user YES=%d  NO=%d", i, s.userYes[i], s.userNo[i]);
        }
        for (uint8 i = 0; i < QUESTION_COUNT; ++i) {
            console.log("        burn NO_%d=%d", i, s.burnNo[i]);
        }
        console.log("  user.col=%d  vault.col=%d", s.userCol, s.vaultCol);
        console.log("  wcol.totalSupply=%d  wcol.underlying=%d", s.wcolSupply, s.wcolUnderlying);
        console.log("  adapter.wcol=%d", s.adapterWcol);
    }

    function _balanceYes(address _user, uint8 _questionIndex) internal view returns (uint256) {
        bytes32 questionId = NegRiskIdLib.getQuestionId(marketId, _questionIndex);
        return adapter.balanceOf(_user, adapter.getPositionId(questionId, true));
    }

    function _balanceNo(address _user, uint8 _questionIndex) internal view returns (uint256) {
        bytes32 questionId = NegRiskIdLib.getQuestionId(marketId, _questionIndex);
        return adapter.balanceOf(_user, adapter.getPositionId(questionId, false));
    }

    function _u8(uint8 a) internal pure returns (uint8[] memory arr) {
        arr = new uint8[](1);
        arr[0] = a;
    }

    function _u8(uint8 a, uint8 b) internal pure returns (uint8[] memory arr) {
        arr = new uint8[](2);
        arr[0] = a;
        arr[1] = b;
    }

    function _u8(uint8 a, uint8 b, uint8 c) internal pure returns (uint8[] memory arr) {
        arr = new uint8[](3);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
    }
}
