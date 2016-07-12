package amqp

import "github.com/streadway/amqp"

func main() {
	conn, err := amqp.Dial("amqplink")
	//error handling
	defer conn.Close()

	ch, err := conn.Channel()
	//error handling

	q, err := ch.QueueDeclare("",
		false,
		false,
		false, false, nil)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
}
