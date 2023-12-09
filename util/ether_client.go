package util

import (
    "context"
    "errors"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "sync"
    "time"

    "github.com/coocood/freecache"
)

var mpEthClient sync.Map
var receiptCache *freecache.Cache

func init() {
    cacheSize := 50 * 1024 * 1024
    receiptCache = freecache.NewCache(cacheSize)
}

func GetReceiptByHashFromCache(ethClient *ethclient.Client, hash string) (*types.Receipt, error) {
    t := time.Now()
    fmt.Println("GetReceiptByHashFromCache in time: %s", t)
    receiptBytes, err := receiptCache.Get([]byte(hash))
    if err == nil {
        if len(receiptBytes) != 0 {
            var retReceipt types.Receipt
            err = retReceipt.UnmarshalBinary(receiptBytes)
            if err != nil {
                fmt.Println("transactionReceipt error:%v", err.Error())
                return nil, err
            }
            fmt.Println("GetReceipt for cache")
            fmt.Println("GetReceiptByHashFromCache elapsed: %s", time.Since(t))
            return &retReceipt, nil
        }
    }
    receipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(hash))
    if err != nil {
        fmt.Println("GetReceiptByHashFromCache error:%v", err.Error())
        return nil, err

    }
    receiptBytes, err = receipt.MarshalBinary()
    if err != nil {
        fmt.Println("GetReceiptByHashFromCache error:%v", err.Error())
        return nil, err
    }
    receiptCache.Set([]byte(hash), receiptBytes, 3600*24)
    fmt.Println("GetReceipt from rpc elapsed: %s", time.Since(t))
    return receipt, nil
}

func GetEtherClientByRpcFromCache(chainId int) *ethclient.Client {
    pEthClient, ok := mpEthClient.Load(chainId)
    if ok {
        pEthClient := pEthClient.(*ethclient.Client)
        pEthClient.ChainID(context.Background())
        chainIdBig, err := pEthClient.ChainID(context.Background())
        if err == nil {
            if chainIdBig.Int64() == int64(chainId) {
                fmt.Println("GetEtherClientByRpc for cache")
                return pEthClient
            }
        }
    }
    return nil
}
func GetEtherClientByRpc(rpcUrl string, chainId int) (client *ethclient.Client, ok bool) {
    pEthClient, ok := mpEthClient.Load(chainId)
    if ok {
        pEthClient := pEthClient.(*ethclient.Client)
        pEthClient.ChainID(context.Background())
        chainIdBig, err := pEthClient.ChainID(context.Background())
        if err == nil {
            if chainIdBig.Int64() == int64(chainId) {
                //fmt.Println("GetEtherClientByRpc for cache")
                return pEthClient, true
            }
        }
    }

    client, err := ethclient.Dial(rpcUrl)
    if err != nil {
        fmt.Println("rpc url[%v] eth client: error:%v", rpcUrl, err.Error())
        return nil, false
    }
    chainIdBig, err := client.ChainID(context.Background())
    if err != nil {
        fmt.Println("rpc url[%v] get chainId error:%v", rpcUrl, err.Error())
        return nil, false
    }
    if chainIdBig.Int64() != int64(chainId) {
        return nil, false
    }
    if err != nil {
        fmt.Println("rpc url[%v] using chainId get client error:%v", rpcUrl, err.Error())
        return nil, false
    }

    //fmt.Println("get ether client chainId = %d  at %v success", chainId, rpcUrl)

    mpEthClient.Store(chainId, client)
    return client, true

}

func GetEtherClientByRpcWithNoCache(rpcUrl string, chainId int) (client *ethclient.Client, ok bool, err error) {

    client, err = ethclient.Dial(rpcUrl)
    if err != nil {
        fmt.Println("rpc url[%v] eth client: error:%v", rpcUrl, err.Error())
        return nil, false, err
    }
    chainIdBig, err := client.ChainID(context.Background())
    if err != nil {
        fmt.Println("rpc url[%v] get chainId error:%v", rpcUrl, err.Error())
        return nil, false, err
    }
    if chainIdBig.Int64() != int64(chainId) {
        return nil, false, errors.New("chain id is not equal")
    }
    if err != nil {
        fmt.Println("rpc url[%v] using chainId get client error:%v", rpcUrl, err.Error())
        return nil, false, err
    }
    return client, true, nil

}
