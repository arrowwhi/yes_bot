package config

import (
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken      string `env:"BOT_TOKEN,required"`
	StickerFileID string `env:"STICKER_FILE_ID,required"`
	LogLevel      string `env:"LOG_LEVEL"`
}

func Load() Config {
	_ = godotenv.Load()

	var cfg Config

	val := reflect.ValueOf(&cfg).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("env")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		key := parts[0]
		required := len(parts) > 1 && parts[1] == "required"

		value, ok := os.LookupEnv(key)
		if !ok || value == "" {
			if required {
				log.Fatalf("%s is not set", key)
			}
			continue
		}

		if field.Type.Kind() == reflect.String {
			val.Field(i).SetString(value)
		}
	}

	return cfg
}
