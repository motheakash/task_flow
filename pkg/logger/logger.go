package logger

import (
	"log"
	"os"
)

func InitLogger(env string) {
	if env == "production" {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(os.Stdout)
	}
}
