// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import {Vm} from "lib/forge-std/src/Vm.sol";
import {IConditionalTokens} from "../../interfaces/IConditionalTokens.sol";

interface IConditionalTokensAdmin {
    function initialize(address initialAdmin, uint48 initialDelay) external;
    function grantRole(bytes32 role, address account) external;
}

/// @notice Deploys the team ConditionalTokens implementation behind an ERC1967Proxy
///         and initializes it. Caller grants SENDER_ROLE to the operator (e.g., the
///         NegRiskAdapter) after the adapter is constructed, because the adapter address
///         is only known after deployment.
/// @dev Bytecode is loaded from prediction-contracts forge artifacts via vm.deployCode.
///      The path is whitelisted in foundry.toml `fs_permissions`. Run `forge build` in
///      prediction-contracts before running these tests.
library CtfDeployer {
    bytes32 internal constant SENDER_ROLE = keccak256("SENDER_ROLE");

    function deploy(Vm vm, address admin) internal returns (IConditionalTokens) {
        address impl = vm.deployCode("../prediction-contracts/out/ConditionalTokens.sol/ConditionalTokens.json");
        bytes memory initData = abi.encodeWithSelector(IConditionalTokensAdmin.initialize.selector, admin, uint48(0));
        address proxy =
            vm.deployCode("../prediction-contracts/out/ERC1967Proxy.sol/ERC1967Proxy.json", abi.encode(impl, initData));
        return IConditionalTokens(proxy);
    }

    function grantSenderRole(Vm vm, IConditionalTokens ctf, address admin, address account) internal {
        vm.prank(admin);
        IConditionalTokensAdmin(address(ctf)).grantRole(SENDER_ROLE, account);
    }
}
