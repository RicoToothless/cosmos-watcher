# Cosmos watcher

Prometheus exporter for telemetry monitoring cosmos metrics.

This project follows the Prometheus community exporter pattern.

## Building and running

```
git clone https://github.com/stakefish/cosmos-watcher.git
cd cosmos-watcher
make build
./bin/cosmos-watcher <flags>
```

To build the Docker image:

```
docker build -t cosmos-watcher .

# for macOS docker desktop
docker build --platform=linux/amd64 -t cosmos-watcher .
```

### Flags

* `help` Show context-sensitive help (also try --help-long and --help-man).
* `cosmos.rest-rpc` Cosmos REST RPC URL. Default is `http://localhost:1317`.
* `cosmos.rpc-timeout` Cosmos RPC connect timeout. Default is `5s`.
* `web.listen-address` Address to listen on for web interface and telemetry. Default is `:5577`.
* `web.telemetry-path` Path under which to expose metrics. Default is `/metrics`.
* `version` Show application version.
* `log.level` Set logging level: one of `debug`, `info`, `warn`, `error`.
* `log.format` Set the log format: one of `logfmt`, `json`.
* `web.config.file` Configuration file to use TLS and/or basic authentication. The format of the file is described [in the exporter-toolkit repository](https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md).

### Environment Variables

* `COSMOS_REST_RPC` Cosmos REST RPC URL. Default is `http://localhost:1317`.
* `COSMOS_RPC_TIMEOUT` Cosmos RPC connect timeout. Default is `5s`.
* `EXPORTER_WEB_LISTEN_ADDRESS` Address to listen on for web interface and telemetry. Default is `:5577`.
* `EXPORTER_WEB_TELEMETRY_PATH` Path under which to expose metrics. Default is `/metrics`.