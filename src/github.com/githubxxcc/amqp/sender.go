package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("somelink")
	FailsOnError(err, "Err in connecting")
	defer conn.Close()

	ch, err := conn.Channel()
	FailsOnError(err, "Err in getting channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"q", false, false, false, false, nil)

	FailsOnError(err, "err in creating queue")

	body := "hellow world"

	err = ch.Publish("", q.Name,
		false,
		false,
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

}

func FailsOnError(err error, msg string) {

}
