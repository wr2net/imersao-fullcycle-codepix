package db

import (
	"github.com/wr2net/imersao-fullcylce-codepix/domain/model"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.ENV")

	if err != nil {
		log.Fatalf("Error loading .ENV files")
	}
}

func ConnectDB(ENV string) *gorm.DB {
	var DSN string
	var db *gorm.DB
	var err error

	if ENV != "test" {
		DSN = os.Getenv("DSN")
		db, err = gorm.Open(os.Getenv("DB_TYPE"), DSN)
	} else {
		DSN = os.Getenv("DSN_TEST")
		db, err = gorm.Open(os.Getenv("DB_TYPETest"), DSN)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
