package redis

type Config struct {
	Addr     string `env:"ADDR,notEmpty"`
	Database int    `env:"DATABASE" envDefault:"0"`
}
