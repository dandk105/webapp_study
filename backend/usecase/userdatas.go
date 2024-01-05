package usecase

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetUserDatalistHandler(w http.ResponseWriter, r *http.Request) {
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

	// WILL: 配列に格納するのではなく、mapに格納してjson形式で返すようにしたい
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

	response := fmt.Sprintf("users: %s", users)
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	// レスポンスを返す
	w.Write([]byte(response))
}

func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("Error:Json marshal %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		// レスポンスを返す
		w.Write(jsonresponse)
	}
}
