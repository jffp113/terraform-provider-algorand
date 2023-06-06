package account

import (
	"context"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

type Accounter interface {
	CreateAccount() Credentials
	FetchAccountInformation(ctx context.Context, addr string) (Account, error)
	ConvertMnemonicToAddress(m string) (string, error)
}

type Core struct {
	algodClient *algod.Client
}

func New(cli *algod.Client) *Core {
	return &Core{
		algodClient: cli,
	}
}

func (c *Core) CreateAccount() Credentials {
	account := crypto.GenerateAccount()

	m, _ := mnemonic.FromPrivateKey(account.PrivateKey)

	return Credentials{
		Address:  account.Address.String(),
		Mnemonic: m,
	}
}

func (c *Core) FetchAccountInformation(ctx context.Context, addr string) (Account, error) {
	acc, err := c.algodClient.AccountInformation(addr).
		Do(ctx)

	if err != nil {
		return Account{},
			fmt.Errorf("fetching account information: %w", err)
	}

	return Account{
		Round:                      acc.Round,
		Address:                    addr,
		Amount:                     acc.Amount,
		PendingReward:              acc.PendingRewards,
		AmountWithoutPendingReward: acc.AmountWithoutPendingRewards,
		Reward:                     acc.Rewards,
		Status:                     acc.Status,
	}, nil
}

func (c *Core) ConvertMnemonicToAddress(m string) (string, error) {
	priv, err := mnemonic.ToPrivateKey(m)
	if err != nil {
		return "", fmt.Errorf("converting mnemonic from private key: %w", err)
	}

	acc, err := crypto.AccountFromPrivateKey(priv)

	if err != nil {
		return "", fmt.Errorf("importing account from private key: %w", err)
	}

	return acc.Address.String(), nil
}

/**
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
*/
