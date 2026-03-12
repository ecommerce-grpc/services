package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv returns the environemnt the application is running.
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// GetTracerBackendOrDefault returns the Tracer backend address or a default value otherwise.
func GetTracerBackendOrDefault() string {
	return getEnvironmentValueOrDefault("TRACER_BACKEND", "http://jaeger-otel.jaeger.svc.cluster.local:14278/api/traces")
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

func getEnvironmentValue(key string) string {
	// TODO: Implement a isEmpty function
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing", key)
	}
	return os.Getenv(key)
}

func getEnvironmentValueOrDefault(key string, defaultValue string) string {
	if os.Getenv(key) == "" && defaultValue != "" {
		return defaultValue
	}
	return os.Getenv(key)
}
