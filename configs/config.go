package configs

type Config struct {
	Host string
}

func NewConfig() Config {
	return Config{
		Host: "0.0.0.0:9999",
	}
}
