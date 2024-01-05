package internal

import (
	"net/http"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
