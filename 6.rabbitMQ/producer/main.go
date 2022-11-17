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
	err = channel.ExchangeDeclare("events-1", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	message := amqp.Publishing{
		Body: []byte("{ \"Command\": \"pull\", \"Type\": \"Q\", \"Age\": 12, \"FirstName\": \"Fawzi\", \"MidleName\": null, \"LastName\": \"Linggo\", \"Msg\": \"query\", \"Timestamp\": \"05:19:08.321\" }"),
	}

	for i := 0; i < 1000; i++ {

		// We publish the message to the exahange we created earlier
		err = channel.Publish("events-1", "random-key", false, false, message)

		if err != nil {
			panic("error publishing a message to the queue:" + err.Error())
		}

		// We create a queue named Test
		_, err = channel.QueueDeclare("test-json", true, false, false, false, nil)

		if err != nil {
			panic("error declaring the queue: " + err.Error())
		}

		// We bind the queue to the exchange to send and receive data from the queue
		err = channel.QueueBind("test-json", "#", "events-1", false, nil)

		if err != nil {
			panic("error binding to the queue: " + err.Error())
		}
		fmt.Printf("Success to send message %s ", message.Body)
	}
}
