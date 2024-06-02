package publisher

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/mdiaas/processor/internal/core/dataprovider"
	"github.com/mdiaas/processor/internal/core/entity"
)

type Publisher struct {
	client *pubsub.Client
	topic  string
}

func (p *Publisher) Publish(ctx context.Context, message entity.UploadMessage) error {
	topic := p.client.Topic(p.topic)
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	_, err = result.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
func New(client *pubsub.Client, topic string) dataprovider.UploadProducer {
	return &Publisher{
		topic:  topic,
		client: client,
	}
}
