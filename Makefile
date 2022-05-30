export GO111MODULE = on
###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
	@echo "building flow_exporter binary..."
	@go build -mod=readonly -o build/flow_exporter ./cmd/flow_exporter
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing flow_exporter binary..."
	@go install -mod=readonly ./cmd/flow_exporter
.PHONY: install