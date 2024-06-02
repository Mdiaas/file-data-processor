package listener

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/mdiaas/processor/internal/app/config"
	"github.com/mdiaas/processor/internal/app/controller/listener/dto"
	"github.com/mdiaas/processor/internal/core/usecase/process"
)

type listener struct {
	client    *pubsub.Client
	cfg       *config.Config
	processUc process.Process
}

func (l *listener) Listen(ctx context.Context) {
	topic := l.client.Topic(l.cfg.Topics.FileUpload)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists: %v", err)
	}
	if !ok {
		log.Fatal("topic not exists")
	}
	sub := l.client.Subscription(l.cfg.Subscribers.FileUpload)
	ok, err = sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if subscription exists: %v", err)
	}
	if !ok {
		log.Fatalf("subscriber not exists")
	}
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var m dto.UploadMessage
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Printf("Failed to unmarshal message data: %v", err)
			msg.Nack()
			return
		}
		l.processUc.Do(ctx, m.FileName)
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}
}
func New(client *pubsub.Client, cfg *config.Config, processUc process.Process) *listener {
	return &listener{
		client:    client,
		cfg:       cfg,
		processUc: processUc,
	}
}
