package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	url := "amqp://guest:guest@developer.alldataint.com:5672"

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the
	// connection itself
	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	// We create an exahange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("from-kafka", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	message, err := channel.Consume("from-confluent-1", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range message {
		//print the message to the console
		fmt.Println("message received: " + string(msg.Body))
		// Acknowledge that we have received the message so it can be removed from the queue
		msg.Ack(true)
	}
	// We close the connection after the operation has completed.
	defer connection.Close()
}
