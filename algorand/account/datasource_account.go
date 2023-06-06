package account

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jffp113/terraform-provider-algorand/algorand/client"
	"os"
	"time"
)

func DataSourceAlgorandAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataResourceAlgorandAccountRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Local name, not persisted on the blockchain",
			},
			"round": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
			"address": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{"address", "mnemonic_envvar", "mnemonic"},
				Description:  "",
			},
			"mnemonic_envvar": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"address", "mnemonic_envvar", "mnemonic"},
			},
			"mnemonic": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"address", "mnemonic_envvar", "mnemonic"},
				Description:  "",
			},
			"amount": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
			"pending_reward": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
			"amount_without_pending_reward": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
			"reward": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

// TODO improve code readability
func dataResourceAlgorandAccountRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	cli := meta.(*client.Client)

	var address string
	var mnemonic string

	if mn, ok := d.GetOk("mnemonic"); ok {
		mnemonic = mn.(string)
	} else if env, ok := d.GetOk("mnemonic_envvar"); ok {
		v, ok := os.LookupEnv(env.(string))
		if !ok {
			return diag.FromErr(fmt.Errorf("env var (%v) not set", env.(string)))
		}
		mnemonic = v
	} else {
		address = d.Get("address").(string)
	}

	if mnemonic != "" {
		addr, err := getAddress(cli, mnemonic)
		if err != nil {
			return diag.FromErr(err)
		}
		address = addr.address
	}

	acc, err := cli.Accounts.FetchAccountInformation(ctx, address)
	if err != nil {
		return diag.Errorf("Error fetching account (%v): %s", address, err)
	}

	d.SetId(address)
	d.Set("round", acc.Round)
	d.Set("amount", acc.Amount)
	d.Set("pending_reward", acc.PendingReward)
	d.Set("amount_without_pending_reward", acc.AmountWithoutPendingReward)
	d.Set("reward", acc.Reward)
	d.Set("status", acc.Status)
	d.Set("mnemonic", mnemonic)
	d.Set("address", address)

	return nil
}

type res struct {
	address  string
	mnemonic string
}

func getAddress(cli *client.Client, mnemonic string) (struct {
	address  string
	mnemonic string
}, error) {
	res := struct {
		address  string
		mnemonic string
	}{
		mnemonic: mnemonic,
	}

	address, err := cli.Accounts.ConvertMnemonicToAddress(mnemonic)
	res.address = address

	if err != nil {
		return res, fmt.Errorf("invaid mnemonic: %v", err)
	}

	return res, nil
}
