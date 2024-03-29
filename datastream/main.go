package main

import (
	"fmt"
	"log"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func main() { 
	topic :="HVSE"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":    "localhost:9092",
			"group.id":             "foo_data",//diffrent group id from other consumer
			"auto.offset.reset":    "smallest"})
		
		if err!=nil {
			log.Fatal(err)
		}
		err = consumer.Subscribe(topic, nil)
		// fmt.Printf("v+/n",err)
		if err!=nil {
			log.Fatal(err)
		}
		for {
			ev := consumer.Poll(100)
			// fmt.Println("+v.n",ev)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("data team reading order: %s\n",string(e.Value))
				// application-specific processing
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}