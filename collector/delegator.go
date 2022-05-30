package collector

import (
	"context"
	"forbole/flow-exporter/types"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/prometheus/client_golang/prometheus"
)

type FlowDelegatorrCollector struct {
	Ctx              context.Context
	FlowClient       *grpc.Client
	DelegatorAddress string

	AccountBalance *prometheus.Desc
}

func NewFlowDelegatorrCollector(ctx context.Context, flowClient *grpc.Client, delegatorAddress string) *FlowDelegatorrCollector {
	return &FlowDelegatorrCollector{
		Ctx:              ctx,
		FlowClient:       flowClient,
		DelegatorAddress: delegatorAddress,

		AccountBalance: prometheus.NewDesc(
			"flow_account_balance",
			"Flow account balance",
			[]string{"delegator_address"},
			types.FLOW_DENOM_LABEL,
		),
	}
}

func (collector *FlowDelegatorrCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.AccountBalance

}

func (collector *FlowDelegatorrCollector) Collect(ch chan<- prometheus.Metric) {
	if account, err := collector.FlowClient.GetAccount(collector.Ctx, flow.HexToAddress(collector.DelegatorAddress)); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.AccountBalance, err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.AccountBalance, prometheus.GaugeValue, types.ConvertToDisplayDenom(account.Balance), collector.DelegatorAddress)
	}

}
