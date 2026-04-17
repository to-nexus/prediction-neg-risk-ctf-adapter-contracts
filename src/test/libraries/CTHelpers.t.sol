// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Test} from "lib/forge-std/src/Test.sol";
import {CTHelpers} from "../../libraries/CTHelpers.sol";

/// @notice Pins the ID derivation formulas in CTHelpers to the keccak forms used by the team
///         ConditionalTokens implementation (prediction-contracts/src/ctf/ConditionalTokens.sol).
/// @dev The expected values in each test are computed inline with the same
///      keccak256(abi.encodePacked(...)) expression used by the team CT contract. If CTHelpers
///      diverges from that shape, these tests fail — which signals an off-chain / on-chain
///      positionId mismatch before it reaches production.
contract CTHelpersTest is Test {
    // ------------------------------------------------------------
    // getConditionId
    // ------------------------------------------------------------

    function test_getConditionId_matchesKeccakFormula_fixedVector() public pure {
        address oracle = address(0xCafEBAbECAFEbAbEcaFEbabECAfebAbEcAFEBaBe);
        bytes32 questionId = keccak256("question-1");
        uint256 outcomeSlotCount = 2;

        bytes32 expected = keccak256(abi.encodePacked(oracle, questionId, outcomeSlotCount));
        bytes32 actual = CTHelpers.getConditionId(oracle, questionId, outcomeSlotCount);

        assertEq(actual, expected, "getConditionId must match team CT _getConditionId keccak formula");
    }

    function testFuzz_getConditionId_matchesKeccakFormula(address oracle, bytes32 questionId, uint256 outcomeSlotCount)
        public
        pure
    {
        bytes32 expected = keccak256(abi.encodePacked(oracle, questionId, outcomeSlotCount));
        bytes32 actual = CTHelpers.getConditionId(oracle, questionId, outcomeSlotCount);
        assertEq(actual, expected);
    }

    // ------------------------------------------------------------
    // getCollectionId
    // ------------------------------------------------------------

    function test_getCollectionId_matchesKeccakFormula_parentZero_indexYes() public pure {
        bytes32 conditionId = keccak256("condition-1");
        uint256 indexSet = 1; // 0b01 — YES

        bytes32 expected = keccak256(abi.encodePacked(bytes32(0), conditionId, indexSet));
        bytes32 actual = CTHelpers.getCollectionId(bytes32(0), conditionId, indexSet);

        assertEq(actual, expected, "getCollectionId(0, cid, 1) must match team CT _getCollectionId");
    }

    function test_getCollectionId_matchesKeccakFormula_parentZero_indexNo() public pure {
        bytes32 conditionId = keccak256("condition-1");
        uint256 indexSet = 2; // 0b10 — NO

        bytes32 expected = keccak256(abi.encodePacked(bytes32(0), conditionId, indexSet));
        bytes32 actual = CTHelpers.getCollectionId(bytes32(0), conditionId, indexSet);

        assertEq(actual, expected);
    }

    function test_getCollectionId_matchesKeccakFormula_nonZeroParent() public pure {
        bytes32 parent = keccak256("parent-collection");
        bytes32 conditionId = keccak256("condition-2");
        uint256 indexSet = 7; // 0b111

        bytes32 expected = keccak256(abi.encodePacked(parent, conditionId, indexSet));
        bytes32 actual = CTHelpers.getCollectionId(parent, conditionId, indexSet);

        assertEq(actual, expected, "Non-zero parent must also use flat keccak, not EC composition");
    }

    function testFuzz_getCollectionId_matchesKeccakFormula(bytes32 parent, bytes32 conditionId, uint256 indexSet)
        public
        pure
    {
        bytes32 expected = keccak256(abi.encodePacked(parent, conditionId, indexSet));
        bytes32 actual = CTHelpers.getCollectionId(parent, conditionId, indexSet);
        assertEq(actual, expected);
    }

    // ------------------------------------------------------------
    // getPositionId
    // ------------------------------------------------------------

    function test_getPositionId_matchesKeccakFormula_fixedVector() public pure {
        address collateral = address(0xDeaDbeefdEAdbeefdEadbEEFdeadbeEFdEaDbeeF);
        bytes32 collectionId = keccak256("collection-1");

        uint256 expected = uint256(keccak256(abi.encodePacked(collateral, collectionId)));
        uint256 actual = CTHelpers.getPositionId(collateral, collectionId);

        assertEq(actual, expected, "getPositionId must match team CT _getPositionId keccak formula");
    }

    function testFuzz_getPositionId_matchesKeccakFormula(address collateral, bytes32 collectionId) public pure {
        uint256 expected = uint256(keccak256(abi.encodePacked(collateral, collectionId)));
        uint256 actual = CTHelpers.getPositionId(collateral, collectionId);
        assertEq(actual, expected);
    }
}
