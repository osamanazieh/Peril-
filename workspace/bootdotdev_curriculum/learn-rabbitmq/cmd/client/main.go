package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/pubsub"
	"github.com/osamaNazieh/Peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const connectionString = "amqp://guest:guest@localhost:5672/"
	amqpTCPConn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	
	clientName, err := gamelogic.ClientWelcome()
	gamestate := gamelogic.NewGameState(clientName)

	// Define queue and subcribe for pause exchange
	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, clientName)
	pubsub.SubscribeJSON(
		amqpTCPConn, 
		routing.ExchangePerilDirect,
		queueName,
		routing.PauseKey,
		pubsub.TransientQueue, 
		handlerPause(gamestate),
	)
	

	// Define a publish queue
	moveQueueName := fmt.Sprintf("%s.%s", routing.ArmyMovesPrefix, clientName)
	publsihCh, err := amqpTCPConn.Channel()

	pubsub.SubscribeJSON(
		amqpTCPConn, 
		routing.ExchangePerilTopic, 
		moveQueueName, 
		routing.ArmyMovesPrefix + ".*", 
		pubsub.TransientQueue, 
		handlerMove(gamestate, publsihCh),
	)


	pubsub.SubscribeJSON(
		amqpTCPConn, 
		routing.ExchangePerilTopic, 
		"war", 
		routing.WarRecognitionsPrefix + ".*", 
		pubsub.DurableQueue, 
		handlerWar(gamestate, publsihCh),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Client REPL
	Loop:
	for {
		input := gamelogic.GetInput()
		if len(input) == 0 {
			continue 
		}
		switch(input[0]) {
			case "spawn": 
			if err := gamestate.CommandSpawn(input); err != nil {
				log.Fatal(err)
			}
		case "move":
			mv, err := gamestate.CommandMove(input)
			err = pubsub.PublishJSON(publsihCh, routing.ExchangePerilTopic, routing.ArmyMovesPrefix+ "." + clientName, mv)
			
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s moved to %s successfully...\n", mv.Player.Username, mv.ToLocation)
		case "status":
			gamestate.CommandStatus()
		case "help":
			gamelogic.PrintClientHelp()
		case "spam": 
			if input[1] == "" {
				log.Fatal("A second argument represinting number of spam must be proiveded")
			}
			spamNumbrt, err := strconv.Atoi(input[1])
			if err != nil {
				fmt.Println("the given argument is not a number")
			}
			maliciousLog := gamelogic.GetMaliciousLog()
			for range spamNumbrt {
				pubsub.PublishGob(publsihCh, routing.ExchangePerilTopic, "game_logs." + gamestate.GetUsername(), routing.GameLog{
				Username:    clientName,
				CurrentTime: time.Now(),
				Message:     maliciousLog,
			})
			}
		case "quit":
				break Loop
		default:
			err := fmt.Errorf("%s is not a regietered command", input[0])
			fmt.Println(err)
		} 

	}
	fmt.Println("Starting Peril client...")
	
	
	interruptSigChan := make(chan os.Signal, 1)
	signal.Notify(interruptSigChan, os.Interrupt)
	if _, ok := <-interruptSigChan; ok {
		fmt.Println("the connection has ended...")
		syscall.Exit(0)
	}
	
}