package main

import (
	"fmt"
	"log"
	"os"

	rmq "github.com/andyollylarkin/rabbit-cli/cmd"
	"github.com/spf13/cobra"
)

var version string = "UNKNOWN"

func main() {
	cmd := cobra.Command{
		Use: "rmqcli", Aliases: []string{"rcli"}, Short: "Run rabbitmq consumer/producer",
		Example: "Usage: rmqcli {consumer|producer} {params}",
		Run:     func(cmd *cobra.Command, args []string) { fmt.Println(cmd.Example); os.Exit(1) },
		Version: version,
	}
	var interactive bool
	cmdConsumer := cobra.Command{
		Use:   "consumer",
		Short: "Run rabbitmq consumer",
		Run: func(cmd *cobra.Command, args []string) {
			c := rmq.Consumer{
				Exchange:    cmd.Flag("exchange").Value.String(),
				RoutingKey:  cmd.Flag("rk").Value.String(),
				Host:        cmd.Flag("host").Value.String(),
				Port:        cmd.Flag("port").Value.String(),
				Username:    cmd.Flag("username").Value.String(),
				Password:    cmd.Flag("password").Value.String(),
				Queue:       cmd.Flag("queue").Value.String(),
				Declare:     cmd.Flag("declare").Value.String(),
				Interactive: interactive,
			}

			err := c.Consume()

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.AddCommand(&cmdConsumer)
	cmdConsumer.PersistentFlags().String("exchange", "", "exchange name")
	cmdConsumer.PersistentFlags().String("rk", "", "routing key")
	cmdConsumer.PersistentFlags().String("queue", "", "queue name")
	cmdConsumer.PersistentFlags().String("declare", "no", "declare queue")
	cmdConsumer.PersistentFlags().String("host", "127.0.0.1", "rabbitmq host addr")
	cmdConsumer.PersistentFlags().String("port", "5672", "rabbitmq port")
	cmdConsumer.PersistentFlags().String("username", "guest", "rabbitmq user")
	cmdConsumer.PersistentFlags().String("password", "guest", "rabbitmq password")
	cmdConsumer.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "interactive mode")

	cmdProducer := cobra.Command{
		Use:   "producer",
		Short: "Run rabbitmq producer",
		Run: func(cmd *cobra.Command, args []string) {
			p := rmq.Producer{
				Exchange:    cmd.Flag("exchange").Value.String(),
				RoutingKey:  cmd.Flag("rk").Value.String(),
				Host:        cmd.Flag("host").Value.String(),
				Port:        cmd.Flag("port").Value.String(),
				Username:    cmd.Flag("username").Value.String(),
				Password:    cmd.Flag("password").Value.String(),
				Interactive: interactive,
			}

			err := p.Produce()

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.AddCommand(&cmdProducer)
	cmdProducer.PersistentFlags().String("exchange", "", "exchange name")
	cmdProducer.PersistentFlags().String("rk", "", "routing key")
	cmdProducer.PersistentFlags().String("host", "127.0.0.1", "rabbitmq host addr")
	cmdProducer.PersistentFlags().String("port", "5672", "rabbitmq port")
	cmdProducer.PersistentFlags().String("username", "guest", "rabbitmq user")
	cmdProducer.PersistentFlags().String("password", "guest", "rabbitmq password")
	cmdProducer.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "interactive mode")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
