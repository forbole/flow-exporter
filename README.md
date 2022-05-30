# flow-exporter
# Getting started
1. `$ make install`
2. `$ mkdir -p $HOME/.flow_exporter`
3. `$ vi $HOME/.flow_exporter/config.yaml`
4. Add following config
```
delegator_address: ""
validator_address: ""
exporter_port: ""
grpc_address: ""
```
5. `$ flow_exporter start --home  $HOME/.flow_exporter` 
