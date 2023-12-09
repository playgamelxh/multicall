package main

import (
    "bufio"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "io"
    "multicall/contracts"
    "multicall/util"
    "os"
)

func main() {
    rpc := "https://bsc.publicnode.com"
    chainId := 56
    contractAddress := common.HexToAddress("0xd80130CDF6f4c4d4Aa71CaB1aA0C7Fff4b211572")

    // client
    client, _ := util.GetEtherClientByRpc(rpc, chainId)
    contract, err := contracts.NewAbc(contractAddress, client)
    if err != nil {
        fmt.Println("new wsteth contract error:%v", err)
        return
    }
    fmt.Println(contract)

    file, _ := os.Open("./less3000.txt")
    reader := bufio.NewReader(file)

    over := false
    addresses := make([]common.Address, 0)
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            over = true
        }
        if err != nil {
            return
        }
        fmt.Println(string(line))
        addresses = append(addresses, common.HexToAddress(string(line)))
        if len(addresses) == 100 || over {
            _, err := contract.ExecuteUSDC(&bind.TransactOpts{}, addresses)
            if err != nil {
                fmt.Println("get allowance error:%v", err.Error())
                return
            }
        }

        if over {
            break
        }
    }
}
