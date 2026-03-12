package e2e_test

import (
	"context"
	"path/filepath"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"

	config "github.com/marcpires/grpc/ecommerce/e2e/config"

	"log"
	"strings"
	"testing"
)

type OrderTestSuite struct {
	suite.Suite
	*mysql.MySQLContainer
}

func down(container *mysql.MySQLContainer) {
	if err := testcontainers.TerminateContainer(container); err != nil {
		log.Printf("Failed to terminate container: %s", err)
	}
}

func setupMySQL() {
	ctx := context.Background()
	mysqlImage := config.GetEnv("MYSQL_TEST_IMAGE")

	mysqlContainer, err := mysql.Run(ctx,
		mysqlImage,
		mysql.WithConfigFile(filepath.Join("testdata", "mysql.cnf")),
		mysql.WithDatabase("order"),
		mysql.WithUsername("root"),
		mysql.WithPassword("verysecretpassword"),
		mysql.WithScripts(filepath.Join("testdata", "testdata.sql")),
	)
	defer down(mysqlContainer)

	if err != nil {
		log.Printf("Failed to start container: %s", err)
		return
	}

}

func (o *OrderTestSuite) SetUpSuite() {
	setupMySQL()

}