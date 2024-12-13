package kafka

type Config struct {
	Brokers []string `env:"BROKERS,notEmpty" envSeparator:","`
}
