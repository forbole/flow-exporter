package collector

import (
	"context"
	"forbole/flow-exporter/types"

	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/prometheus/client_golang/prometheus"
)

type FlowValidatorCollector struct {
	Ctx              context.Context
	FlowClient       *grpc.Client
	ValidatorAddress string

	NetworkStakedTotal      *prometheus.Desc
	NetworkEpochTokenPayout *prometheus.Desc
	ValidatorCommission     *prometheus.Desc
	ValidatorDelegatorCount *prometheus.Desc
	ValidatorStakedTotal    *prometheus.Desc
	ValidatorRewardedTotal  *prometheus.Desc
	ValidatorCommittedTotal *prometheus.Desc
}

func NewFlowValidatorCollector(ctx context.Context, flowClient *grpc.Client, validatorAddress string) *FlowValidatorCollector {
	return &FlowValidatorCollector{
		Ctx:              ctx,
		FlowClient:       flowClient,
		ValidatorAddress: validatorAddress,

		NetworkStakedTotal: prometheus.NewDesc(
			"flow_staked_total",
			"Total activated staked of the network",
			[]string{"role"},
			types.FLOW_DENOM_LABEL,
		),
		NetworkEpochTokenPayout: prometheus.NewDesc(
			"flow_network_epoch_token_payout_total",
			"Total token payout of the network last epoch",
			nil,
			types.FLOW_DENOM_LABEL,
		),
		ValidatorCommission: prometheus.NewDesc(
			"flow_validator_commission_rate",
			"Commission rate of the validator",
			[]string{"validator_address"},
			nil,
		),
		ValidatorDelegatorCount: prometheus.NewDesc(
			"flow_validator_delegators_count",
			"Number of delegators per validator",
			[]string{"validator_address"},
			nil,
		),
		ValidatorStakedTotal: prometheus.NewDesc(
			"flow_validator_token_staked",
			"Token staked per validator",
			[]string{"validator_address", "role"},
			types.FLOW_DENOM_LABEL,
		),
		ValidatorCommittedTotal: prometheus.NewDesc(
			"flow_validator_token_committed",
			"Token committed per validator",
			[]string{"validator_address"},
			types.FLOW_DENOM_LABEL,
		),
		ValidatorRewardedTotal: prometheus.NewDesc(
			"flow_validator_token_rewarded",
			"Token reward per validator",
			[]string{"validator_address"},
			types.FLOW_DENOM_LABEL,
		),
	}
}

func (collector *FlowValidatorCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.NetworkStakedTotal
	ch <- collector.NetworkEpochTokenPayout
	ch <- collector.ValidatorCommission
	ch <- collector.ValidatorCommittedTotal
	ch <- collector.ValidatorDelegatorCount
	ch <- collector.ValidatorStakedTotal
	ch <- collector.ValidatorRewardedTotal
}

func (collector *FlowValidatorCollector) Collect(ch chan<- prometheus.Metric) {
	if totalStakedByRole, err := types.GetTotalStaked(collector.Ctx, collector.FlowClient); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.NetworkStakedTotal, err)
	} else {
		for role, totalStaked := range totalStakedByRole {
			ch <- prometheus.MustNewConstMetric(collector.NetworkStakedTotal, prometheus.GaugeValue, types.ConvertToDisplayDenom(totalStaked), types.ParseRoleByID(role))
		}
	}

	if nodeInfo, err := types.GetNodeInfo(collector.Ctx, collector.FlowClient, collector.ValidatorAddress); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.ValidatorDelegatorCount, err)
		ch <- prometheus.NewInvalidMetric(collector.ValidatorStakedTotal, err)
		ch <- prometheus.NewInvalidMetric(collector.ValidatorRewardedTotal, err)

	} else {
		ch <- prometheus.MustNewConstMetric(collector.ValidatorDelegatorCount, prometheus.GaugeValue, float64(nodeInfo.DelegatorIDCounter), collector.ValidatorAddress)
		ch <- prometheus.MustNewConstMetric(collector.ValidatorStakedTotal, prometheus.GaugeValue, types.ConvertToDisplayDenom(nodeInfo.TokensStaked), collector.ValidatorAddress, types.ParseRoleByID(nodeInfo.Role))
		ch <- prometheus.MustNewConstMetric(collector.ValidatorRewardedTotal, prometheus.GaugeValue, types.ConvertToDisplayDenom(nodeInfo.TokensRewarded), collector.ValidatorAddress)
	}

	if commission, err := types.GetRewardCutPercentage(collector.Ctx, collector.FlowClient); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.ValidatorCommission, err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.ValidatorCommission, prometheus.GaugeValue, float64(commission), collector.ValidatorAddress)
	}

	if epochPayout, err := types.GetEpochTokenPayout(collector.Ctx, collector.FlowClient); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.NetworkEpochTokenPayout, err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.NetworkEpochTokenPayout, prometheus.GaugeValue, types.ConvertToDisplayDenom(epochPayout))
	}

	if totalTokenCommitted, err := types.GetTotalCommittedWithDelegators(collector.Ctx, collector.FlowClient, collector.ValidatorAddress); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.ValidatorCommittedTotal, err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.ValidatorCommittedTotal, prometheus.GaugeValue, types.ConvertToDisplayDenom(totalTokenCommitted), collector.ValidatorAddress)
	}
}
