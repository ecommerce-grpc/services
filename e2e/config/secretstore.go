package config

type SecretStore interface {
	GetSecret(secretStore string, secretKey string) string
}
