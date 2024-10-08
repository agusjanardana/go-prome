package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

func New(filenames ...string) Config {
	_ = godotenv.Load(filenames...)
	return &config{}
}

type config struct {
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
