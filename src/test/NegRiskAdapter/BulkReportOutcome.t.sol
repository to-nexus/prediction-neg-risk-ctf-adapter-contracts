// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Vm} from "lib/forge-std/src/Vm.sol";
import {NegRiskAdapter_SetUp} from "src/test/NegRiskAdapter/NegRiskAdapterSetUp.sol";
import {NegRiskIdLib} from "src/libraries/NegRiskIdLib.sol";

contract NegRiskAdapter_BulkReportOutcome_Test is NegRiskAdapter_SetUp {
    uint8 constant QUESTION_COUNT = 3;
    bytes32 marketId;
    bytes32[] questionIds;
    bytes32[] conditionIds;

    function setUp() public override {
        NegRiskAdapter_SetUp.setUp();

        uint256 feeBips = 0;
        bytes memory data = new bytes(0);

        vm.startPrank(oracle);
        marketId = nrAdapter.prepareMarket(feeBips, data);

        for (uint8 i = 0; i < QUESTION_COUNT; i++) {
            bytes32 qId = nrAdapter.prepareQuestion(marketId, data);
            questionIds.push(qId);
            conditionIds.push(nrAdapter.getConditionId(qId));
        }
        vm.stopPrank();
    }

    // ---------------------------------------------------------------
    //  Happy Path
    // ---------------------------------------------------------------

    function test_bulkReportOutcome_winnerFirst() public {
        uint256 winnerIndex = 0;

        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, winnerIndex);

        // winner (index 0) resolved true
        _assertQuestionResolved(0, true);
        // losers resolved false
        _assertQuestionResolved(1, false);
        _assertQuestionResolved(2, false);
    }

    function test_bulkReportOutcome_winnerLast() public {
        uint256 winnerIndex = 2;

        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, winnerIndex);

        _assertQuestionResolved(0, false);
        _assertQuestionResolved(1, false);
        _assertQuestionResolved(2, true);
    }

    function test_bulkReportOutcome_twoQuestions() public {
        // prepare a fresh 2-question market
        bytes memory data = new bytes(0);

        vm.startPrank(oracle);
        bytes32 mId2 = nrAdapter.prepareMarket(1, data);
        bytes32 q0 = nrAdapter.prepareQuestion(mId2, data);
        bytes32 q1 = nrAdapter.prepareQuestion(mId2, data);
        vm.stopPrank();

        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(mId2, 1);

        // question 0 -> false
        bytes32 cond0 = nrAdapter.getConditionId(q0);
        assertEq(ctf.payoutNumerators(cond0, 0), 0);
        assertEq(ctf.payoutNumerators(cond0, 1), 1);

        // question 1 -> true
        bytes32 cond1 = nrAdapter.getConditionId(q1);
        assertEq(ctf.payoutNumerators(cond1, 0), 1);
        assertEq(ctf.payoutNumerators(cond1, 1), 0);

        assertTrue(nrAdapter.getDetermined(mId2));
        assertEq(nrAdapter.getResult(mId2), 1);
    }

    function test_bulkReportOutcome_events() public {
        uint256 winnerIndex = 1;

        vm.recordLogs();

        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, winnerIndex);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        // Collect OutcomeReported events
        bytes32 outcomeReportedSig = keccak256("OutcomeReported(bytes32,bytes32,bool)");
        bytes32 bulkOutcomeSig = keccak256("BulkOutcomeReported(bytes32,uint256,uint256)");

        uint256 outcomeCount;
        bool foundBulk;

        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].topics[0] == outcomeReportedSig) {
                bytes32 emittedMarketId = logs[i].topics[1];
                bytes32 emittedQuestionId = logs[i].topics[2];
                bool emittedOutcome = abi.decode(logs[i].data, (bool));

                bytes32 expectedQId = NegRiskIdLib.getQuestionId(marketId, uint8(outcomeCount));
                bool expectedOutcome = (outcomeCount == winnerIndex);

                assertEq(emittedMarketId, marketId);
                assertEq(emittedQuestionId, expectedQId);
                assertEq(emittedOutcome, expectedOutcome);

                outcomeCount++;
            } else if (logs[i].topics[0] == bulkOutcomeSig) {
                bytes32 emittedMarketId = logs[i].topics[1];
                (uint256 emittedWinner, uint256 emittedCount) = abi.decode(logs[i].data, (uint256, uint256));

                assertEq(emittedMarketId, marketId);
                assertEq(emittedWinner, winnerIndex);
                assertEq(emittedCount, QUESTION_COUNT);

                foundBulk = true;
            }
        }

        assertEq(outcomeCount, QUESTION_COUNT);
        assertTrue(foundBulk);
    }

    function test_bulkReportOutcome_marketDetermined() public {
        uint256 winnerIndex = 1;

        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, winnerIndex);

        assertTrue(nrAdapter.getDetermined(marketId));
        assertEq(nrAdapter.getResult(marketId), winnerIndex);
    }

    // ---------------------------------------------------------------
    //  Revert Cases
    // ---------------------------------------------------------------

    function test_bulkReportOutcome_revert_notOracle() public {
        vm.expectRevert(OnlyOracle.selector);
        vm.prank(alice);
        nrAdapter.bulkReportOutcome(marketId, 0);
    }

    function test_bulkReportOutcome_revert_marketNotPrepared() public {
        bytes32 fakeMarketId = keccak256("fake");

        vm.expectRevert(MarketNotPrepared.selector);
        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(fakeMarketId, 0);
    }

    function test_bulkReportOutcome_revert_indexOutOfBounds() public {
        vm.expectRevert(IndexOutOfBounds.selector);
        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, QUESTION_COUNT);
    }

    function test_bulkReportOutcome_revert_alreadyDetermined() public {
        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, 0);

        vm.expectRevert(MarketAlreadyDetermined.selector);
        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, 1);
    }

    // ---------------------------------------------------------------
    //  Integration
    // ---------------------------------------------------------------

    function test_bulkReportOutcome_redeemAfterResolve() public {
        uint256 amount = 100 ether;
        uint256 winnerIndex = 1;

        // brian buys YES on question 1 (the winner)
        // carly buys YES on question 0 (a loser)
        {
            // alice splits positions on all questions
            vm.startPrank(alice);
            for (uint256 i = 0; i < QUESTION_COUNT; i++) {
                usdc.mint(alice, amount);
                usdc.approve(address(nrAdapter), amount);
                nrAdapter.splitPosition(conditionIds[i], amount);
            }

            // transfer YES token of question 1 to brian
            uint256 winnerYesId = nrAdapter.getPositionId(questionIds[winnerIndex], true);
            ctf.safeTransferFrom(alice, brian, winnerYesId, amount, "");

            // transfer YES token of question 0 to carly
            uint256 loserYesId = nrAdapter.getPositionId(questionIds[0], true);
            ctf.safeTransferFrom(alice, carly, loserYesId, amount, "");

            vm.stopPrank();
        }

        // bulk resolve
        {
            vm.prank(oracle);
            nrAdapter.bulkReportOutcome(marketId, winnerIndex);
        }

        // brian redeems winner YES tokens -> gets collateral back
        {
            vm.startPrank(brian);
            ctf.setApprovalForAll(address(nrAdapter), true);
            uint256[] memory amounts = new uint256[](2);
            amounts[0] = amount; // YES tokens
            nrAdapter.redeemPositions(conditionIds[winnerIndex], amounts);
            assertEq(usdc.balanceOf(brian), amount);
            vm.stopPrank();
        }

        // carly redeems loser YES tokens -> gets 0
        {
            vm.startPrank(carly);
            ctf.setApprovalForAll(address(nrAdapter), true);
            uint256[] memory amounts = new uint256[](2);
            amounts[0] = amount; // YES tokens
            nrAdapter.redeemPositions(conditionIds[0], amounts);
            assertEq(usdc.balanceOf(carly), 0);
            vm.stopPrank();
        }
    }

    function test_bulkReportOutcome_thenIndividualReportReverts() public {
        vm.prank(oracle);
        nrAdapter.bulkReportOutcome(marketId, 0);

        // individual reportOutcome with true should revert (market already determined)
        bytes32 qId = NegRiskIdLib.getQuestionId(marketId, 1);
        vm.expectRevert(MarketAlreadyDetermined.selector);
        vm.prank(oracle);
        nrAdapter.reportOutcome(qId, true);
    }

    // ---------------------------------------------------------------
    //  Internal helpers
    // ---------------------------------------------------------------

    function _assertQuestionResolved(uint256 _index, bool _expectedOutcome) internal view {
        bytes32 condId = conditionIds[_index];
        assertEq(ctf.payoutDenominator(condId), 1);
        assertEq(ctf.payoutNumerators(condId, 0), _expectedOutcome ? 1 : 0);
        assertEq(ctf.payoutNumerators(condId, 1), _expectedOutcome ? 0 : 1);
    }
}
