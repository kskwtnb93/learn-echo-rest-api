package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err) // プログラム強制終了。
		}
	}
	// 環境変数からデータベースに接続するためのURLの文字列を作成。
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		// 第二引数以降の引数が上記の%sに当てはめられていく。
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(
		postgres.Open(url),
		&gorm.Config{}, // 空の構造体を渡してデフォルトの設定で起動する。
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
