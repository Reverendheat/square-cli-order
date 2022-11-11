package dotenv

import (
	"os"
	"testing"
)

func TestSetEnvString(t *testing.T) {
	parsed := "TEST_KEY=FunTestsWithRevernedHeat135135"
	want := "FunTestsWithRevernedHeat135135"
	setEnv(parsed)
	found := os.Getenv("TEST_KEY")
	if want != found {
		t.Fatalf("%s wanted, however %s returned.", want, found)
	}
}

func TestSetEnvIncorrectSeperator(t *testing.T) {
	parsed := "TEST_KEY/FunTestsWithRevernedHeat135135"
	want := "FunTestsWithRevernedHeat135135"
	err := setEnv(parsed)
	found := os.Getenv("TEST_KEY")
	if err == nil {
		t.Fatalf("Error not raised for want %s, found %s.", want, found)
	}
}
