package tendermint

import "github.com/ybbus/jsonrpc"

func Status(url string) (map[string]interface{}, error) {
	rpcClient := jsonrpc.NewClient(url)
	resp, er := rpcClient.Call("status")
	if er != nil {
		return nil, er
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp.Result.(map[string]interface{}), er
}
