package databaseClientConfig

import (
	"context"
	"fmt"
	"log"
	"os"
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
func createDataSourceName() context.Context {
	// createDBc
	envs := createDBConfigEnvs()
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=*** host=%s sslmode=disable",
		envs.dbUser, envs.dbName, envs.dbHost)

	log.Printf("Set Dsn %s\n", dsn)
	// DatabaseSourceNameをContextに渡すのはありな気がする
	// Context.withValue(context.background(),"DatabaseSourceName",dsn)
	c := context.WithValue(context.Background(), "DatabaseSourceName", dsn)

	return c
}
