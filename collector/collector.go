package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "cosmos_watcher"
)

type Exporter struct {
	cosmosRestRpc    string
	cosmosRpcTimeout time.Duration
	logger           log.Logger

	// Metrics
	LatestBlockHeight *prometheus.Desc
}

func New(cosmosRestRpc string, cosmosRpcTimeout time.Duration, logger log.Logger) *Exporter {
	return &Exporter{
		cosmosRestRpc:    cosmosRestRpc,
		cosmosRpcTimeout: cosmosRpcTimeout,
		logger:           logger,
		LatestBlockHeight: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "latest_block_height"),
			"The latest block height of the Cosmos chain.",
			nil,
			nil,
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.LatestBlockHeight
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	client := http.Client{
		Timeout: e.cosmosRpcTimeout,
	}

	resp, err := client.Get(e.cosmosRestRpc + "/blocks/latest")
	if err != nil {
		level.Error(e.logger).Log("msg", "No response from getting latest block request", "err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(e.logger).Log("msg", "Fail to read latest block response body", "err", err)
		return
	}

	var result map[string]map[string]map[string]string
	json.Unmarshal([]byte(body), &result)

	latestBlock, err := strconv.Atoi(result["block"]["header"]["height"])
	if err != nil {
		level.Error(e.logger).Log("msg", "Fail convert latest block string to int", "err", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(e.LatestBlockHeight, prometheus.CounterValue, float64(latestBlock))

	level.Info(e.logger).Log("msg", "The latest block is", "data", latestBlock)
}
