package config

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/r1cebucket/gopkg/pkg/log"
)

type configure struct {
	Logger   logger   `json:"logger"`
	Database database `json:"database"`
	Redis    redis    `json:"redis"`
	Email    email    `json:"email"`
	// TODO add new conf here
}

var conf configure

// user pointer to save space
var Logger *logger
var Database *database
var Redis *redis
var Email *email

// TODO add new conf var here

type logger struct {
	Level string `json:"level"`
}
type database struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type email struct {
}

// TODO add new conf struct here

func Parse(confPath string) {
	if strings.HasSuffix(confPath, ".json") {
		log.Info().Msg("use json config")
		ParseJson(confPath)
	}
}

func ParseJson(confPath string) {
	// open and read config file
	confFile, err := os.Open(confPath)
	if err != nil {
		log.Fatal().Msg("Cannot open config file:" + err.Error())
		return
	}
	defer confFile.Close()
	data, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatal().Msg("Cannot read config file:" + err.Error())
		return
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		return
	}

	Database = &conf.Database
	Logger = &conf.Logger
	Redis = &conf.Redis

	// TODO add new conf here
}