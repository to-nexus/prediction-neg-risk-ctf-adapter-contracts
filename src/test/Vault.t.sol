// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {TestHelper} from "../dev/TestHelper.sol";
import {Vault} from "../Vault.sol";
import {USDC} from "./mock/USDC.sol";

contract VaultTest is TestHelper {
    Vault vault;
    USDC usdc;

    function setUp() public {
        vm.prank(alice);
        vault = new Vault();
        usdc = new USDC();
    }

    function test_transferERC20(uint64 _a, uint64 _b, uint64 _c) public {
        uint256 s = uint256(_a);
        uint256 m = s + uint256(_b);
        uint256 l = m + uint256(_c);

        usdc.mint(brian, l);

        vm.prank(brian);
        usdc.transfer(address(vault), m);

        vm.prank(alice);
        vault.transferERC20(address(usdc), devin, s);

        assertEq(usdc.balanceOf(devin), s);
        assertEq(usdc.balanceOf(address(vault)), m - s);
        assertEq(usdc.balanceOf(brian), l - m);
    }
}
