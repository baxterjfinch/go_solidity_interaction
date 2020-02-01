package lor

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

  lor "github.com/hemlokc/go_solidity_interaction/compiled/lor"
)

func LorInteract(network string, key string, string_address string) {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress(string_address)

    instance, err := lor.NewLor(address, client)
    if err != nil {
      log.Fatal(err)
    }

    callOpts := &bind.CallOpts{
      Pending: false,
    }

    name, err := instance.LorCaller.Name(callOpts)

    fmt.Println("")
    fmt.Println("--------")
    fmt.Println("Contract Name: ", name)

    totalSupply, err := instance.LorCaller.TotalSupply(callOpts)

    fmt.Println("Total Supply: ", totalSupply)
}

func GetLockedSupply(network string, string_address string) {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress(string_address)

    instance, err := lor.NewLor(address, client)
    if err != nil {
      log.Fatal(err)
    }

    callOpts := &bind.CallOpts{
      Pending: false,
    }

    locked, err := instance.LorCaller.GetLockedSupply(callOpts)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println("Locked Supply: ", locked)
}

func LockSupply(network string, key string, string_address string, amount int64)  {
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
    instance, err := lor.NewLor(address, client)
    if err != nil {
      log.Fatal(err)
    }

    lockedtx, err := instance.LorTransactor.SetLockedSupply(auth, big.NewInt(amount))
    if err != nil {
      log.Fatal(err)
    }
    // transaction, err := mint.MarshalJSON()
    // if err != nil {
    //   log.Fatal(err)
    // }
    txString := lockedtx.Hash().Hex()
    fmt.Println()
    fmt.Println("Locked")
    fmt.Println("--------")
    fmt.Println()
    fmt.Println("Transaction")
    fmt.Println("ID: ", txString)

    // fmt.Println("Size: ", transaction.size)
    // fmt.Println("From: ", transaction.from)
}
