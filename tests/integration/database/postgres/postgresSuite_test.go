package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const MigrationsPath = "file://../../../../migrations"

type PostgresDatabaseSuiteTest struct {
	db  database.IDB
	app bootstrap.IApp
	suite.Suite
}

func (pgDb *PostgresDatabaseSuiteTest) SetupSuite() {
	app := bootstrap.Boot()
	pgDb.db = app.GetActiveConnection()
	pgDb.app = app
}

func (pgDb *PostgresDatabaseSuiteTest) SetupTest() {
	driver, err := postgres.WithInstance(pgDb.db.SqlDB(), &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(MigrationsPath, constants.PostgresDriverName, driver)
	assert.NoError(pgDb.T(), err)

	err = m.Up()
	if err != nil && err == migrate.ErrNoChange {
		return
	} else if err != nil {
		panic(err)
	}
}

func (pgDb *PostgresDatabaseSuiteTest) TearDownTest() {
	driver, err := postgres.WithInstance(pgDb.db.SqlDB(), &postgres.Config{})
	assert.NoError(pgDb.T(), err)
	m, err := migrate.NewWithDatabaseInstance(MigrationsPath, constants.PostgresDriverName, driver)
	assert.NoError(pgDb.T(), err)
	assert.NoError(pgDb.T(), m.Down())
}

func TestPostgresRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &PostgresDatabaseSuiteTest{})
}
