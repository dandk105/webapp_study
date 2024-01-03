package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func CheckGetedDBsEnvs(keys) bool {
	// 入力された
	for keys {
		_, looked := os.LookupEnv(key)
		looked != ok
		return false
		else
		return true
	}
	return false
}

// DBの設定をまとめている構造体
// @Deprecated 関数の方が管理が楽なので、構造体は廃止する
type DBConfiger struct {
	// databaseへの接続を行う名前を保持する
	DbSourceName string
}

// TODO: 環境変数を取得できなかった時に、空文字列を変数に入力してしまもうので、
// その後の工程において、ランタイムエラーが発生してしまう。これを防ぐために、何らかのデフォルト値
// を設定する必要がある
func (configer *DBConfiger) CreateDataSourceName() string {
	// TODO: 環境変数からDBの情報を取得する部分を外部関数に切り出した方が、汎用性が高い
	// Will: Default値を設定するなりして、環境変数が取得できなかった時の挙動を定義する
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	GetedDbEnvs := CheckGetedDBsEnvs()

	switch GetedDbEnvs {
	case true:
	case false:
		_ = os.Setenv("DB_USER", "default")
		_ = os.Setenv("DB_NAME", "default")
		_ = os.Setenv("DB_PASSWORD", "default")
		_ = os.Setenv("DB_HOST", "default")
	}

	/*
		これすれば、一々取得する関数を打たなくていい省略にはなるけど、
		もしどれかの環境変数を取得出来なかった時に、デフォルトの値で起動できる様にする
		事まで考えると環境変数を取得出来たかどうかの判断が難しい
		os.ExpandEnv("user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} host=${DB_HOsT} sslmode=disable")
	*/

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost)

	log.Printf("Set Dsn %s\n", dsn)
	return dsn
}

// DBの通信を担う構造体
type Client struct {
	dbConfig DBConfiger
	// DBの接続についての設定は基本的にデフォルトで設定されている値が適応される
	DataBaseConnection *sql.DB
	logger             *log.Logger
}

/*
Postgressへの接続を確保する為に使用されるメソッド

dsn string DataSourceNameの略、Databaseへ接続する際の名前を受け取る
 */
func (client *Client) CreateConnection(dsn string) {
	// postgresのみを対象としているためそれ以外のドライバーの事は考慮していない
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		client.logger.Printf("Failed Open DB\n %v", err)
	}
	defer db.Close()
	// DBへの接続をPing通信を行う事で確認している
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
