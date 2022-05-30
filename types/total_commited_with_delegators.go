package types

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access"
)

func GetTotalCommittedWithDelegators(ctx context.Context, client access.Client, address string) (uint64, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s

	pub fun main(nodeID: String): UFix64 {
		let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
		return nodeInfo.totalCommittedWithDelegators()
	}`, FLOWIDTABLESTAKING)

	args := []cadence.Value{cadence.String(address)}
	value, err := client.ExecuteScriptAtLatestBlock(ctx, []byte(script), args)
	if err != nil {
		return 0, err
	}

	s := value.(cadence.UFix64).ToGoValue().(uint64)

	return s, nil

}
