package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Client interface {
	OpenConnection(url, topic string, partition int) error
	Consume() error
}

type client struct {
	conn *kafka.Conn
}

func NewClient() Client {
	return &client{}
}

func (c *client) OpenConnection(url, topic string, partition int) error {
	ctx := context.Background()

	conn, err := kafka.DialLeader(ctx, "tcp", url, topic, partition)
	if err != nil {
		return fmt.Errorf("kafka connection failed: %v", err)
	}

	c.conn = conn

	return nil
}

func (c *client) Consume() error {
	if err := c.conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return fmt.Errorf("set read deadline failed: %v", err)
	}

	batch := c.conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	bytes := make([]byte, 10e3) // 10KB max per message

	for {
		n, err := batch.Read(bytes)
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}

		log.Println(string(bytes[:n]))
	}
}
