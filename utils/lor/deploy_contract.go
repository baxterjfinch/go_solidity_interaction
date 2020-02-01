package lor

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    lor "github.com/hemlokc/go_solidity_interaction/compiled/lor"

)

func DeployContract(network string, key string) {
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
    auth.Value = big.NewInt(0)     // in wei
    // auth.GasLimit = uint64(300000) // in units
    // auth.GasPrice = big.NewInt(10000000)

    // input := "1.0"
    address, tx, instance, err := lor.DeployLor(auth, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("")
    fmt.Println("")
    fmt.Println("--------")
    fmt.Println("Deploying Contract From Address:")
    fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
    fmt.Println("--------")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("--------")
    fmt.Println("Den Contract Deployed! Hooray! Tx:")
    fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
    fmt.Println("--------")
    fmt.Println("")
    fmt.Println("")
    _ = instance
}
