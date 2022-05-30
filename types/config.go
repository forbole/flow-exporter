package types

// Config defines all necessary parameters
type Config struct {
	DelegatorAddress string `mapstructure:"delegator_address"`
	ValidatorAddress string `mapstructure:"validator_address"`
	ExporterPort     string `mapstructure:"exporter_port"`
	GrpcAddress      string `mapstructure:"grpc_address"`
}

// NewConfig builds a new Config instance
func NewConfig(
	delegatorAddress string, validatorAddress string, exporterPort string, grpcAddress string,
) Config {
	return Config{
		DelegatorAddress: delegatorAddress,
		ValidatorAddress: validatorAddress,
		ExporterPort:     exporterPort,
		GrpcAddress:      grpcAddress,
	}
}
