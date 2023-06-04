package account

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jffp113/terraform-provider-algorand/algorand/client"
	"time"
)

func ResourceAlgorandAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAlgorandAccountCreate,
		ReadContext:   resourceAlgorandAccountRead,
		UpdateContext: resourceAlgorandAccountUpdate,
		DeleteContext: resourceAlgorandAccountDelete,
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			"mnemonic": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "",
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

func resourceAlgorandAccountCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	cli := meta.(*client.Client)

	cred := cli.Accounts.CreateAccount()

	d.SetId(cred.Address)
	d.Set("mnemonic", cred.Mnemonic)

	return resourceAlgorandAccountRead(ctx, d, meta)
}

func resourceAlgorandAccountRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	cli := meta.(*client.Client)

	acc, err := cli.Accounts.FetchAccountInformation(ctx, d.Id())
	if err != nil {
		//d.SetId("")
		return diag.Errorf("Error fetching account: %s", err)
	}

	d.Set("round", acc.Round)
	d.Set("amount", acc.Amount)
	d.Set("pending_reward", acc.PendingReward)
	d.Set("amount_without_pending_reward", acc.AmountWithoutPendingReward)
	d.Set("reward", acc.Reward)
	d.Set("status", acc.Status)

	return nil
}

func resourceAlgorandAccountUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	//if d.HasChange("name") {
	//}

	return resourceAlgorandAccountRead(ctx, d, meta)
}

func resourceAlgorandAccountDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//TODO opt out from smart contracts
	//TODO opt out from assets
	//TODO transfer algos to main account
	d.SetId("")
	return nil
}
