package rmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	Exchange    string
	Queue       string
	RoutingKey  string
	Declare     string
	Host        string
	Port        string
	Username    string
	Password    string
	Interactive bool
}

func (c Consumer) Consume() error {
	dsn := "amqp://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/"

	conn, err := amqp.Dial(dsn)
	if err != nil {
		return err
	}

	defer conn.Close()

	fmt.Printf("Connected to %s\n", dsn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	if c.Declare == "yes" {
		_, err := ch.QueueDeclare(c.Queue, false, true, false, true, nil)
		if err != nil {
			return err
		}
	}

	err = ch.QueueBind(c.Queue, c.RoutingKey, c.Exchange, false, nil)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(c.Queue, "", true, true, false, false, nil)

	for msg := range msgs {
		fmt.Printf("Received new message. Headers: %v\nMessage: %s\n", msg.Headers, msg.Body)

		if !c.Interactive {
			return nil
		}
	}

	log.Println("Done consuming")

	return nil
}
