package metrics

/*
	This package requests the country-aggregated statistics for a single bitly link and
	computes the average click rate per country over a 30d timeframe
*/

import (
	"encoding/json"
	"fmt"

	httphelper "../http"
)

type Metric struct {
	Value  string  `json:"value"`
	Clicks float32 `json:"clicks"`
}

type Metrics struct {
	Units     int      `json:"units"`
	Unit      string   `json:"unit"`
	Facet     string   `json:"facet"`
	Reference string   `json:"unit_reference"`
	Metrics   []Metric `json:"metrics"`
}

func getMetricsPerUrl(token, bitlink string) Metrics {
	endpoint := fmt.Sprintf("https://api-ssl.bitly.com/v4/bitlinks/%s/countries", bitlink)
	bytes := httphelper.GetBytes(token, endpoint)

	var metrics Metrics
	if err := json.Unmarshal(bytes, &metrics); err != nil {
		panic(err)
	}

	return metrics
}

type MetricResponse struct {
	Clicks    []Metric `json:"metrics"`
	Reference string   `json:"unit_reference"`
	Link      string   `json:"link"`
}

// This is the only exported function in this package, which does the calculation for average click rates per country for a single bitly link.
func GetAverageClickPerCountry(token, bitlink string) MetricResponse {
	var metrics []Metric

	perUrl := getMetricsPerUrl(token, bitlink).Metrics

	for _, metric := range perUrl {
		metrics = append(metrics, Metric{Value: metric.Value, Clicks: metric.Clicks / 30})
	}

	return MetricResponse{
		Clicks:    metrics,
		Reference: "average",
		Link:      bitlink,
	}
}
