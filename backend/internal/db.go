package internal

import (
	"context"
	"database/sql"
	"log"
)

func InitDatabaseClient(ctx context.Context) *sql.DB {

	// MEMO: context経由でDataSourceNameを持ってきても、型がanyになってしまうので、型を明示的に指定する必要あり
	db, err := sql.Open("postgres", ctx.Value("DataSourceName").string())
	if err != nil {
		log.Printf("Failed Open DB: %v", err)
	}
	dbErr := db.Ping()
	if dbErr != nil {
		log.Printf("Failed Ping DB: %v", dbErr)
	}
	log.Printf("Successfully connected!")

	return db
}

func InitDatabaseTx(db *sql.DB, ctx context.Context) *sql.Tx {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	return tx
}
