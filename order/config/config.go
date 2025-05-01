package config

import (
	"log"
	"os"
	"strconv"
)

// GetPaymentServiceURL returns the address of the Payment service.
func GetPaymentServiceURL() string {
	return getEnvironmentValue("PAYMENT_SERVICE_URL")
}

// GetEnv returns the environemnt the application is running.
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// GetDataSourceURL the connection URL for database being used.
func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

// GetApplicationPort returns the application port being used.
func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("Port %s is invalid", portStr)
	}
	return port
}

// getEnvironementValue gets the environment variable informed.
func getEnvironmentValue(key string) string {
	// TODO: Implement a isEmpty function
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing", key)
	}
	return os.Getenv(key)
}
