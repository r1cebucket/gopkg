package test_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/config"
)

func TestParse(t *testing.T) {
	config.Parse("../configs/conf.json")
}