package main

import (
	"fmt"
	"log"
	"net/http"
)

// グローバルにハンドラのマップを定義
var routes = make(map[string]string)

// ルートを追加する関数
func addRoute(pattern string, description string, handlerFunc http.HandlerFunc) {
	routes[pattern] = description
	http.HandleFunc(pattern, handlerFunc)
}

func main() {

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
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	w.Write([]byte(response))
}
