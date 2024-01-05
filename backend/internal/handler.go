package internal

import (
	"github.com/rs/cors"
	"net/http"
)

func CreateHandler() *http.Handler {
	r := setupRouter()

	// CORSの設定をしている部分。AllowsOriginsには許可するオリジンとしてフロントエンドのドメインを指定する
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*"},
		Debug:          true,
	})

	handler := c.Handler(mux)
	return &handler
}
