package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"strconv"

	"github.com/joho/godotenv"
)

var (
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
	AT_DAYS    time.Duration
    JWTSecret  []byte
	HostPort   int
)

func LoadConfig() error {
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading env variables from file, trying anyway...")
    }

	var err error;

	// I hate go's error handling
    DBHost, err = getEnv("DB_HOST");
	if (err != nil) {return err}

    DBUser, err = getEnv("DB_USER");
	if (err != nil) {DBUser = "none"}
   
	if (DBUser != "none") {
		DBPassword, err = getEnv("DB_PASSWORD");
		if (err != nil) {DBPassword = "none"}
	}

	DBName, err = getEnv("DB_NAME");
	if (err != nil) {return err}

	days, err := time.ParseDuration(getEnvDefault("AT_DAYS", "7") + "h");
	if (err != nil) {return err}
	AT_DAYS = 24 * days;
	
	jwtSecretString, err := getEnv("JWT_SECRET");
	if (err != nil) {return err}
	JWTSecret = []byte(jwtSecretString)
	
	port_str, err := getEnv("PORT");
	if (err != nil) {return err}

	HostPort, err = strconv.Atoi(port_str);
	if (err != nil) {return err}

	return nil;
}

func getEnv(key string) (string, error) {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value, nil
    }
	return "", fmt.Errorf("%s env variable is not present", key);
}

func getEnvDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
