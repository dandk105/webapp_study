package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/greet", greetHandler) // /api/greet へのアクセス時に greetHandler を呼び出す

	http.ListenAndServe(":8080", nil) // 8080ポートでサーバーを起動
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// GETメソッドのみを受け入れる
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	response := fmt.Sprintf("Hello, %s!", name)
	w.Write([]byte(response))
}
