package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Client interface {
	OpenConnection(url string) error
	Consume() error
}

type client struct {
	conn *kafka.Conn
}

func NewClient() Client {
	return &client{}
}

func (c *client) OpenConnection(url string) error {
	ctx := context.Background()

	conn, err := kafka.DialLeader(ctx, "tcp", url, "", 0)
	if err != nil {
		return fmt.Errorf("kafka connection failed: %v", err)
	}

	c.conn = conn

	return nil
}

func (c *client) Consume() error {
	return nil
}
