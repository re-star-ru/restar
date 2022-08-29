package configs

type Config struct {
	Host          string
	DiscoveryHost string
	Postgres      string
}

func NewConfig() Config {
	return Config{
		Host:          "0.0.0.0:8080",
		DiscoveryHost: "0.0.0.0:9999",
		Postgres:      "postgresql://restar:restar@localhost:5432/restar",
	}
}
