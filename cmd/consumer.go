package rmq

import (
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	Exchange     string
	ExchangeType string
	Queue        string
	RoutingKey   string
	Host         string
	Port         string
	Username     string
	Password     string
}

func (c Consumer) Consume() error {
	if c.Exchange != "" && c.ExchangeType == "" {
		return errors.New("invalid exchange type")
	}

	dsn := "amqp://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/"
	conn, err := amqp.Dial(dsn)

	defer conn.Close()

	if err != nil {
		return err
	}

	fmt.Printf("Connected to %s\n", dsn)

	ch, err := conn.Channel()

	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(c.Exchange, c.ExchangeType, false, true, false, false, nil)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(c.Queue, false, true, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.QueueBind(q.Name, c.RoutingKey, c.Exchange, false, nil)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, true, false, false, nil)

	for msg := range msgs {
		fmt.Printf("Receive new message. Headers: %#v\t. Message: %s\n", msg.Headers, msg.Body)
	}

	log.Println("Done consuming")

	return nil
}
