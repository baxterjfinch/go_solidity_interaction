package matter

import (
  "context"
  "crypto/ecdsa"
  "fmt"
  "log"
  "math/big"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/common"
  // "github.com/ethereum/go-ethereum/core/types"

    // matter "github.com/hemlokc/go_solidity_interaction/test_go_contracts"
    matter "github.com/hemlokc/go_solidity_interaction/compiled/matter"
)

func MatterInteract(network string, key string, string_address string) {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress(string_address)

    instance, err := matter.NewMatter(address, client)
    if err != nil {
      log.Fatal(err)
    }

    callOpts := &bind.CallOpts{
      Pending: false,
    }

    name, err := instance.MatterCaller.Name(callOpts)

    fmt.Println("")
    fmt.Println("--------")
    fmt.Println("Contract Name: ", name)

    lastMinted, err := instance.MatterCaller.LastMinted(callOpts)
    fmt.Println("Last Minted: ", lastMinted)

    maxSupply, err := instance.MatterCaller.MaxSupply(callOpts)
    fmt.Println("Maximum Supply: ", maxSupply)

    minMintInterval, err := instance.MatterCaller.MinMintInterval(callOpts)
    fmt.Println("Minimum Mint Interval: ", minMintInterval)

    mintHour, err := instance.MatterCaller.MintHour(callOpts)
    fmt.Println("Mint Hour: ", mintHour)
    fmt.Println("--------")
    fmt.Println("")
}

func MintMatter(network string, key string, string_address string)  {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = 900000

    address := common.HexToAddress(string_address)
    instance, err := matter.NewMatter(address, client)
    if err != nil {
      log.Fatal(err)
    }

    mint, err := instance.MatterTransactor.Mint(auth)
    if err != nil {
      log.Fatal(err)
    }
    // transaction, err := mint.MarshalJSON()
    // if err != nil {
    //   log.Fatal(err)
    // }
    txString := mint.Hash().Hex()
    fmt.Println()
    fmt.Println("Minted")
    fmt.Println("--------")
    fmt.Println()
    fmt.Println("Transaction")
    fmt.Println("ID: ", txString)

    // fmt.Println("Size: ", transaction.size)
    // fmt.Println("From: ", transaction.from)
}
