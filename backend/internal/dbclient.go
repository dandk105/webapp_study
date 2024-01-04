package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Envs struct {
	DbUser     string
	DbName     string
	DbPassword string
	DbHost     string
}

func NewDBConfigEnvs() *Envs {

	// WILL: いずれ共通化して読みやすい様にする
	// 対象の環境変数が存在しているか判断して、存在している場合は何もせず、存在していない時にはデフォルトの値を入力する
	Username, exists := os.LookupEnv("DB_USER")
	if !exists || Username == "" {
		Username = "default"
	}
	DatabaseName, exists := os.LookupEnv("DB_NAME")
	if !exists || DatabaseName == "" {
		DatabaseName = "default"
	}
	DatabasePassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists || DatabasePassword == "" {
		DatabasePassword = "default"
	}
	DatabaseHost, exists := os.LookupEnv("DB_HOST")
	if !exists || DatabaseHost == "" {
		DatabaseHost = "localhost"
	}

	a := &Envs{
		DbUser:     Username,
		DbName:     DatabaseName,
		DbPassword: DatabasePassword,
		DbHost:     DatabaseHost,
	}

	// Passwordをそのまま標準出力させない為に、マスキングを行っている。
	// なお、Passwordに空白の値が設定される事はなく、デフォルトの値か環境変数で設定された値のどちらかである
	// TODO: 開発環境の時は、取得した設定の値を全て出力する様に切り替えたい
	log.Printf("Create Config Struct of Database from Envs \n %s, %s, ***, %s \n", a.DbUser, a.DbName, a.DbHost)

	return a
}

// DBの設定をまとめている構造体
type DatabaseSourceConfig struct {
	// databaseへの接続を行う名前を保持する
	DBSourceName string
}

// TODO: 環境変数を取得できなかった時に、空文字列を変数に入力してしまもうので、
// その後の工程において、ランタイムエラーが発生してしまう。これを防ぐために、何らかのデフォルト値
// を設定する必要がある

func (config *DatabaseSourceConfig) CreateDataSourceName() *DatabaseSourceConfig {
	envs := NewDBConfigEnvs()

	/*
		これすれば、一々取得する関数を打たなくていい省略にはなるけど、
		もしどれかの環境変数を取得出来なかった時に、デフォルトの値で起動できる様にする
		事まで考えると環境変数を取得出来たかどうかの判断が難しい
		os.ExpandEnv("user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} host=${DB_HOsT} sslmode=disable")
	*/

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", envs.DbUser, envs.DbName, envs.DbPassword, envs.DbHost)

	log.Printf("Set Dsn %s\n", dsn)
	return &DatabaseSourceConfig{DBSourceName: dsn}
}

// DBの通信を担う構造体
type Client struct {
	SourceName *DatabaseSourceConfig
	// DBの接続についての設定は基本的にデフォルトで設定されている値が適応される
	DataBaseConnection *sql.DB
	logger             *log.Logger
}

/*
Postgresへの接続を確保する為に使用されるメソッド

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
