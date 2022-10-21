package configs

import (
	"errors"
	"fmt"
	"os"

	"manager/utils"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load()
	utils.HandleErr(err, "Error loading .env file")
}

func GetEnvVar(envKey string) string {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		errmsg := missingEnvVarMsg(envKey)
		utils.HandleErr(errors.New("missing .env key"), errmsg)
	}
	return val
}

// return type []interface... expected by Sprintf
// (Does NOT want type []string...)
func GetEnvVars(envVars ...string) (result []interface{}) {
	for _, envKey := range envVars {
		val := GetEnvVar(envKey)
		result = append(result, val)
	}
	return result
}

func missingEnvVarMsg(variableName string) string {
	return fmt.Sprintf("No %s variable found in environment. Verify .env file.", variableName)
}

func DBConnStr() string {
	variables := GetEnvVars("DB_HOST", "DB_USER", "DB_NAME", "DB_PW", "DB_PORT")
	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		variables...,
	)

	return dbUri
}
