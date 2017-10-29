package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupTest()
	// call flag.Parse() here if TestMain uses flags
	test := m.Run()
	tearDownTest()
	os.Exit(test)
}

// TODO: create test database and collections
func setupTest() {
}

// TODO: drop test database
func tearDownTest() {

}
