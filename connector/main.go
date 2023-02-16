package main

import (
	"flag"
	"log"

	"github.com/amirhnajafiz/debezium/connector/internal/kafka"
)

func main() {
	var (
		KafkaURL       = flag.String("url", "kafka:9092", "kafka cluster address")
		KafkaTopic     = flag.String("topic", "kafka.debezium.users", "kafka topic")
		KafkaPartition = flag.Int("partition", 0, "kafka partition")
	)

	// parsing flags
	flag.Parse()

	// create a new kafka client
	c := kafka.NewClient()
	if er := c.OpenConnection(*KafkaURL, *KafkaTopic, *KafkaPartition); er != nil {
		panic(er)
	}

	log.Println("start consuming ...")

	// start consuming
	if er := c.Consume(); er != nil {
		panic(er)
	}
}
