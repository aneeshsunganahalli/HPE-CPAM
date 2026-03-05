package internal

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	// "github.com/prometheus/common/model"
)


type PrometheusClient struct {
	api v1.API
}

// To create a new Prometheus client for a particular address
func NewPrometheusClient(address string) (*PrometheusClient, error) {
	client, err := api.NewClient(api.Config{Address: address})
	if err != nil {
		return nil, err
	}

	return &PrometheusClient{api: v1.NewAPI(client)}, nil
}

// Query functions after we decide the key metrics we want from Prometheus
