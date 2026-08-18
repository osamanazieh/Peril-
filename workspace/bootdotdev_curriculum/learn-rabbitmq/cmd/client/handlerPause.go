package main

import (
	"fmt"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/routing"
)

func handlerPause(gs *gamelogic.GameState) func(routing.PlayingState) routing.AckType {
	return func(ps routing.PlayingState) routing.AckType {
		defer fmt.Print("> ")
		gs.HandlePause(ps)
		return routing.Ack
	}
} 