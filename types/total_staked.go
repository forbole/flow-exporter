package types

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access"
)

type TotalStaked map[uint8]uint64

func GetTotalStaked(ctx context.Context, client access.Client) (TotalStaked, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s

	pub fun main(): {UInt8: UFix64} {
		let staked = FlowIDTableStaking.getTotalTokensStakedByNodeType()
	
		return staked
	}`, FLOWIDTABLESTAKING)

	args := []cadence.Value{}
	value, err := client.ExecuteScriptAtLatestBlock(ctx, []byte(script), args)
	if err != nil {
		return nil, err
	}

	s := value.(cadence.Dictionary).ToGoValue().(map[interface{}]interface{})
	totalStaked := make(TotalStaked)
	for keys, value := range s {
		totalStaked[keys.(uint8)] = value.(uint64)
	}

	return totalStaked, nil

}
