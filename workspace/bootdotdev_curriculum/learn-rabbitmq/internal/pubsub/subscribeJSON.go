package pubsub

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)



func SubscribeJSON[T any](
	conn *amqp.Connection, 
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, 
	handler func(T) routing.AckType,
) error {
	
	ch, _, err := DeclareAndBind(conn, exchange, queueName, key, queueType)
	if err != nil {
		return err 
	}

	deliveryCh, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return err 
	}

	go func() {
		for msg := range deliveryCh {
			var info T 
			if err := json.Unmarshal(msg.Body, &info); err != nil {
				log.Fatal(err)
			}
			ack := handler(info)
			switch ack {
				case routing.Ack:
					msg.Ack(false)
					fmt.Println("Ack")
				case routing.NackRequeue:
					msg.Nack(false, true)
					fmt.Println("Nack with Requeue")
				case routing.NackDiscard:
					msg.Nack(false, false)
					fmt.Println("Nack with Discard")
			}
		}
	}()
	
	return nil 
}