package main

import (
	"context"
	"github.com/dandk105/webapp_study/backend/internal"
)

func main() {
	confctx := context.Background()
	// データベースとか外部クライアントは起動の早いタイミングで初期化はしたい
	db := internal.InitDatabaseClient(confctx)
	// txは実際にデータの永続化を行うタイミングで必要なので、ここで利用したいわけでもない
	// contextに渡して上げると、必要なタイミングでtxを利用できるから良いかも？
	tx := internal.InitDatabaseTx(db, confctx)
	context.WithValue(context.Background(), "tx", tx)
}
