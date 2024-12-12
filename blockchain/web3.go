package blockchain

import (
	"encoding/json"
	"fmt"
	"log"

	"org.donghyusn.com/chain/collector/constant"
	"org.donghyusn.com/chain/collector/utils"
)

type Web3RpcRequest struct {
	// RpcUrl  string        `json:"rpcUrl"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type Web3RpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Web3 Block Number
func (result *Web3RpcResponse) GetBlockNumber(rpcUrl string) error {
	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["BLOCK_NUMBER"],
		Params:  []interface{}{"latest", true}, // Get the latest block with full details
		ID:      1,
	}

	res, postErr := utils.Post(rpcUrl, request)

	if postErr != nil {
		return postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Response Error: %v", parseErr)
		return parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return fmt.Errorf("%s", response.Error.Message)
	}

	return nil
}
