package utils

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func Client(network string) {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client // we'll use this in the upcoming sections
}
