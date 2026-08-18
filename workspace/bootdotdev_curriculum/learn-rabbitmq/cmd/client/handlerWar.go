package main

import (
	"fmt"
	"time"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/pubsub"
	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func publish(msg, username string, ch *amqp.Channel) routing.AckType {
	if err := pubsub.PublishGob(
		ch,
		routing.ExchangePerilTopic, 
		routing.GameLogSlug + "." + username, 
		routing.GameLog{
			CurrentTime:  time.Now(),
			Message:	  msg, 
			Username: 	  username,
		},
	); err != nil {
		fmt.Println(err)
		return routing.NackRequeue
	}
	return routing.Ack
}

func handlerWar(gs *gamelogic.GameState, ch *amqp.Channel) func(gamelogic.RecognitionOfWar) routing.AckType {
	return func(row gamelogic.RecognitionOfWar) routing.AckType {
		defer fmt.Print("> ")
		outcome, winner, loser := gs.HandleWar(row)
		switch outcome {
			case gamelogic.WarOutcomeNotInvolved:
				return routing.NackRequeue
			case gamelogic.WarOutcomeNoUnits:
				return routing.NackDiscard
			case gamelogic.WarOutcomeYouWon:
				msg := fmt.Sprintf("%s won a war against %s", winner, loser)
				return publish(msg, gs.GetUsername(), ch)
			case gamelogic.WarOutcomeOpponentWon:
				msg := fmt.Sprintf("%s won a war against %s", winner, loser)
				return publish(msg, gs.GetUsername(), ch)
			case gamelogic.WarOutcomeDraw:
				msg := fmt.Sprintf("A war between %s and %s resulted in a draw", winner, loser)
				return publish(msg, gs.GetUsername(), ch)
			default:
				fmt.Println("Something went wronge")
				return routing.NackDiscard
		}
	}
}