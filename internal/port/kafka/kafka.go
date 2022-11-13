package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "sql-kafka"
	partition = 0
)

func OpenConnection(ctx context.Context, url string) (*kafka.Conn, error) {
	return kafka.DialLeader(ctx, "tcp", url, topic, partition)
}
