package server

import (
	handlers "github.com/dandk105/webapp_study/backend/usecase"
	"net/http"
)

var handlersDescription = map[string]string{
	"users":    "/api/users",
	"dbstatus": "/api/dbstatus",
	"userData": "/api/userdata",
}

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc(handlersDescription["users"], handlers.GetUserDatalistHandler)
	mux.HandleFunc(handlersDescription["dbstatus"], handlers.DatabaseStatusCheckHandler)
	mux.HandleFunc(handlersDescription["userData"], handlers.DatabaseStatusCheckHandler)

	return mux
}
