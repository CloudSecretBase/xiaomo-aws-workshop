package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("rabbitmq consumer")

	conn, err := amqp.Dial("amqps://xiaomo:Xiaomo19921021.@b-af5e9a19-f27d-4a54-9f9f-72b1d7847ad6.mq.ap-northeast-1.amazonaws.com")
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
