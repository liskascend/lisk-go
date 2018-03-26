package main

import (
	"context"
	"log"

	"github.com/liskascend/lisk-go/api"
	"github.com/liskascend/lisk-go/transactions"
)

func main() {
	transaction := transactions.NewTransaction("16313739661670634666L", 20, "wagon stock borrow episode laundry kitten salute link globe zero feed marble", "", 0)

	client := api.NewClient()
	res, err := client.SendTransaction(context.Background(), transaction)

	log.Printf("%v\n", res.Result.Message)
	log.Print(err)
}
