package main

import (
    "bufio"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "io"
    "multicall/util"
    "os"
)

func main() {
    rpc := "https://bsc.publicnode.com"
    chainId := 56
    multicall := "0x1ee38d535d541c55c9dae27b12edf090c608e6fb"
    spender := common.HexToAddress("0xc9316435aef3762f25d2d0a31ec95fe4e6798085")
    //usdcToken := common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d")

    caller, err := util.NewMultiCall(rpc, multicall)
    if err != nil {
        return
    }

    var userCalls = make([]util.Call, 0)

    file, _ := os.Open("./less3000.txt")
    if err != nil {
        return
    }
    reader := bufio.NewReader(file)

    over := false
    addresses := make([]string, 0)
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            over = true
        }
        if err != nil {
            return
        }
        addresses = append(addresses, string(line))
        if len(addresses) == 100 || over {
            for _, address := range addresses {
                // read contract
                userCalls = append(userCalls, util.GetMethodaCall("DillAllowance", spender, common.HexToAddress(address)))
                response, err := caller.Execute(userCalls)
                if err != nil {
                    fmt.Println("rpc:", rpc, "multicall:", multicall, "userAddress:", address, "chainId:", chainId)
                    return
                }
                fmt.Println(response)

                // test balanc
                //userCalls = append(userCalls, util.GetBalanceCall("DillAllowance", usdcToken, common.HexToAddress(address)))
                //response, err := caller.Execute(userCalls)
                //if err != nil {
                //    fmt.Println("rpc:", rpc, "multicall:", multicall, "userAddress:", address, "chainId:", chainId)
                //    return
                //}
                //fmt.Println(response)
            }
        }

        if over {
            break
        }
    }

}
