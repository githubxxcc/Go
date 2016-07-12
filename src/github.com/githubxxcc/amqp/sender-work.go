package amqp

import (
	"os"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("url")
	//error handling
	defer conn.Close()

	ch, err := conn.Channel()
	//error handling
	defer ch.Close()

	q, err := ch.QueueDeclare("",
		false,
		false,
		false,
		false,
		nil)

	msg := os.Args

	err = ch.Publish("",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         []byte(msg),
		})
}
