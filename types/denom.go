package types

import "math"

const DISPLAYDENOM = "flow"
const EXPONENT = 8

var FLOW_DENOM_LABEL = map[string]string{
	"denom": "flow",
}

func ConvertToDisplayDenom(n uint64) float64 {
	return float64(float64(n) / math.Pow10(EXPONENT))
}

func ParseRoleByID(n uint8) string {
	switch n {
	case 1:
		return "collection"
	case 2:
		return "consensus"
	case 3:
		return "execution"
	case 4:
		return "verification"
	default:
		return "access"
	}
}
