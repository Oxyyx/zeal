package models

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := InitDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("Couldn't initialize database.")

		os.Exit(1)
	}

	LoadTestObjects()

	os.Exit(m.Run())
}

func LoadTestObjects() {
	db.Create(&Product{Code: "TestingProduct", Price: 999})
}