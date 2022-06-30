package configs

import (
	"log"
	"strings"

	"github.com/danial-riazati/http-monitoring-server/database"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
)

const (
	Prefix = "HTTP_MONITORING_"
)

type Config struct {
	Listen   string          `koanf:"listen"`
	DataBase database.Config `koanf:"database"`
}

func New() Config {
	var c Config
	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := k.Load(env.Provider(
		Prefix,
		".",
		func(s string) string {
			return strings.ReplaceAll(
				strings.ToLower(strings.TrimPrefix(s, Prefix)),
				"__", ".",
			)
		}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	return c
}
