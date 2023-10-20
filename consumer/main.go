package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("consumer application")

	// Connect Proccess to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("error connecting to rabbitmq: %v", err)
	}
	defer conn.Close()

	fmt.Println("succes connect to rabbitmq")

	// Connect to RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("error channel: %+v", err)
	}
	defer ch.Close()

	// Declare queues
	msgs, err := ch.Consume(
		"testQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// Create channel to receive messages
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Receive: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connect to rabbitmq instance")
	fmt.Println("[*] waiting message ...")
	<-forever
}
