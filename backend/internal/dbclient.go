package db

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

// TODO: 環境変数を取得できなかった時に、空文字列を変数に入力してしまもうので、
// その後の工程において、ランタイムエラーが発生してしまう。これを防ぐために、何らかのデフォルト値
// を設定する必要がある
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
type Client struct {
	initer             DBInitializer
	DataBaseConnection *sql.DB
	logger             *log.Logger
}

/*
*
 */
func (client *Client) CreateConnection() {
	// CreateDataSourceNameを呼び出す場所は設定関連の作業を行う
	// 所で一括対応してあげた方が締まりが良い気がする
	client.initer.CreateDataSourceName()
	db, err := sql.Open("postgres", client.initer.dsn)
	if err != nil {
		client.logger.Printf("Failed Open DB\n %v", err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr == nil {
		client.logger.Print("Success DB Ping Connection")
	} else {
		client.logger.Print("Failed DB Ping Connection")
	}
	client.DataBaseConnection = db
}

func (client *Client) SetDataBaseClientLogger() {
	client.logger = log.New(os.Stdout, "DBClient: ", log.Ldate|log.Ltime|log.Lshortfile)
}
