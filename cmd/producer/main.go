package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	defer producer.Close()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go DeliveryReport(deliveryChan, wg) // async

	for i := 0; i < 100; i++ {
		Publish(fmt.Sprint("transferiu", i+1), "test", producer, []byte("transferecia"), deliveryChan)
	}

	Publish("terminou", "test", producer, []byte("transferecia"), deliveryChan)

	wg.Wait()

	//e := <-deliveryChan
	//msg := e.(*kafka.Message)
	//if msg.TopicPartition.Error != nil {
	//	fmt.Println("Erro ao enviar")
	//} else {
	//	fmt.Println("Mensagem enviada:", msg.TopicPartition)
	//}
	//
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			fmt.Println(string(ev.Value))
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else if string(ev.Value) == "terminou" {
				fmt.Println("TERMINOU")
				return
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
