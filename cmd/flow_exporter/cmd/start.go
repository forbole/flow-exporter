package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/forbole/flow-exporter/collector"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start exporting flow metrics",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		err := viper.Unmarshal(&config)

		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		flowClient, err := grpc.NewClient(config.GrpcAddress)
		if err != nil {
			panic(err)
		}

		registry := prometheus.NewPedanticRegistry()
		registry.MustRegister(
			collector.NewFlowValidatorCollector(ctx, flowClient, config.ValidatorAddress),
			collector.NewFlowDelegatorrCollector(ctx, flowClient, config.DelegatorAddress),
		)

		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
			ErrorLog:      log.New(os.Stderr, log.Prefix(), log.Flags()),
			ErrorHandling: promhttp.ContinueOnError,
		})

		http.Handle("/metrics", handler)
		log.Fatal(http.ListenAndServe(config.ExporterPort, nil))
		return nil
	},
}
