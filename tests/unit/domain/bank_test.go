package domain_test

import (
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/tests"
	"testing"
)

func TestX(t *testing.T) {
	t.Parallel()
	tests.PrepareTestDatabase()
	fmt.Println("hay")
	// Your test here ...
}

func TestZ(t *testing.T) {
	tests.PrepareTestDatabase()
	fmt.Println("hay Z")
	// Your test here ...
}
