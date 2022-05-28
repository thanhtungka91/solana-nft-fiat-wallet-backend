package models

import (
	"gorm.io/gorm"
	"time"
)

type Balance struct {
	ID               uint           `json:"id" gorm:"primary_key"`
	SolWalletAddress string         `json:"sol_wallet_address"`
	Balance          float64        `json:"balance"`
	ExpiredDate      *time.Time     `json:"expired_date"`
	CreatedAt        *time.Time     `json:"created_at,omitempty"`
	UpdatedAt        *time.Time     `json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt `sql:"index" json:"-"`
}

func GetBalance(db *gorm.DB, add string) (balance *Balance, err error) {
	balance = new(Balance)

	err = db.
		Where("sol_wallet_address = ?", add).
		First(&balance).Error

	return

}
