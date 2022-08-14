package utils

import (
	"log"

	"github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {
	err := viper.ReadInConfig()

	Fatal(err, "Error while reading config file")

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
