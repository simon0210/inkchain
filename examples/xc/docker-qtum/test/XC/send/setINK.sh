#! /bin/bash

source p_init.sh

qcli sendtocontract $XCHexAddress `solar encode contracts/XC.sol setINK '["'$INKHexAddress'"]'` 0 6000000 0.0000004 $XCOwner