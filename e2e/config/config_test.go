package config_test

import (
	"testing"

	config "github.com/marcpires/grpc/ecommerce/e2e/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ConfigTestSuite defines the package test suite.
type ConfigTestSuite struct {
	suite.Suite
	MySQLPassword string
}

// SetUpTest setups the required types before running tests.
func (suite *ConfigTestSuite) SetUpTest() {
	suite.MySQLPassword = "verystrongpassword"

}

// TestGetSecretNotEmpty defines a test suite case for empty strings.
func TestGetSecretNotEmpty(t *testing.T) {
	assert.NotEmpty(t, config.GetSecret("default", "MYSQL_TEST_IMAGE"))
}

// TestConfigTestSuite runs the package test suite.
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
