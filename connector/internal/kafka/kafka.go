package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Client interface {
	OpenConnection(ctx context.Context, url string) error
	Consume() error
}

type client struct {
	conn *kafka.Conn
}
