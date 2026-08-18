package main

import (
	"fmt"

	"github.com/osamaNazieh/Peril/internal/gamelogic"
	"github.com/osamaNazieh/Peril/internal/routing"
)

func logHandler() func(routing.GameLog) routing.AckType {
	return func(gl routing.GameLog) routing.AckType {
		defer fmt.Println(gl.Message)
		if err := gamelogic.WriteLog(gl); err != nil {
			return routing.NackRequeue
		}
		return routing.Ack
	}
}