package main

import (
	"context"
	"log"

	"github.com/sumup/sumup-go"
	"github.com/sumup/sumup-go/merchant"
)

func main() {
	client := sumup.NewClient()

	account, err := client.Merchant.Get(context.Background(), merchant.GetAccountParams{})
	if err != nil {
		log.Printf("[ERROR] get merchant account: %v", err)
		return
	}

	log.Printf("[INFO] merchant code: %s", *account.MerchantProfile.MerchantCode)
}
