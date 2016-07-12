package amqp

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqplink")
	FailsOnError(err, "Problems getting connection.")
	defer conn.Close()

	ch, err := conn.Channel()
	FailsOnError(err, "Errs on getting channel.")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)

	FailsOnError(err, "on getting queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)

	forever := make(chan struct{})

	FailsOnError(err, "getting messages")

	go func() {
		for msg := range msgs {
			log.Println(msg.Body)
		}
	}()

	<-forever

}

func FailsOnError(err error, msg string) {

}
