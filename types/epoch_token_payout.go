package types

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access"
)

type EpochTokenPayout = uint64

func GetEpochTokenPayout(ctx context.Context, client access.Client) (EpochTokenPayout, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s

	pub fun main(): UFix64 {
		return FlowIDTableStaking.getEpochTokenPayout()
	}`, FLOWIDTABLESTAKING)

	args := []cadence.Value{}
	value, err := client.ExecuteScriptAtLatestBlock(ctx, []byte(script), args)
	if err != nil {
		return 0, err
	}

	s := value.(cadence.UFix64).ToGoValue().(uint64)

	return s, nil
}
