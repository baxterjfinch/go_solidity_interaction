package utils

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func Balances(ethAddress string, network string) {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress(ethAddress)
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance) // 25893180161173005034
}
