package main

import (
	"fmt"
	"log"
	// "os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderPlacer struct{
	producer *kafka.Producer
	topic string
	deliveryCh chan kafka.Event
}

func NewOrderPlacer(p *kafka.Producer,topic string) *OrderPlacer{
	return &OrderPlacer{
		producer: p,
		topic: topic,
		deliveryCh: make(chan kafka.Event, 100000),
	}
}

func (op *OrderPlacer) placeOrder(orderType string, size int) error {
	// var	{	
		
	format := fmt.Sprintf("%s-%d",orderType, size)
		payload := []byte(format)
	// }:

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &op.topic, 
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	},
		op.deliveryCh,
	)
//delivery chan makes sure that something is produced
	if err!=nil {
		log.Fatal(err)
	}
	<-op.deliveryCh
	fmt.Println("placed order on the queue %s", format)
	return nil
}
func main() { 
	// topic :="HVSE"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id": "foo",
		"acks": "all"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	
	op := NewOrderPlacer(p,"HVSE")

	for i:=0;i<1000;i++{
		if err := op.placeOrder("market",i+1); err != nil{
			log.Fatal(err)
		}
		time.Sleep(time.Second * 3)
		// fmt.
	}
	// deliveryCh := make(chan kafka.Event, 10000)
	// for {
    // 	err = p.Produce(&kafka.Message{
    // 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    // 		Value: []byte("FOO"),
	// 	},
    // 		deliveryCh,
	// )
	
	// //delivery chan makes sure that something is produced
	
	// 	if err!=nil {
	// 		log.Fatal(err)
	// 	}
	// 	<-deliveryCh
	// 	time.Sleep(time.Second*3)
	// }
	// fmt.Printf("%+v\n",p)
}
