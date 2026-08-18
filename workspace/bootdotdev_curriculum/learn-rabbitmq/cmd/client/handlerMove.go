package main

import (
	"fmt"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/pubsub"
	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)


func handlerMove(gs *gamelogic.GameState, ch *amqp.Channel) func(gamelogic.ArmyMove) routing.AckType {
	return func(amv gamelogic.ArmyMove) routing.AckType {
		defer fmt.Print("> ")
		
		mvOutcome := gs.HandleMove(amv)
		switch mvOutcome {
		case gamelogic.MoveOutcomeSamePlayer:
			fallthrough
		case gamelogic.MoveOutComeSafe:
			return routing.Ack
		case gamelogic.MoveOutcomeMakeWar:
			err := pubsub.PublishJSON(
				ch, 
				routing.ExchangePerilTopic, 
				routing.WarRecognitionsPrefix+ "." + gs.GetUsername(), 
				gamelogic.RecognitionOfWar{
				Attacker: amv.Player,
				Defender: gs.GetPlayerSnap(),
				},
			)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return routing.NackRequeue
			}
			return routing.Ack
		}
		return routing.NackDiscard
	}
}