package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Model 全テーブル共通のカラムについて、gorm.Modelの定義
type Model struct {
	// レコードの生成日時
	CreatedAt time.Time `gorm:"column:ins_ts;type:timestamp"`

	// レコードの更新日時
	UpdatedAt time.Time `gorm:"column:upd_ts;type:timestamp"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))
	db, err = gorm.Open("postgres", conn)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
