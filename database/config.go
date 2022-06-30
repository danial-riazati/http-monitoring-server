package database

import "time"

type Config struct {
	ConnectionString  string        `koanf:"connection_string"`
	ConnectionTimeout time.Duration `koanf:"connection_timeout"`
}
