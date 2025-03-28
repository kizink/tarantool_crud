package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DBconfig
}

type DBconfig struct {
	ADDRESS    string
	SPACE_NAME string
	// USER     string
	// PASSWORD string
}

const loadConfigErr = "LoadConfig failed: error loading .env file"

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic(loadConfigErr)
	}
	return &Config{
		Db: DBconfig{
			ADDRESS:    os.Getenv("ADDRESS"),
			SPACE_NAME: os.Getenv("SPACE_NAME"),
			//USER:     os.Getenv("USER"),
			//PASSWORD: os.Getenv("PASSWORD"),
		},
	}
}
