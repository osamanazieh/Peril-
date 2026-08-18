package pubsub

import (
	amqp "github.com/rabbitmq/amqp091-go"
)
type SimpleQueueType string

const (
	DurableQueue SimpleQueueType = "durable"
	TransientQueue SimpleQueueType = "transient"
)

func DeclareAndBind(
	conn *amqp.Connection, 
	exchang,
	queueName,
	key string,
	queueType SimpleQueueType,
)	(*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return &amqp.Channel{}, amqp.Queue{}, err
	}
	var isDurable, autoDelete, isExclusive bool 
	if queueType == DurableQueue {
		isDurable = true

	} else {

		autoDelete = true
		isExclusive = true
	}
	queue, err := ch.QueueDeclare(queueName, isDurable, autoDelete, isExclusive, false, amqp.Table{
		"x-dead-letter-exchange": "peril_dlx",
	})

	if err != nil {
		return &amqp.Channel{}, amqp.Queue{}, err
	}
	
	if err := ch.QueueBind(queueName, key, exchang, false, nil); err != nil {
		return &amqp.Channel{}, amqp.Queue{}, err
	}

	return ch, queue, nil 
}