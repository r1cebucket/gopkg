package pkg_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/kafka"
	"github.com/r1cebucket/gopkg/log"
)

func TestProducer(t *testing.T) {
	p := kafka.NewProducer()
	if p == nil {
		t.Error()
		return
	}

	err := kafka.Produce(p, "test", []byte("hello"))
	if err != nil {
		log.Err(err).Msg("Faile to produce message")
		t.Error()
	}

	p.Close()
}
