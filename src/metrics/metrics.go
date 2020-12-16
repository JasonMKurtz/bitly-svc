package metrics

import (
	"encoding/json"
	"fmt"

	httphelper "../http"
)

type Metrics struct {
	Units     int      `json:"units"`
	Unit      string   `json:"unit"`
	Facet     string   `json:"facet"`
	Reference string   `json:"unit_reference"`
	Metrics   []string `json:"metrics"`
}

func GetMetricsPerUrl(token, bitlink string) Metrics {
	endpoint := fmt.Sprintf("https://api-ssl.bitly.com/v4/bitlinks/%s/countries", bitlink)
	bytes := httphelper.GetBytes(token, endpoint)

	var metrics Metrics
	if err := json.Unmarshal(bytes, &metrics); err != nil {
		panic(err)
	}

	return metrics
}
