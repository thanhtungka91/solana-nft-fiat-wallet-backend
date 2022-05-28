package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thanhtungka91/solana-nft-fiat-wallet-backend/models"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	db := models.ConnectDatabase()
	fmt.Println(db)
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/balance/:wallet_address", func(c *gin.Context) {
		walletAddress := c.Param("wallet_address")

		result, _ := models.GetBalance(db, walletAddress)

		c.JSON(http.StatusOK, result)

	})

	type Deposit struct {
		WalletAddress string `form:"wallet_address" json:"wallet_address" xml:"wallet_address"  binding:"required"`
		Amount        string `form:"amount" json:"amount" xml:"amount" binding:"required"`
	}

	router.POST("/balance", func(c *gin.Context) {
		var deposit Deposit
		var err error
		if err = c.ShouldBindJSON(&deposit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var s float64

		if s, err = strconv.ParseFloat(deposit.Amount, 64); err == nil {
			fmt.Println(s) // 3.14159265
		}

		result := models.Deposit(db, models.Balance{
			SolWalletAddress: deposit.WalletAddress,
			Balance:          s,
		})

		c.JSON(http.StatusOK, result)
	})

	router.Run(":9000")
}
