package models

type Balance struct {
	ID               uint    `json:"id" gorm:"primary_key"`
	SolWalletAddress string  `json:"sol_wallet_address"`
	Balance          float64 `json:"balance"`
}
