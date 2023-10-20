package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	// Connecting proccess to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("error connecting to rabbitmq: %+v", err)
	}
	defer conn.Close()
	fmt.Println("Connect to rabbitmq")

	// Connnect to RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("error channel %+v", err)
	}
	defer ch.Close()

	// Declare Queue
	q, err := ch.QueueDeclare(
		"testQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("error queue %+v", err)
	}

	// // Send data to RabbitMQ struct
	type Person struct {
		Name     string
		Email    string
		Password string
	}

	ogut := &Person{
		Name:     "FianGumilar",
		Email:    "fiangumilar@gmail.com",
		Password: "rahasia123",
	}

	ogutJson, _ := json.Marshal(ogut)

	// Publish queues to RabbitMQ
	if err = ch.Publish(
		"",     // publish to an exchange
		q.Name, // exchange namerouting to 0 or more queues
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(ogutJson),
		},
	); err != nil {
		fmt.Printf("error publishing queue: %+v", err)
	}
	fmt.Println("Successfully publish message to queue.")
}
