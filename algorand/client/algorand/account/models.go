package account

type Credentials struct {
	Address  string
	Mnemonic string
}

type Account struct {
	Round uint64

	Address string
	Amount  uint64

	PendingReward              uint64
	AmountWithoutPendingReward uint64
	Reward                     uint64

	Status string
}
