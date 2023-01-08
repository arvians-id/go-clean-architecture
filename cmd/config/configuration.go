package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	if filenames != nil {
		err := godotenv.Load(filenames...)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &configImpl{}
}
