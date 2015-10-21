package main

import (
	"encoding/json"
	"log"

	"github.com/nothingelsematters7/golang_rabbit/config"
	"github.com/nothingelsematters7/golang_rabbit/utils"
	"github.com/streadway/amqp"
)

// Message provides container for message to be sent and received
type Message struct {
	Type string
	Data map[string]string
}

func main() {
	log.Printf("AMQP URL: %s", config.Conf.AMQPUrl())
	conn, err := amqp.Dial(config.Conf.AMQPUrl())
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	data := map[string]string{
		"order_id": "1",
		"amount":   "10.00"}

	message := Message{"Update", data}
	body, err := json.Marshal(message)
	utils.FailOnError(err, "Failed to dump message to json")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	log.Printf(" [x] Sent %s", body)
	utils.FailOnError(err, "Failed to publish a message")
}
