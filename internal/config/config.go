package config

import (
	"log"
	"os"
)

type GameTypeEnum int

const (
	GameHannibal GameTypeEnum = iota
	GameAto
)

type Config struct {
	DatabaseURL string
	ServerAddr  string
	Env         string
	Game        GameTypeEnum
}

func (config *Config) IsDev() bool {
	return config.Env == "dev"
}

func (config *Config) IsProd() bool {
	return config.Env == "prod"
}

func Load() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":2137"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	gameTypeString := os.Getenv("SELECTED_GAME")
	if gameTypeString == "" {
		gameTypeString = "hannibal"
	}

	gameTypeEnum := gameTypeFromString(gameTypeString)

	return Config{
		DatabaseURL: dbURL,
		ServerAddr:  serverAddr,
		Env:         env,
		Game:        gameTypeEnum,
	}
}

func gameTypeFromString(gameTypeString string) GameTypeEnum {
	switch gameTypeString {
	case "hannibal":
		return GameHannibal
	case "ato":
		return GameAto
	default:
		return GameHannibal
	}
}
