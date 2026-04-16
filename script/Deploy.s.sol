// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {Script, console} from "forge-std/Script.sol";

import {Vault} from "src/Vault.sol";
import {NegRiskAdapter} from "src/NegRiskAdapter.sol";

interface IAccessControl {
    function grantRole(bytes32 role, address account) external;
}

/// @title Deploy
/// @notice Neg-Risk Adapter 컨트랙트 배포 스크립트 (Adapter, WrappedCollateral, Vault)
/// @dev design-neg-risk-integration.md §3.1 참조.
///      NegRiskCtfExchange, NegRiskFeeModule, NegRiskOperator는 사용하지 않음 —
///      기존 PredictionExchange, IFeeController, AccessControl 기반 직접 호출로 대체.
///
///      Usage:
///        forge script script/Deploy.s.sol \
///          --sig "deploy(address,address)" <ctf> <collateral> \
///          --rpc-url <rpc> --broadcast
contract Deploy is Script {
    bytes32 internal constant SENDER_ROLE = keccak256("SENDER_ROLE");

    struct DeployedContracts {
        address vault;
        address negRiskAdapter;
        address wrappedCollateral;
    }

    /// @notice Vault + NegRiskAdapter 배포 (WrappedCollateral은 Adapter 내부에서 자동 생성)
    /// @param _ctf ConditionalTokens 주소 (prediction-contracts에서 배포된 것)
    /// @param _collateral 담보 토큰 주소 (e.g., BILL)
    function deploy(address _ctf, address _collateral) public returns (DeployedContracts memory deployed) {
        vm.startBroadcast();

        Vault vault = new Vault();
        deployed.vault = address(vault);
        console.log("Vault:", deployed.vault);

        NegRiskAdapter negRiskAdapter = new NegRiskAdapter(_ctf, _collateral, deployed.vault);
        deployed.negRiskAdapter = address(negRiskAdapter);
        deployed.wrappedCollateral = address(negRiskAdapter.wcol());
        console.log("NegRiskAdapter:", deployed.negRiskAdapter);
        console.log("WrappedCollateral:", deployed.wrappedCollateral);

        vm.stopBroadcast();
    }
}
 