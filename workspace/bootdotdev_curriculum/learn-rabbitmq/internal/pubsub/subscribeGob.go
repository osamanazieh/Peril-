package pubsub

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"

	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)



func gobDecode[T any](data []byte) (T, error) {
	reader := bytes.NewReader(data)
	gobDecoer := gob.NewDecoder(reader)
	var info T 
	if err := gobDecoer.Decode(&info); err != nil {
		var zero T 
		return zero, err
	}
	return info, nil
}


func SubscribeGob[T any](
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

	err = ch.Qos(10, 0, false)
	if err != nil {
		return fmt.Errorf("could not set QoS: %v", err)
	}
	deliveryCh, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return err 
	}

	go func(){
		for msg := range deliveryCh {
			info, err := gobDecode[T](msg.Body)
			if err != nil {
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