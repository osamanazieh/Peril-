package pubsub

import (
	"bytes"
	"context"
	"encoding/gob"

	amqp "github.com/rabbitmq/amqp091-go"
)

func gobEncode[T any](info T) ([]byte, error) {
	var buffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&buffer)
	if err := gobEncoder.Encode(info); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil 
}
func PublishGob[T any](
	ch *amqp.Channel, 
	exchage,
	key string, 
	val T,
) error {
	gobData, err := gobEncode(val)	
	if err != nil {
		return err 
	}

	if err := ch.PublishWithContext(context.Background(), exchage, key, false, false, amqp.Publishing{
		ContentType: "application/gob",
		Body: gobData,

	}); err != nil {
		return err 
	}
	return nil 
}