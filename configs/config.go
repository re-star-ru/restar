package configs

import "time"

type Config struct {
	Host          string
	Port          string
	DiscoveryHost string
	Postgres      string

	ServerTimeout time.Duration
	StorePath     string
}

func NewConfig() Config {
	return Config{
		Host:          "0.0.0.0",
		Port:          "8090",
		DiscoveryHost: "0.0.0.0:9999",
		Postgres:      "postgresql://restar:restar@localhost:5432/restar",
		ServerTimeout: time.Second * 180,
		StorePath:     "http://localhost:8888",
	}
}
