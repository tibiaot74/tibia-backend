package helpers

import (
	"fmt"
	"os"
)

func GetEnv(envVarName string) string {
	if envVarName == "" {
  		return os.Getenv(envVarName)
  	}
	panic(fmt.Sprintf("The environment variable named %s must be set!", envVarName))
}
