package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nats-io/nats.go"
	"gitlab.com/syntropynet/amberdm/sdk/pubsub-go/pubsub"
)

const (
	natsUrl     = "34.107.87.29"
	accessToken = "SAAEPLBNRA56YZTG4XN674JQXJ6L5KKVNUB7XUW5YOFJXKV2PYQ2FJJ4ZU"
	aptosTopic  = "syntropy.aptos.mainnet.tx"
)

// RepublishData receives a message on a given subject and republishes it to another subject.
// It takes a context, the service instance, and the data (message) as input arguments.
func PrintData(ctx context.Context, service *pubsub.NatsService, data []byte) error {
	var aptosTx AptosTransaction
	if err := json.Unmarshal(data, &aptosTx); err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil
	}

	if strings.Compare(aptosTx.Txn.Type, "user_transaction") != 0 {
		return nil
	}
	address := aptosTx.Txn.Sender
	events := aptosTx.Txn.Events

	deposit := false
	withdraw := false
	depositAddr := ""
	depositAmount := ""
	withdrawAmount := ""

	for _, evt := range events {
		t := evt.Type
		depositAddr = evt.GUID.AccountAddress
		if strings.Contains(t, "::coin::DepositEvent") {
			if deposit {
				log.Println("Multiple deposit events found!", depositAmount, depositAddr)
			}
			depositAmount = evt.Data.Amount
			deposit = true
		}
		if strings.Contains(t, "::coin::WithdrawEvent") {
			if withdraw {
				log.Println("Multiple withdraw events found!", withdrawAmount)
			}
			withdrawAmount = evt.Data.Amount
			withdraw = true
		}
	}
	message := fmt.Sprintf("Got msg on: %s, sender address: %s, receiver address: %s, withdraw amount: %s, deposit amount: %s",
		aptosTopic, address, depositAddr, withdrawAmount, depositAmount)

	log.Println(message)

	return nil
}

func main() {
	jwt, _ := pubsub.CreateAppJwt(accessToken)

	// Set user credentials and options for NATS connection.
	opts := []nats.Option{}
	opts = append(opts, nats.UserJWTAndSeed(jwt, accessToken))

	// Connect to the NATS server using the provided options.
	service := pubsub.MustConnect(
		pubsub.Config{
			URI:  natsUrl,
			Opts: opts,
		})
	log.Println("Connected to NATS server.")

	// Create a context with a cancel function to control the cancellation of ongoing operations.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Add a handler function to process messages received on the exampleSubscribeSubject.
	service.AddHandler(aptosTopic, func(data []byte) error {
		return PrintData(ctx, service, data)
	})

	// Set up signal handling to gracefully shut down when receiving SIGINT or SIGTERM signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()
	// Start serving messages and processing them using the registered handler function.
	service.Serve(ctx)
}
