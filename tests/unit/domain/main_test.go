package domain_test

import (
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/hendrorahmat/golang-clean-architecture/tests"
	"testing"
)

var fixtures *testfixtures.Loader

func TestMain(m *testing.M) {
	tests.FixturesLoad(m)
}
