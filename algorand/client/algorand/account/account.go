package account

type Accounter interface {
	CreateAccount() Credentials
}

type Core struct{}

func (c *Core) CreateAccount() Credentials {

	return Credentials{
		Address:  "",
		Mnemonic: "",
	}
}

func (c *Core) FetchAccountInformation() Account {
	//return
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