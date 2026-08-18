# Peril - RabbitMQ Game

A distributed multiplayer game demonstrating RabbitMQ message queuing and pub/sub patterns in Go.

> **Note**: This is a completed project based on Boot.dev's [Learn Pub/Sub](https://learn.boot.dev/learn-pub-sub) course by Osama Nazieh.

## Overview

Peril is a strategy game where multiple players compete by managing units and territories. The game uses RabbitMQ to handle real-time communication between a central server and multiple game clients, showcasing practical applications of message-oriented middleware.

## Features

- **Multi-client architecture**: Support for multiple concurrent players
- **RabbitMQ messaging**: Uses topic and direct exchanges for game state synchronization
- **Real-time game events**: Publish/subscribe pattern for game log broadcasting
- **Player pause mechanism**: Direct messaging to individual players
- **Goroutine-safe state management**: Thread-safe game state using mutexes
- **Game logging**: Centralized server-side logging of all game events

## Project Structure

```
.
├── cmd/
│   ├── client/           # Game client application
│   │   ├── main.go
│   │   ├── handlerMove.go
│   │   ├── handlerPause.go
│   │   └── handlerWar.go
│   └── server/           # Game server application
│       ├── main.go
│       └── handlerLog.go
├── internal/
│   ├── gamelogic/        # Core game mechanics
│   │   ├── gamestate.go
│   │   ├── gamedata.go
│   │   ├── gamelogic.go
│   │   ├── move.go
│   │   ├── pause.go
│   │   ├── spawn.go
│   │   ├── war.go
│   │   └── logs.go
│   ├── pubsub/           # RabbitMQ message operations
│   │   ├── publishGob.go
│   │   ├── publishJSON.go
│   │   ├── subscribeGob.go
│   │   └── subscribeJSON.go
│   └── routing/          # Message routing configuration
│       ├── routing.go
│       └── models.go
├── go.mod
├── go.sum
├── Dockerfile            # Docker configuration for RabbitMQ
├── rabbit.sh             # RabbitMQ server startup script
├── multiserver.sh        # Multi-client server startup script
└── README.md
```

## Prerequisites

- Go 1.22 or later
- RabbitMQ server (can run via Docker)
- amqp091-go library

## Installation

1. Clone the repository:
```bash
cd /home/osama/workspace/bootdotdev_curriculum/learn-rabbitmq
```

2. Install dependencies:
```bash
go mod download
```

## Running the Game

### 1. Start RabbitMQ

Using Docker:
```bash
docker run -d --name rabbitmq -p 5672:15672 rabbitmq:management
```

Or using the provided script:
```bash
bash rabbit.sh
```

### 2. Start the Server

```bash
go run ./cmd/server/main.go
```

The server will connect to RabbitMQ at `amqp://guest:guest@localhost:5672/` and wait for player connections.

### 3. Start Client(s)

In separate terminals, run:
```bash
go run ./cmd/client/main.go
```

Each client will prompt for a username and then connect to the game server.

### Run Multiple Clients

Use the provided script to spawn multiple clients:
```bash
bash multiserver.sh
```

## Architecture

### Message Exchange Pattern

- **Topic Exchange (`peril_topic`)**: Broadcasts game logs and events to all subscribed clients
- **Direct Exchange (`peril_direct`)**: Routes pause/resume messages to specific players

### Queue Types

- **Durable Queues**: Persist game logs for all clients
- **Transient Queues**: Handle temporary player-specific messages

### Serialization

- **Gob encoding**: Used for complex game state objects
- **JSON encoding**: Used for simple messages like pause commands

## Game Commands

The game supports various commands:
- **Move**: Move units to different territories
- **Spawn**: Create new units
- **War**: Attack opponent units
- **Pause**: Pause/resume game state

## Dependencies

- `github.com/rabbitmq/amqp091-go`: RabbitMQ Go client library

## Learning Outcomes

This project demonstrates:
- RabbitMQ pub/sub messaging patterns
- Topic-based and direct exchange routing
- Goroutine-safe concurrent game state management
- Message serialization (Gob and JSON)
- Distributed system design patterns
- Real-time game synchronization

## Configuration

Default RabbitMQ connection: `amqp://guest:guest@localhost:5672/`

Modify the connection string in `cmd/server/main.go` and `cmd/client/main.go` if using different credentials or host.

## License

Part of Boot.dev curriculum materials.
