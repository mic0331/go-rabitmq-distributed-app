package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)


const SensorDiscoveryExchange = "SensorDiscovery"
const SensorListQueue = "SensorList"

// GetChannel return a message broker connection and channel
func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		false, // durable bool
		false, // autoDelete bool
		false, // exclusive bool
		false, // noWait bool
		nil)   // args ampq.Table
	failOnError(err, "Failed to declate queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
