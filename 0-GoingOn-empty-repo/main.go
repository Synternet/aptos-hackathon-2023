package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SyntropyNet/pubsub-go/pubsub-go/pubsub"
	"github.com/nats-io/nats.go"
)

const (
	natsUrl     = "34.107.87.29"
	accessToken = "SAAEPLBNRA56YZTG4XN674JQXJ6L5KKVNUB7XUW5YOFJXKV2PYQ2FJJ4ZU"
)

var bot *tgbotapi.BotAPI

// RepublishData receives a message on a given subject and republishes it to another subject.
// It takes a context, the service instance, and the data (message) as input arguments.
func PrintData(ctx context.Context, service *pubsub.NatsService, data []byte) error {

	message := "hello"
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
