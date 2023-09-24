package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

// グローバルにハンドラのマップを定義
var routes = make(map[string]string)

// ルートを追加する関数
// pattern: ルートのパターン
// description: ルートの説明
// handlerFunc: ルートにアクセスした時に呼び出されるハンドラ関数
func addRoute(pattern string, description string, handlerFunc http.HandlerFunc) {
	routes[pattern] = description
	http.HandleFunc(pattern, handlerFunc)
}

func main() {
	// データベースを初期化
	db := initDB()

	// /api/greet へのアクセス時に greetHandler を呼び出す
	addRoute("/api/greet", "Greet API", greetHandler)

	// サーバー起動時にハンドラの一覧をログに表示
	log.Println("Registered routes:")
	for route, desc := range routes {
		log.Printf("%s: %s\n", route, desc)
	}

	log.Println("Starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil)) // 8080ポートでサーバーを起動. エラーが発生した時のみログに出力

}

// dbクライアントを初期化する関数
func initDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	// データベースに接続するための文字列を生成
	d := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost)
	db, err := sql.Open("postgres", d)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

func greetHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// GETメソッドのみを受け入れる
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// データベースからデータを取得する
	rows, err := db.Query("SELECT message FROM greetings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	response := fmt.Sprintf("Hello, %s!", name)
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	w.Write([]byte(response))
}
