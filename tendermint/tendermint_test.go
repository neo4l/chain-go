package tendermint

import (
	"fmt"
	"testing"
)

func Test_Status(t *testing.T) {
	//https://stargate.cosmos.network
	reply, err := Status("http://54.255.215.155:16657")
	fmt.Printf("Status: %s,\nerror: %s", reply, err)
}
