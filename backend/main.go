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

// 予約システムのユーザー情報を格納する構造体
type User struct {
	// IDはUUIDになるから、stringでまとめてしまうのはいささか乱雑な気がする
	ID       string
	Name     string
	Birthday time.Time
}

// 予約システムの部屋情報を格納する構造体
type Room struct {
	ID       string
	Name     string
	Capacity int
}

// 予約システムの予約情報を格納する構造体
type Reservation struct {
	ID        string
	UserID    string
	RoomID    string
	StartTime time.Time
	EndTime   time.Time
}

type DBClient struct {
	DB *sql.DB
}

type UsersResponse struct {
	Users User `json:"users"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func main() {
	// データベースを初期化
	db := initDB()
	defer initDB()

	HandlersDescription := map[string]string{
		"users":    "/api/users",
		"status":   "/api/status",
		"userData": "/api/userdata",
	}

	mux := http.NewServeMux()
	mux.HandleFunc(HandlersDescription["users"], func(w http.ResponseWriter, r *http.Request) {
		getUserslistHandler(w, r, db)
	})
	mux.HandleFunc(HandlersDescription["status"], func(w http.ResponseWriter, r *http.Request) {
		statuscheckHandler(w, r, db)
	})
	mux.HandleFunc(HandlersDescription["userData"], func(w http.ResponseWriter, r *http.Request) {
		getUserDataHandler(w, r, db)
	})

	// CORSの設定をしている部分。AllowsOriginsには許可するオリジンとしてフロントエンドのドメインを指定する
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		Debug:          true,
	})

	handler := c.Handler(mux)

	// サーバー起動時にハンドラの一覧をログに表示
	log.Println("Registered routes:")

	log.Printf("%s: %s\n", HandlersDescription["users"], "Get all users list")
	log.Printf("%s: %s\n", HandlersDescription["status"], "Get server status")
	log.Printf("%s: %s\n", HandlersDescription["userData"], "Get user data")

	port := os.Getenv("PORT")

	log.Printf("Starting server on :%s\n", port)
	// 環境変数PORTでサーバーを起動. エラーが発生した時のみログに出力
	l := ":" + port
	//ここの行を分離しないと、エラーになった時に常にサーバーがOS.Exit(1)で終了してしまう
	log.Fatal(http.ListenAndServe(l, handler))

}

// dbクライアントを初期化する関数
// 環境変数の読み込みとデータベースへの接続を行う
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

	/*
		/Will 何れしたいDB周りのリファクタリングについて記載
					type DBClient struct {
						DB *sql.DB
					}

					定義した構造を初期化
					DBClient := &DBClient{DB: db}
					return DBClient, nil
	*/
	return db
}

func getUserslistHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// GETメソッドを受け入れる
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
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

	// users: {"id":"",...}  というjson形式で返すために加工する必要がある現在は下記の形で返している
	// users: %s[{a938f06c-3f90-4dc2-97df-0f8dd456eba9 Alice 1990-01-01 00:00:00 +0000 +0000} {ed99e003-349e-42b7-ae49-74594c7faa29 Bob 1992-05-15 00:00:00 +0000 +0000} {fc73ce46-bbda-476b-b991-3c4fe63e4af5 Charlie 1988-11-23 00:00:00 +0000 +0000}]

	response := fmt.Sprint("users: %s", users)
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
		}

		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
	}
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

}

func getUserDataHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// GETメソッドを受け入れる
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	// "name"が空の時はデフォルトのユーザー情報を返す所謂疎通確認用の処理
	// switch文で書いた方が良さそう
	if name == "" {
		var user User
		user.Name = "World"
		user.Birthday = time.Now()
		user.ID = "1"
		jsonresponse, err := json.Marshal(user)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonresponse)
	} else {
		// データベースからデータを取得する
		// TODO: URLのクエリーによって取得するデータを変えるようにする
		// 現在はUSERSテーブルの全てのデータを取得しているからよろしくない
		rows, err := db.Query("SELECT * FROM USERS WHERE NAME = $1", name)
		if err != nil {
			log.Print(err)
		}
		defer rows.Close()

		var user User
		for rows.Next() {
			if err := rows.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
				log.Print(err)
			}
		}

		if err := rows.Err(); err != nil {
			log.Print(err)
		}

		jsonresponse, err := json.Marshal(user)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		// レスポンスを返す
		w.Write(jsonresponse)
	}
}
