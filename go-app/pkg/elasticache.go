package pkg

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
)

func InitElasticSearchClient(ctx context.Context, addresses []string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{Addresses: addresses}
	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to ElasticSearch")
	return client, nil

}
