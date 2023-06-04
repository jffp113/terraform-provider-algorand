package algorand

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jffp113/terraform-provider-algorand/algorand/client"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"algod_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALGORAND_ALDOG_ENDPOINT", nil),
				Description: "The algorand node endpoint",
			},
			"algod_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALGORAND_ALDOG_TOKEN_ENDPOINT", ""),
				Description: "The algorand node access token",
			},
			"indexer_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALGORAND_INDEXER_ENDPOINT", nil),
				Description: "The indexer node endpoint",
			},
			"indexer_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALGORAND_INDEXER_TOKEN_ENDPOINT", ""),
				Description: "The indexer node access token",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			//TODO random version
			terraformVersion = "1.0+compatible"
		}

		return providerConfiguration(d)
	}

	return p
}

func providerConfiguration(d *schema.ResourceData) (any, diag.Diagnostics) {
	conf := client.Config{
		AlgodEndpoint:   d.Get("algod_endpoint").(string),
		AlgodToken:      d.Get("algod_token").(string),
		IndexerEndpoint: d.Get("indexer_endpoint").(string),
		IndexerToken:    d.Get("indexer_token").(string),
	}

	return client.New(conf)
}
