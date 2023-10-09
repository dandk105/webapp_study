package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
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

type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	ID       string
	Name     string
	Birthday time.Time
}

func main() {
	// データベースを初期化
	db := initDB()
	defer initDB()

	// /api/greet へのアクセス時に greetHandler を呼び出す
	addRoute("/api/greet", "Greet API", greetHandlerWrapper(db))
	addRoute("/api/status", "Status Check API", statuscheckHandlerWrapper(db))

	// サーバー起動時にハンドラの一覧をログに表示
	log.Println("Registered routes:")
	for route, desc := range routes {
		log.Printf("%s: %s\n", route, desc)
	}

	port := os.Getenv("PORT")

	log.Printf("Starting server on :%s\n", port)
	// 環境変数PORTでサーバーを起動. エラーが発生した時のみログに出力
	l := ":" + port
	log.Fatal(http.ListenAndServe(l, nil))

}

// dbクライアントを初期化する関数
func initDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	// データベースに接続するための文字列を生成
	// TIPS: hostはdocker-compose.ymlで定義したサービス名を指定する必要がある(localhostだと接続できない)
	d := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost)
	db, err := sql.Open("postgres", d)
	if err != nil {
		log.Print(err)
	}

	// データベースに接続できるか確認
	// THINK: DBへのpingが失敗した時のステータスを保持する変数がpingErrなのは
	// ちょっと分かりづらい気がする
	pingErr := db.Ping()
	if pingErr != nil {
		log.Print(pingErr)
	}
	log.Print("Successfully connected!")

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
	// TODO: URLのクエリーによって取得するデータを変えるようにする
	// 現在はUSERSテーブルの全てのデータを取得しているからよろしくない
	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	log.Print(users)

	if err := rows.Err(); err != nil {
		log.Print(err)
	}

	response := fmt.Sprintf("Hello, %s!", name)
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	w.Write([]byte(response))
}

func statuscheckHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	dbErr := db.Ping()
	if dbErr == nil {
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
	}

}

// greetHandlerをhttp.HandlerFuncに変換する
// dbとhttp.HandlerFuncを受け取り、http.HandlerFuncを返す方が
// 結合度が下がって他のハンドラーにも使うことができそう
func greetHandlerWrapper(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		greetHandler(w, r, db)
	}
}

func statuscheckHandlerWrapper(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statuscheckHandler(w, r, db)
	}
}
