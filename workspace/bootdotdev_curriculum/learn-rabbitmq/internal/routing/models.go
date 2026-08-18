package routing

import "time"

type PlayingState struct {
	IsPaused bool
}

type GameLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}

type AckType string
const (
	Ack AckType = "ack"
	NackRequeue AckType = "nackR"
	NackDiscard AckType = "nackD"
)