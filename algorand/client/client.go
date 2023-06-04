package client

import (
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/jffp113/terraform-provider-algorand/algorand/client/algorand/account"
)

type Config struct {
	AlgodEndpoint   string
	AlgodToken      string
	IndexerEndpoint string
	IndexerToken    string
}

type Client struct {
	Accounts account.Accounter
}

func New(cfg Config) (*Client, diag.Diagnostics) {
	cli, err := algod.MakeClient(cfg.AlgodEndpoint, cfg.AlgodToken)

	if err != nil {
		return &Client{}, diag.Errorf("creating algod client: %s", err)
	}

	return &Client{
		Accounts: account.New(cli),
	}, nil
}
