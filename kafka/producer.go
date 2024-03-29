package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

// ref: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/producer_example/producer_example.go

const PartitionAny int32 = kafka.PartitionAny

func NewProducer() *kafka.Producer {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":   strings.Join(config.Kafka.Servers, ","),
			"api.version.request": "true",
			"message.max.bytes":   1000000,
			"linger.ms":           10,
			"retries":             30,
			"retry.backoff.ms":    1000,
			"acks":                "1",
		})
	if err != nil {
		panic(err)
	}

	// monitor
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Err(ev.TopicPartition.Error).Msg(fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition))
				} else {
					log.Info().Msg(fmt.Sprintf("Delivered message to %v\n", ev.TopicPartition))
				}
			}
		}
	}()

	return p
}

func Produce(p *kafka.Producer, topic string, partition int32, content []byte) error {
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			// Partition: kafka.PartitionAny,
			Partition: partition,
		},
		Value: content,
	}, nil)
	return err
}

func Flush(p *kafka.Producer) {
	unflushed := p.Len()
	for unflushed > 0 {
		unflushed = p.Flush(15 * 1000)
	}
}

func TransFlush(ctx context.Context, p *kafka.Producer, topic string, partition int32, content []byte) error {
	err := p.InitTransactions(ctx)
	if err != nil {
		return err
	}
	err = p.BeginTransaction()
	if err != nil {
		p.AbortTransaction(ctx)
		return err
	}
	/*unflushed := */ p.Flush(15 * 1000) // TODO
	p.CommitTransaction(ctx)
	return nil
}
