package configs

import (
	"log"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
)

const (
	Prefix = "HTTP_MONITORING_"
)

type DbConfig struct {
	ConnectionString string        `koanf:"connection_string"`
	Timeout          time.Duration `koanf:"connection_timeout"`
}
type UserConfig struct {
	NoOfUrls int32 `koanf:"no_of_urls"`
}
type CallerConfig struct {
	Sleep time.Duration `koanf:"sleep"`
}
type Config struct {
	Listen    string       `koanf:"listen"`
	DataBase  DbConfig     `koanf:"database"`
	SECRETKEY string       `koanf:"secret_key"`
	Caller    CallerConfig `koanf:"caller"`
	User      UserConfig   `koanf:"user"`
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

var Cfg = New()
