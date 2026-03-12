package config

import (
	"log"
	"os"
)

// Deprecated: Not yet Unimplemented
type Config struct {
	secretStore SecretStore
}

// Deprecated: Not yet Unimplemented
func NewConfig(secret SecretStore) *Config {

	return &Config{
		secretStore: secret,
	}
}

// Deprecated: Not yet Unimplemented
// DefaultSecretStore returs the default secret store used.
func DefaultSecretStore() string {
	return "tmp/secrets"
}

// GetEnv returns the environemnt the application is running.
func GetEnv(key string) string {
	return getEnvironmentValue(key)
}

// Deprecated: Not yet Unimplemented
// GetSecret returns a secret from the secret store.
func (c *Config) GetSecret(secretStore string, secretKey string) string {
	if secretStore == "" {
		secretStore = DefaultSecretStore()
	}
	return secretValue(secretStore, secretKey)
}

// Deprecated: Not yet Unimplemented
func secretValue(secretStore string, secretKey string) string {
	return ""
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
