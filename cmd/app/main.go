package main

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/controller"
	"github.com/mdiaas/processor/internal/app/controller/listener/dto"
	"github.com/mdiaas/processor/internal/core/usecase/process"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var cfg config.Config
var processUc process.Process

func init() {
	cfg = config.Load()

	processUc = process.New()
}
func main() {
	ctx := context.Background()
	listener(ctx)
	controller := controller.New(&cfg)
	if err := controller.Start(); err != nil {
		log.WithError(err).Error("failed to start application")
	}
}

func listener(ctx context.Context) {
	opt := option.WithCredentialsFile(cfg.Google.Credentials)
	client, err := pubsub.NewClient(ctx, "teak-span-419621", opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	topic := client.Topic(cfg.Topics.FileUpload)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists: %v", err)
	}
	if !ok {
		log.Fatal("topic not exists")
	}
	sub := client.Subscription(cfg.Subscribers.FileUpload)
	ok, err = sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if subscription exists: %v", err)
	}
	if !ok {
		log.Fatalf("subscriber not exists")
	}
	go listen(ctx, sub)

}

func listen(ctx context.Context, sub *pubsub.Subscription) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var m dto.UploadMessage
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Printf("Failed to unmarshal message data: %v", err)
			msg.Nack()
			return
		}
		processUc.Do(ctx, m.FileName)
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}

}
