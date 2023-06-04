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
		CreateContext: nil,
		ReadContext:   nil,
		UpdateContext: nil,
		DeleteContext: nil,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"round": {
				Type:        schema.TypeString,
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
	return nil
}

func resourceAlgorandAccountRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceAlgorandAccountUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceAlgorandAccountDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
