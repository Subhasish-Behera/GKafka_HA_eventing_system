# GKafka HA Eventing System

A Go template for building event-driven microservices on top of Kafka.

Gkafka demonstrates a simple order eventing flow with one producer and multiple consumers. A producer publishes order events to a Kafka topic, while separate consumer services read the same stream using different consumer groups. This models how independent services can process the same event stream without blocking each other.

## Components

- `main.go` - Kafka producer that publishes order events.
- `processor/` - consumer service for processing orders.
- `datastream/` - separate consumer service for data/event stream reading.
- `docker-compose.yml` - local Kafka and Zookeeper setup.

## Architecture

```text
Order Producer
     |
     v
Kafka topic: HVSE
     |
     +--> processor consumer group: foo
     |
     +--> datastream consumer group: foo_data
```

Both consumers subscribe to the same topic, but use different consumer groups. This lets each service receive and process the full event stream independently.

## Requirements

- Go 1.21+
- Docker
- Docker Compose

## Start Kafka

```bash
docker compose up -d
```

This starts:

- Zookeeper on `localhost:2181`
- Kafka broker on `localhost:9092`

## Run the Producer

From the repository root:

```bash
go run .
```

The producer publishes `market-*` order events to the `HVSE` topic.

## Run the Processor Consumer

```bash
cd processor
go run .
```

Example output:

```text
processing order: market-1
processing order: market-2
```

## Run the Data Stream Consumer

From the repository root:

```bash
go run ./datastream
```

Example output:

```text
data team reading order: market-1
data team reading order: market-2
```

## Kafka Topic

The code uses the topic:

```text
HVSE
```

If your Kafka setup does not auto-create topics, create it manually:

```bash
docker exec -it broker kafka-topics \
  --bootstrap-server localhost:9092 \
  --create \
  --topic HVSE \
  --partitions 1 \
  --replication-factor 1
```

## Project Layout

```text
.
├── docker-compose.yml      # Kafka + Zookeeper local setup
├── main.go                 # order producer
├── datastream/
│   └── main.go             # data stream consumer
└── processor/
    ├── go.mod
    └── main.go             # order processor consumer
```

## Current Behavior

- Producer connects to `localhost:9092`
- Producer publishes to topic `HVSE`
- Producer waits for Kafka delivery confirmation
- Processor consumer uses group `foo`
- Datastream consumer uses group `foo_data`
- Both consumers poll Kafka continuously and print received events
