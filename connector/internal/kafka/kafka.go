package kafka

import (
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
	return nil
}

func (c *client) Consume() error {
	return nil
}
