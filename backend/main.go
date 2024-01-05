package main

import (
	handlers "github.com/dandk105/webapp_study/backend/usecase"
	"log"
	"net/http"

	engine "github.com/dandk105/webapp_study/backend/internal"
	"github.com/rs/cors"
)

func main() {

	HandlersDescription := map[string]string{
		"users":    "/api/users",
		"dbstatus": "/api/dbstatus",
		"userData": "/api/userdata",
	}

	mux := http.NewServeMux()
	mux.HandleFunc(HandlersDescription["users"], handlers.GetUserDatalistHandler)
	mux.HandleFunc(HandlersDescription["dbstatus"], handlers.DatabaseStatusCheckHandler)
	mux.HandleFunc(HandlersDescription["userData"], handlers.GetUserDataHandler)

	// CORSの設定をしている部分。AllowsOriginsには許可するオリジンとしてフロントエンドのドメインを指定する
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*"},
		Debug:          true,
	})

	handler := c.Handler(mux)

	// サーバー起動時にハンドラの一覧をログに表示
	log.Println("Registered routes:")

	log.Printf("%s: %s\n", HandlersDescription["users"], "Get all users list")
	log.Printf("%s: %s\n", HandlersDescription["status"], "Get server status")
	log.Printf("%s: %s\n", HandlersDescription["userData"], "Get user data")

	server := engine.CreateEngine()
	log.Fatal(server.ListenAndServe())

}
