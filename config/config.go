package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  []byte
	HostPort   int
)

func LoadConfig() error {
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading env variables from file, trying anyway...")
    }

	var err error;

	// gestione degli errori davvero ottimale
	// ( mi taglio il pisello )
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
