package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func parseLevel(path string) (zerolog.Level, error) {
	env := os.Getenv(path)
	if env == "" {
		return 0, fmt.Errorf("%s is empty", path)
	}

	level, err := zerolog.ParseLevel(env)
	if err != nil {
		return 0, err
	}

	return level, nil
}
