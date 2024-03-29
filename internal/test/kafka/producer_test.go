package kafka_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/kafka"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	config.Parse("../../configs/conf.json")
	log.Setup("debug")
}

func TestProducer(t *testing.T) {
	p := kafka.NewProducer()
	if p == nil {
		t.Error()
		return
	}

	err := kafka.Produce(p, "test", kafka.PartitionAny, []byte("hello0"))
	if err != nil {
		log.Err(err).Msg("Faile to produce message")
		t.Error()
	}

	p.Close()
}
