package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/pubsub"
	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"
	amqpTCPConn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer amqpTCPConn.Close()
	
	fmt.Println("Connection has been set seccessfully...")
	
	
	amqpChannel, err := amqpTCPConn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	
	pubsub.SubscribeGob(amqpTCPConn, routing.ExchangePerilTopic, routing.GameLogSlug, "game_logs.*", pubsub.DurableQueue, logHandler())
		
	fmt.Println("Starting Peril server...")
	gamelogic.PrintServerHelp()
	for true {
		input := gamelogic.GetInput()
		if len(input) == 0 {
			continue
		}
		
		if input[0] == "pause"{
			fmt.Println("Sending Pause message...")
			
			if err := pubsub.PublishJSON(amqpChannel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
				IsPaused: true,
				}); err != nil {
					log.Fatal(err)
			}

		} else if input[0] == "resume" {  
			fmt.Println("Sending resume message...")
			if err := pubsub.PublishJSON(amqpChannel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
				IsPaused: false,
				}); err != nil {
				log.Fatal(err)
			}

		} else if input[0] == "quit" {
			fmt.Println("Qutting the game...")
			break
		} else {
			fmt.Println("Command format is not correct...")
		}	
	}
	
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	if _, ok := <-signalChan; ok {
		fmt.Println("the connection has been closed...")
	}

}
