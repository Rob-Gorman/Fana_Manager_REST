package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load()
	ErrLog.Printf("unable to load .env file: %v", err)
}

func GetEnvVar(envKey string) string {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		err := EnvVarError(envKey)
		ErrLog.Printf("%v", err)
	}
	return val
}

// return type []interface... expected by Sprintf
// (Does NOT want type []string...)
// UPDATE: fix this eventually
func GetEnvVars(envVars ...string) (result []interface{}) {
	for _, envKey := range envVars {
		val := GetEnvVar(envKey)
		result = append(result, val)
	}
	return result
}
