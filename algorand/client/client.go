package client

import "github.com/hashicorp/terraform-plugin-sdk/v2/diag"

type Config struct {
	AlgodEndpoint   string
	AlgodToken      string
	IndexerEndpoint string
	IndexerToken    string
}

type Client struct {
}

func New(_ Config) (Client, diag.Diagnostics) {

	return Client{}, nil
}
