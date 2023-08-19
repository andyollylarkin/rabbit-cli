package rmq

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type Producer struct {
	Exchange   string
	RoutingKey string
	Host       string
	Port       string
	Username   string
	Password   string
}

func (p Producer) Produce() error {
	dsn := "amqp://" + p.Username + ":" + p.Password + "@" + p.Host + ":" + p.Port + "/"
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

	cl := make(chan *amqp.Error)
	nc := ch.NotifyClose(cl)

	textout := produce()

	for {
		var text string
		select {
		case e := <-nc:
			log.Fatal(e)
		case text = <-textout:
		}

		msg := amqp.Publishing{
			Body: []byte(text),
		}
		err := ch.Publish(p.Exchange, p.RoutingKey, false, false, msg)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func produce() <-chan string {
	reader := bufio.NewReader(os.Stdin)
	out := make(chan string, 0)

	go func() {
		for {
			fmt.Print("-> ")

			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)
			out <- text
		}
	}()

	return out
}
