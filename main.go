package main

import (
    "fmt"
    //LOR
    lor "github.com/hemlokc/go_solidity_interaction/utils/lor"

    //MATTER
    matter "github.com/hemlokc/go_solidity_interaction/utils/matter"
)

const ethAddress = ""
const network = "https://ropsten.infura.io/v3/b8238578cbe444438b2f80c928c18df1"
const privateKey = ""

// Deployed Matter Contract address
const matter_contract = ""
const lor_contract = ""

func main() {
    // utils.Client(network)
    // utils.Address()
    // utils.Balances(ethAddress, network)

    // LOR Contract Interaction
    // lor.DeployContract(network, privateKey)
    lor.LorInteract(network, privateKey, lor_contract)
    // utils.LockSupply(network, privateKey, lor_contract, 4338920)
    lor.GetLockedSupply(network, lor_contract)

    fmt.Println("--------")

    // Matter Contract Interaction
    matter.MatterInteract(network, privateKey, matter_contract)
    matter.MintMatter(network, privateKey, matter_contract)
}
