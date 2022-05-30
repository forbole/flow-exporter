package types

import (
	"context"
	"fmt"
	"math"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access"
)

type NodeDelegatingRewardCut = float64

func GetRewardCutPercentage(ctx context.Context, client access.Client) (NodeDelegatingRewardCut, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s

	pub fun main(): UFix64 {
		return FlowIDTableStaking.getRewardCutPercentage()
	}`, FLOWIDTABLESTAKING)

	args := []cadence.Value{}
	value, err := client.ExecuteScriptAtLatestBlock(ctx, []byte(script), args)
	if err != nil {
		return 0, err
	}

	s := value.(cadence.UFix64).ToGoValue().(uint64)

	return float64(float64(s) / math.Pow10(8)), nil

}
