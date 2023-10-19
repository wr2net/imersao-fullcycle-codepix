package cmd

import (
	//ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"github.com/wr2net/imersao-fullcylce-codepix/application/grpc"
	//"github.com/wr2net/imersao-fullcylce-codepix/application/kafka"
	"github.com/wr2net/imersao-fullcylce-codepix/infrastructure/db"
	"os"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and a Kafka Consumer",

	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, portNumber)

		//deliveryChan := make(chan ckafka.Event)
		//producer := kafka.NewKafkaProducer()
		//go kafka.DeliveryReport(deliveryChan)
		//kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		//kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 500051, "gRPC Port")
}
