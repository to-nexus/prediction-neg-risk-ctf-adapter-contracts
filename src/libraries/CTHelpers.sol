// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity >=0.5.1;

/// @title CTHelpers
/// @notice ID derivation helpers aligned with the team ConditionalTokens implementation.
/// @dev The three formulas below MUST remain byte-for-byte identical to
///      `prediction-contracts/src/ctf/ConditionalTokens.sol::_getConditionId`,
///      `_getCollectionId`, and `_getPositionId`. Any divergence breaks positionId
///      parity between off-chain (this library) and on-chain (the CT contract),
///      causing ERC1155 balance mismatches during splitPosition / mergePositions /
///      safeBatchTransferFrom. A unit test in `src/test/libraries/CTHelpers.t.sol`
///      pins this invariant.
library CTHelpers {
    /// @dev Constructs a condition ID from an oracle, a question ID, and the outcome slot count.
    function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(oracle, questionId, outcomeSlotCount));
    }

    /// @dev Constructs an outcome collection ID from parent collection, condition, and index set.
    /// @dev The neg-risk adapter always passes `parentCollectionId == bytes32(0)`, so nested
    ///      collection composition is not exercised; the formula is flat keccak either way.
    function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(parentCollectionId, conditionId, indexSet));
    }

    /// @dev Constructs a position ID (ERC-1155 token ID) from a collateral token and a collection.
    function getPositionId(address collateralToken, bytes32 collectionId) internal pure returns (uint256) {
        return uint256(keccak256(abi.encodePacked(collateralToken, collectionId)));
    }
}
