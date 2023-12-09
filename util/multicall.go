package util

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "multicall/contracts"
    MultiCall2 "multicall/contracts/MultiCall"
    "strings"
    "time"
)

type Call struct {
    Name     string         `json:"name"`
    Target   common.Address `json:"target"`
    CallData []byte         `json:"call_data"`
}

type CallResponse struct {
    Success    bool   `json:"success"`
    ReturnData []byte `json:"returnData"`
}

var erc20Abi, _ = abi.JSON(strings.NewReader(contracts.ERC20MetaData.ABI))

func (call Call) GetMultiCall() MultiCall2.Multicall2Call {
    return MultiCall2.Multicall2Call{Target: call.Target, CallData: call.CallData}
}

func randomSigner() *bind.TransactOpts {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        panic(err)
    }

    signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
    if err != nil {
        panic(err)
    }

    signer.NoSend = true
    signer.Context = context.Background()
    signer.GasPrice = big.NewInt(0)

    return signer
}

type EthMultiCaller struct {
    Signer          *bind.TransactOpts
    Client          *ethclient.Client
    Abi             abi.ABI
    ContractAddress common.Address
}

func NewMultiCall(rawurl, contract string) (EthMultiCaller, error) {
    // timeout after 10 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    //client, err := ethclient.Dial(rawurl)
    client, err := ethclient.DialContext(ctx, rawurl)
    if err != nil {
        return EthMultiCaller{}, err
    }

    // Load Multicall abi for later use
    mcAbi, err := abi.JSON(strings.NewReader(MultiCall2.MultiCall2ABI))
    if err != nil {
        return EthMultiCaller{}, err
    }

    contractAddress := common.HexToAddress(contract)

    return EthMultiCaller{
        Signer:          randomSigner(),
        Client:          client,
        Abi:             mcAbi,
        ContractAddress: contractAddress,
    }, nil
}

//func NewTVLMultiCall(rawurl, contract string) (EthMultiCaller, error) {
//    // timeout after 10 seconds
//    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//    defer cancel()
//
//    //client, err := ethclient.Dial(rawurl)
//    client, err := ethclient.DialContext(ctx, rawurl)
//    if err != nil {
//        return EthMultiCaller{}, err
//    }
//
//    // Load Multicall abi for later use
//    mcAbi, err := abi.JSON(strings.NewReader(tvl.MantaRouterABI))
//    if err != nil {
//        return EthMultiCaller{}, err
//    }
//
//    contractAddress := common.HexToAddress(contract)
//
//    return EthMultiCaller{
//        Signer:          randomSigner(),
//        Client:          client,
//        Abi:             mcAbi,
//        ContractAddress: contractAddress,
//    }, nil
//}

func (caller *EthMultiCaller) Execute(calls []Call) ([][]byte, error) {
    var responses [][]byte

    var multiCalls = make([]MultiCall2.Multicall2Call, 0, len(calls))

    // Add calls to multicall structure for the contract
    for _, call := range calls {
        multiCalls = append(multiCalls, call.GetMultiCall())
    }

    // Prepare calldata for multicall
    callData, err := caller.Abi.Pack("aggregate", multiCalls)
    if err != nil {
        fmt.Println("pack err:", err)
        return responses, err
    }

    // Perform multicall
    resp, err := caller.Client.CallContract(context.Background(), ethereum.CallMsg{To: &caller.ContractAddress, Data: callData}, nil)
    if err != nil {
        fmt.Println("call contract err:", err)
        return responses, err
    }
    unpackedResp, _ := caller.Abi.Unpack("aggregate", resp)

    a, err := json.Marshal(unpackedResp[1])
    if err != nil {
        fmt.Println("marshal err:", err)
        return responses, err
    }

    err = json.Unmarshal(a, &responses)
    if err != nil {
        fmt.Println("unmarshal err:", err)
        return responses, err
    }

    return responses, nil
}

func (caller *EthMultiCaller) TVLExecute(calls []Call) ([][]byte, error) {
    var responses [][]byte

    var multiCalls = make([]MultiCall2.Multicall2Call, 0, len(calls))

    // Add calls to multicall structure for the contract
    for _, call := range calls {
        multiCalls = append(multiCalls, call.GetMultiCall())
    }

    // Prepare calldata for multicall
    callData, err := caller.Abi.Pack("aggregate", multiCalls)
    if err != nil {
        return responses, err
    }

    // Perform multicall
    resp, err := caller.Client.CallContract(context.Background(), ethereum.CallMsg{To: &caller.ContractAddress, Data: callData}, nil)
    if err != nil {
        return responses, err
    }
    unpackedResp, _ := caller.Abi.Unpack("aggregate", resp)

    a, err := json.Marshal(unpackedResp[1])
    if err != nil {
        return responses, err
    }

    err = json.Unmarshal(a, &responses)
    if err != nil {
        return responses, err
    }

    return responses, nil
}
func GetBalanceCall(name string, tokenAddress common.Address, userAddress common.Address) Call {
    parsedData, err := erc20Abi.Pack("balanceOf", userAddress)
    if err != nil {
        panic(err)
    }
    return Call{
        Target:   tokenAddress,
        CallData: parsedData,
        Name:     name,
    }
}

func GetApproveCall(name string, tokenAddress common.Address, userAddress, spender common.Address) Call {
    parsedData, err := erc20Abi.Pack("allowance", userAddress, spender)
    if err != nil {
        panic(err)
    }
    return Call{
        Target:   tokenAddress,
        CallData: parsedData,
        Name:     name,
    }
}

//func GetMantaCall(name string, userAddress, router common.Address) Call {
//    parsedData, err := TvlRouter.Pack("getTotalValue", userAddress)
//    if err != nil {
//        panic(err)
//    }
//    return Call{
//        Target:   router,
//        CallData: parsedData,
//        Name:     name,
//    }
//}
func GetMethodaCall(name string, target, userAddress common.Address) Call {
    parsedData, err := erc20Abi.Pack("executeUSDC", userAddress)
    if err != nil {
        panic(err)
    }
    return Call{
        Target:   target,
        CallData: parsedData,
        Name:     name,
    }
}

//func (caller *EthMultiCaller) Execute(calls []Call) map[string]CallResponse {
//    var responses []CallResponse
//
//    var multiCalls = make([]MultiCall2.Multicall2Call, 0, len(calls))
//
//    // Add calls to multicall structure for the contract
//    for _, call := range calls {
//        multiCalls = append(multiCalls, call.GetMultiCall())
//    }
//
//    // Prepare calldata for multicall
//    callData, err := caller.Abi.Pack("tryAggregate", false, multiCalls)
//    if err != nil {
//        panic(err)
//    }
//
//    // Perform multicall
//    resp, err := caller.Client.CallContract(context.Background(), ethereum.CallMsg{To: &caller.ContractAddress, Data: callData}, nil)
//    if err != nil {
//        panic(err)
//    }
//
//    // Unpack results
//    unpackedResp, _ := caller.Abi.Unpack("tryAggregate", resp)
//
//    a, err := json.Marshal(unpackedResp[0])
//    if err != nil {
//        panic(err)
//    }
//
//    err = json.Unmarshal(a, &responses)
//    if err != nil {
//        panic(err)
//    }
//
//    // Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
//    results := make(map[string]CallResponse)
//    for i, response := range responses {
//        results[calls[i].Name] = response
//    }
//
//    return results
//}
