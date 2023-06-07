package main

import (
	"listener/event"
	"listener/rabbitMQ"
	"log"
	"os"
)

func main() {
	//try to connect to RabbitMQ
	connection, err := rabbitMQ.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer connection.Close()

	//start listening for messages
	log.Println("Listening for messages...")

	//create a consumer
	consumer, err := event.NewConsumer(connection)
	if err != nil {
		log.Fatal(err)
	}
	//watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.ERROR", "log.WARNING"})
	if err != nil {
		log.Fatal(err)
	}
}
