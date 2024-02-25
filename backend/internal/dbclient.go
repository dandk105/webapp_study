package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type dbEnvs struct {
	dbUser     string
	dbName     string
	dbPassword string
	dbHost     string
}

func createDBConfigEnvs() *dbEnvs {

	// WILL: いずれ共通化して読みやすい様にする
	// 対象の環境変数が存在しているか判断して、存在している場合は何もせず、存在していない時にはデフォルトの値を入力する
	// envsKye = {"ENV_KEY":"DB_USER","DEFAULT_KEY": "default"}
	//
	// for i range envsKye {res, e := os.lookupEnv(i) if !e || res == ""{ res = default["i"] }}
	d := &dbEnvs{dbUser: "default", dbName: "default", dbPassword: "default", dbHost: "localhost"}
	Username := os.Getenv("DB_USER")
	DatabaseName := os.Getenv("DB_NAME")
	DatabasePassword := os.Getenv("DB_PASSWORD")
	DatabaseHost := os.Getenv("DB_HOST")

	// Think: どれか一つでも空だった場合　という風に変更できないかな
	if Username == "" || DatabaseName == "" || DatabasePassword == "" || DatabaseHost == "" {
		log.Printf("Create Config Struct of Database from Envs \n %s, %s, ***, %s \n", d.dbUser, d.dbName, d.dbHost)
		return d
	}

	a := &dbEnvs{
		dbUser:     Username,
		dbName:     DatabaseName,
		dbPassword: DatabasePassword,
		dbHost:     DatabaseHost,
	}

	// Passwordをそのまま標準出力させない為に、マスキングを行っている。
	// なお、Passwordに空白の値が設定される事はなく、デフォルトの値か環境変数で設定された値のどちらかである
	// TODO: 開発環境の時は、取得した設定の値を全て出力する様に切り替えたい
	log.Printf("Create Config Struct of Database from Envs \n %s, %s, ***, %s \n", a.dbUser, a.dbName, a.dbHost)

	return a
}

// DatabaseSourceConfig DBの設定をまとめている構造体
type DatabaseSourceConfig struct {
	// databaseへの接続を行う名前を保持する
	DBSourceName string
}

// DataSourceNameを作成する関数
// 引数を受け取らずにDatabaseSourceNameが設定された構造体を返却する
func createDataSourceName() *DatabaseSourceConfig {
	// createDBc
	envs := createDBConfigEnvs()
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=*** host=%s sslmode=disable",
		envs.dbUser, envs.dbName, envs.dbHost)

	log.Printf("Set Dsn %s\n", dsn)
	return &DatabaseSourceConfig{DBSourceName: dsn}
}

// Client
// Databaseに関係するログと接続についてまとめている
type Client struct {
	// DBの接続についての設定は基本的にデフォルトで設定されている値が適応される
	DataBaseConnection *sql.DB
	Log                *log.Logger
}

// createConnection Postgresへの接続を確保する為に使用されるメソッド
// 接続が成功した場合には、DatabaseClientの接続を上書きする
// dsn string DataSourceNameの略、Databaseへ接続する際の名前を受け取る
func (client *Client) createConnection(dsn string) error {
	// postgresのみを対象としているためそれ以外のドライバーの事は考慮していない
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed Open DB: %v", err)
		return err
	}
	defer db.Close()
	// DBへの接続をPing通信を行う事で確認している
	pingErr := db.Ping()
	if pingErr == nil {
		log.Print("Success DB Ping Connection")
	} else {
		log.Printf("Failed DB Ping Connection　%v", pingErr)
		return pingErr
	}
	// 処理が最後まで問題なく実行された場合に構造体のDatabaseConnectionを
	// 上書きして新しい状態として処理を行う
	client.DataBaseConnection = db
	return nil
}

func (client *Client) SetDataBaseClientLogger() {
	client.Log = log.New(os.Stdout, "DBClient: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (client *Client) FatalSQLF(err error) {
	client.Log.Fatalf("SQL Error: %v", err)
}

// CreateConnectedDatabaseClient 初期化されたClient構造体を返却する関数
// ここで提供されるClient構造体はDBの接続が確立されていて、かつLoggerが専用に設定されているものである
// DBの接続が何らかの理由で失敗した場合は、Errorを返却する
func CreateConnectedDatabaseClient() (*Client, error) {
	// createDatasourceName()は実質、createDBConfigsも兼ねているわけだが、
	// 少なくともこのコネクション作成済の関数に対して、毎度呼ばれるのは冗長である
	conf := createDataSourceName()
	client := Client{}
	err := client.createConnection(conf.DBSourceName)

	if err != nil {
		return nil, errors.New("Failed Creating of Database Client")
	}

	// TODO:必ず返り値にDBClient構造体を返却する様にする
	// 現在はこの関数内でハンドリングが出来ていないため、依存している関数で何らかの
	// panicが発生した際に処理する事ができず、DatabaseClientが必ず帰ってくる信用がないため
	return &client, nil
}
