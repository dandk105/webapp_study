package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	ID       string
	Name     string
	Birthday time.Time
}

type UsersResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func main() {
	// データベースを初期化
	db := initDB()
	defer initDB()

	HandlersDescription := map[string]string{
		"users":  "/api/users",
		"status": "/api/status",
	}

	mux := http.NewServeMux()
	mux.HandleFunc(HandlersDescription["users"], func(w http.ResponseWriter, r *http.Request) {
		getUserslistHandler(w, r, db)
	})
	mux.HandleFunc(HandlersDescription["status"], func(w http.ResponseWriter, r *http.Request) {
		statuscheckHandler(w, r, db)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		Debug:          true,
	})

	handler := c.Handler(mux)

	// サーバー起動時にハンドラの一覧をログに表示
	log.Println("Registered routes:")

	log.Printf("%s: %s\n", HandlersDescription["users"], "Get all users list")
	log.Printf("%s: %s\n", HandlersDescription["status"], "Get server status")

	port := os.Getenv("PORT")

	log.Printf("Starting server on :%s\n", port)
	// 環境変数PORTでサーバーを起動. エラーが発生した時のみログに出力
	l := ":" + port
	//ここの行を分離しないと、エラーになった時に常にサーバーがOS.Exit(1)で終了してしまう
	log.Fatal(http.ListenAndServe(l, handler))

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

func getUserslistHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// GETメソッドを受け入れる
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
		// 取得したユーザーリストの情報を全てUser型のスライスに格納し、usersに代入
		users = append(users, user)
	}

	log.Print(users)

	if err := rows.Err(); err != nil {
		log.Print(err)
	}

	// users: {"id":"",...}  というjson形式で返すために加工する必要がある
	response := "users"
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	// レスポンスを返す
	w.Write([]byte(response))
}

// TODO: handlerを全て別のファイルに分離した方が良さそう
func statuscheckHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	dbErr := db.Ping()
	if dbErr == nil {
		// DBへの接続が成功した時はJSON形式でstatus:OK を返す
		response := StatusResponse{Status: "OK"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
	}
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

}

// getUserslistHandlerをhttp.HandlerFuncに変換する
// dbとhttp.HandlerFuncを受け取り、http.HandlerFuncを返す方が
// 結合度が下がって他のハンドラーにも使うことができそう
// func getUserslistHandlerWrapper(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		getUserslistHandler(w, r, db)
// 	}
// }

// func statuscheckHandlerWrapper(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		statuscheckHandler(w, r, db)
// 	}
// }
