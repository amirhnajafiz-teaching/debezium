package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "henry"
	partition = 0
)

func OpenConnection(ctx context.Context, url string) (*kafka.Conn, error) {
	return kafka.DialLeader(ctx, "tcp", url, topic, partition)
}

func Consume(connection *kafka.Conn) error {
	if err := connection.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return err
	}

	batch := connection.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	bytes := make([]byte, 10e3) // 10KB max per message

	for {
		n, err := batch.Read(bytes)
		if err != nil {
			return err
		}

		log.Println(string(bytes[:n]))
	}
}
