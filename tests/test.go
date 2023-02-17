package tests

import (
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/utils"
	"os"
	"testing"
)

var fixtures *testfixtures.Loader

func FixturesLoad(m *testing.M) {
	application := bootstrap.Boot()
	rootPath := utils.GetRootPath()
	var dialect = ""
	if application.GetConfig().Database[constants.DefaultConnectionDB].Driver == constants.POSTGRES {
		dialect = "postgresql"
	}

	var err error
	fixtures, err = testfixtures.New(
		testfixtures.Template(),
		testfixtures.Database(application.GetActiveConnection().SqlDB()), // You database connection
		testfixtures.Dialect(dialect),                                    // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(rootPath+"/testdata/fixtures"),            // The directory containing the YAML files
		testfixtures.UseDropConstraint(),
	)
	if err != nil {
		application.GetLogger().Fatal(err)
	}
	os.Exit(m.Run())
}

func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
}
