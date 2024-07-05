package internal

import (
	"fmt"

	env "github.com/joho/godotenv"
)

func LoadEnv() error {
	err := env.Load(".env")
	if err != nil {
		return fmt.Errorf("error in loading env: %v", err);
	}
	return nil;
}