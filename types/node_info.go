package types

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access"
)

type NodeInfo struct {
	Id                 string
	Role               uint8 // https://github.com/onflow/flow-core-contracts/blob/master/contracts/FlowIDTableStaking.cdc#L112
	TokensStaked       uint64
	TokensCommitted    uint64
	TokensRewarded     uint64
	DelegatorIDCounter uint32
}

func GetNodeInfo(ctx context.Context, client access.Client, nodeID string) (*NodeInfo, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s

	// This script gets all the info about a node and returns it
	pub fun main(nodeID: String): FlowIDTableStaking.NodeInfo {
		return FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	}`, FLOWIDTABLESTAKING)

	args := []cadence.Value{cadence.String(nodeID)}
	value, err := client.ExecuteScriptAtLatestBlock(ctx, []byte(script), args)

	if err != nil {
		return nil, err
	}

	s := value.(cadence.Struct)

	nodeInfo := &NodeInfo{
		Id:                 s.Fields[0].ToGoValue().(string),
		Role:               s.Fields[1].ToGoValue().(uint8),
		TokensStaked:       s.Fields[5].ToGoValue().(uint64),
		TokensCommitted:    s.Fields[6].ToGoValue().(uint64),
		TokensRewarded:     s.Fields[9].ToGoValue().(uint64),
		DelegatorIDCounter: s.Fields[11].ToGoValue().(uint32),
	}
	return nodeInfo, nil
}
