package main

import (
    "github.com/hemlokc/go_solidity_interaction/utils"
)

const ethAddress = ""
const network = ""
const privateKey = ""

func main() {
    // utils.Client(network)
    // utils.Address()
    // utils.Balances(ethAddress, network)
    utils.DeployERC1155(network, privateKey)
}
