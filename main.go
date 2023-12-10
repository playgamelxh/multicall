package main

import (
    "bufio"
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "io"
    "math/big"
    "multicall/contracts"
    "multicall/util"
    "os"
)

func main() {
    //rpc := "https://bsc.publicnode.com"
    rpc := "https://bsc-mainnet.nodereal.io/v1/1f0d76d59db449b3be10903fb7f5adc7"
    chainId := 56
    contractAddress := common.HexToAddress("0x55416aBFc4B431684Ab2b37E5e3a6B622109A00E")
    //
    //rpc := "https://eth-mainnet.nodereal.io/v1/1f0d76d59db449b3be10903fb7f5adc7"
    //chainId := 1
    //contractAddress := common.HexToAddress("0x55416aBFc4B431684Ab2b37E5e3a6B622109A00E")

    PrivateKey, _ := crypto.HexToECDSA("")
    PublicAddress := crypto.PubkeyToAddress(PrivateKey.PublicKey)
    fmt.Println(PublicAddress)

    // client
    client, _ := util.GetEtherClientByRpc(rpc, chainId)
    contract, err := contracts.NewAbc(contractAddress, client)
    if err != nil {
        fmt.Println("new wsteth contract error:%v", err)
        return
    }
    fmt.Println(contract)

    file, _ := os.Open("./less3000.txt")
    //file, _ := os.Open("./test.txt")
    reader := bufio.NewReader(file)

    over := false
    addresses := make([]common.Address, 0)
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            over = true
        }
        if err != nil {
            fmt.Println("read file error:%v", err.Error())
            //return
        }
        fmt.Println(string(line))
        addresses = append(addresses, common.HexToAddress(string(line)))
        if len(addresses) >= 50 || over {
            opts, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(int64(chainId)))
            opts.GasLimit = uint64(10000000)

            gasPrice, err := client.SuggestGasPrice(context.Background())
            opts.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(11)).Div(gasPrice, big.NewInt(10))

            nonce, err := client.PendingNonceAt(context.Background(), PublicAddress)
            if err != nil {
                fmt.Println("get nonce err:" + err.Error())
            }
            opts.Nonce = big.NewInt(int64(nonce))

            transaction, err := contract.ExecuteUSDC(opts, addresses)
            if err != nil {
                fmt.Println("execute error:%v", err.Error())
                return
            }

            receipt, err := bind.WaitMined(context.Background(), client, transaction)
            if err != nil {
                fmt.Println("wait mined err:" + err.Error())
            }
            fmt.Println("transaction receipt : ", receipt.Status)

            addresses = make([]common.Address, 0)

        }

        if over {
            break
        }
    }
}
