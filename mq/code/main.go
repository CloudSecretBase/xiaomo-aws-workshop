package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("rabbitmq consumer")

	conn, err := amqp.Dial("amqps://xiaomo:xiaomo123456@b-998c8a2d-10e8-42de-b22a-8f2a4535ebf8.mq.ap-northeast-1.amazonaws.com")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connected to the rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	fmt.Printf("declared queue %s\n", q.Name)

	msgs, err := ch.Consume(
		"TestQueue",
		"xiaomo",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message is: %s\n", d.Body)
		}
	}()
	<-forever
}
