package main

import (
	"fmt"
	"log"
	"os"

	rmq "github.com/andyollylarkin/rabbit-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use: "rmqcli", Aliases: []string{"rcli"}, Short: "Run rabbitmq consumer/producer",
		Example: "rmqcli consumer",
		Run:     func(cmd *cobra.Command, args []string) { fmt.Println(cmd.Example); os.Exit(1) },
	}
	cmdConsumer := cobra.Command{
		Use:   "consumer",
		Short: "Run rabbitmq consumer",
		Run: func(cmd *cobra.Command, args []string) {
			c := rmq.Consumer{
				Exchange:     cmd.Flag("e").Value.String(),
				ExchangeType: cmd.Flag("et").Value.String(),
				RoutingKey:   cmd.Flag("r").Value.String(),
				Host:         cmd.Flag("h").Value.String(),
				Port:         cmd.Flag("p").Value.String(),
				Username:     cmd.Flag("U").Value.String(),
				Password:     cmd.Flag("P").Value.String(),
				Queue:        cmd.Flag("q").Value.String(),
			}

			err := c.Consume()

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.AddCommand(&cmdConsumer)
	cmdConsumer.PersistentFlags().String("e", "", "exchange name")
	cmdConsumer.PersistentFlags().String("et", "", "exchange type")
	cmdConsumer.PersistentFlags().String("r", "", "routing key")
	cmdConsumer.PersistentFlags().String("q", "", "queue")
	cmdConsumer.PersistentFlags().String("h", "127.0.0.1", "rabbitmq host addr")
	cmdConsumer.PersistentFlags().String("p", "5672", "rabbitmq port")
	cmdConsumer.PersistentFlags().String("U", "guest", "rabbitmq user")
	cmdConsumer.PersistentFlags().String("P", "guest", "rabbitmq password")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
