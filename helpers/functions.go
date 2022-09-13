package helpers

import (
	"fmt"
	"os"
)

func GetEnv(envVarName string) string {
	if os.Getenv(envVarName) != "" {
		return os.Getenv(envVarName)
	}
	panic(fmt.Sprintf("The environment variable named %s must be set!", envVarName))
}
