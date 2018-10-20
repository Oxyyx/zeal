package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

var db *gorm.DB

type Cohort struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func InitDB() error {
	var err error
	db, err = gorm.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_CONFIG"))

	if err != nil {
		return err
	}

	MigrateDB()

	return nil
}

func MigrateDB() {
	db.AutoMigrate(&Product{})
}