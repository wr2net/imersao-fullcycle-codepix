package cmd

import (
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		//deliveryChan := make(chan ckafka.Event)
		//database := db.ConnectDB(os.Getenv("env"))
		//producer := kafka.NewKafkaProducer()
		//
		////kafka.Publish("Ola Cosumer", "teste", producer, deliveryChan)
		//go kafka.DeliveryReport(deliveryChan)
		//
		//kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		//kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
