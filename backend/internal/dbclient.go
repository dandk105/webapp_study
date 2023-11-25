package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DBの初期設定を行う構造体
type DBInitializer struct {
	dsn string
}

func (init *DBInitializer) CreateDataSourceName() {
	// TODO: 環境変数からDBの情報を取得する部分を外部関数に切り出した方が、汎用性が高い
	// Will: Default値を設定するなりして、環境変数が取得できなかった時の挙動を定義する
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost)

	log.Printf("Set Dsn %s\n", dsn)
	init.dsn = dsn
}

// DBの通信を担う構造体
type DBClient struct {
	init         DBInitializer
	DBConneciton *sql.DB
	logger       *log.Logger
}

func (client *DBClient) CreateConnection() {
	client.init.CreateDataSourceName()
	db, err := sql.Open("postgres", client.init.dsn)
	if err != nil {
		log.Printf("Failed Open DB\n %v", err)
		client.DBConneciton = nil
	}
	defer db.Close()
	client.DBConneciton = db
}

func (client *DBClient) SetDBClientLogger() {
	client.logger = log.New(os.Stdout, "DBClient: ", log.Ldate|log.Ltime|log.Lshortfile)
}
