package kafka

import (
	"context"

	"github.com/amirhnajafiz/debezium/app/internal/port/kafka"

	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use: "kafka",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	url := "localhost:9092"
	ctx := context.Background()

	// opening kafka connection
	conn, err := kafka.OpenConnection(ctx, url)
	if err != nil {
		panic(err)
	}

	// consume over topic
	if er := kafka.Consume(conn); er != nil {
		panic(er)
	}
}
