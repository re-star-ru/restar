package configs

type Config struct {
	Host          string
	DiscoveryHost string
}

func NewConfig() Config {
	return Config{
		Host:          "0.0.0.0:8080",
		DiscoveryHost: "0.0.0.0:9999",
	}
}
