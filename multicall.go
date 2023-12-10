package main

import (
    "bufio"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/shopspring/decimal"
    "io"
    "math/big"
    "multicall/util"
    "os"
)

func main() {
    rpc := "https://bsc.publicnode.com"
    //chainId := 56
    multicall := "0x1ee38d535d541c55c9dae27b12edf090c608e6fb"
    spender := common.HexToAddress("0xc9316435aef3762f25d2d0a31ec95fe4e6798085")
    //usdcToken := common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d")
    //usdtToken := common.HexToAddress("0x55d398326f99059ff775485246999027b3197955")
    //ethToken := common.HexToAddress("0x2170ed0880ac9a755fd29b2688956bd959f933f8")
    wbtcToken := common.HexToAddress("0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c")

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

    usdcFile, _ := os.OpenFile("./wbtc.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

    over := false
    addresses := make([]string, 0)
    for {
        line, _, err := reader.ReadLine()
        //fmt.Println(line)
        if err == io.EOF {
            fmt.Println("read over", err.Error())
            over = true
        }
        if err != nil {
            return
        }
        addresses = append(addresses, string(line))
        if len(addresses) == 100 || over {
            userCalls = make([]util.Call, 0)
            for _, address := range addresses {
                //// read contract
                //userCalls = append(userCalls, util.GetMethodaCall("DillAllowance", spender, common.HexToAddress(address)))
                //response, err := caller.Execute(userCalls)
                //if err != nil {
                //    fmt.Println("rpc:", rpc, "multicall:", multicall, "userAddress:", address, "chainId:", chainId)
                //    return
                //}
                //fmt.Println(response)

                // test balanc
                userCalls = append(userCalls, util.GetApproveCall("DillAllowance", wbtcToken, common.HexToAddress(address), spender))
            }
            response, err := caller.Execute(userCalls)
            if err != nil {
                fmt.Println("multicall execute error:", err.Error())
                //return
            }
            msg := ""
            for key, res := range response {
                num := big.NewInt(0).SetBytes(res)
                d := decimal.NewFromBigInt(num, 1)
                d = d.Div(decimal.NewFromInt(1e18))
                msg += fmt.Sprintf("%s %s\n", addresses[key], d.String())
            }
            usdcFile.WriteString(msg)
            fmt.Println(msg)

            // clear
            addresses = make([]string, 0)
        }

        if over {
            fmt.Println("over")
            break
        }
    }

}
