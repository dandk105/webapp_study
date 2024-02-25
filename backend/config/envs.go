package config

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

func setDBConfigEnvs() *dbEnvs {

	// WILL: いずれ共通化して読みやすい様にする
	// 対象の環境変数が存在しているか判断して、存在している場合は何もせず、存在していない時にはデフォルトの値を入力する
	// envsKye = {"ENV_KEY":"DB_USER","DEFAULT_KEY": "default"}
	//
	// for i range envsKye {res, e := os.lookupEnv(i) if !e || res == ""{ res = default["i"] }}
	var a *dbEnvs
	Username := os.Getenv("DB_USER")
	DatabaseName := os.Getenv("DB_NAME")
	DatabasePassword := os.Getenv("DB_PASSWORD")
	DatabaseHost := os.Getenv("DB_HOST")

	// Think: どれか一つでも空だった場合defaultの値を設定する
	// もしかして何か一つでも空だった場合には、デフォルトの値を返却するのではなく、エラーを返却するべきでは？
	if Username == "" || DatabaseName == "" || DatabasePassword == "" || DatabaseHost == "" {
		a = &dbEnvs{dbUser: "default", dbName: "default", dbPassword: "default", dbHost: "localhost"}
	}

	a = &dbEnvs{
		dbUser:     Username,
		dbName:     DatabaseName,
		dbPassword: DatabasePassword,
		dbHost:     DatabaseHost,
	}

	// MEMO: 開発環境の時のみDBに関係する環境変数の値をログに出力する
	if IsDebug() {
		log.Printf("Read DB Config from Envs: %s, %s, %s, %s \n", a.dbUser, a.dbName, a.dbPassword, a.dbHost)
	}
	log.Printf("Successfully read DB Config from Envs\n")

	return a
}

// SetDataSourceNameWithContext はデータベースの接続情報をコンテキストに設定する関数
func SetDataSourceNameWithContext(ctx context.Context) context.Context {
	envs := setDBConfigEnvs()
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s sslmode=disable",
		envs.dbUser, envs.dbName, envs.dbPassword, envs.dbHost)

	if IsDebug() {
		log.Printf("Sucess Create Database Source Name: %s\n", dsn)
	}
	c := context.WithValue(ctx, "DatabaseSourceName", dsn)

	return c
}
