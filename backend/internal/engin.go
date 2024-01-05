package internal

import (
	"log"
	"net/http"
	"os"
)

type serverEnvs struct {
	serverAddr string
}

func newServerEnvConfig() *serverEnvs {
	port, e := os.LookupEnv("PORT")
	if e != false || port == "" {
		port = "5001"
	}
	host, e := os.LookupEnv("HOST")
	if e != false || host == "" {
		host = ":"
	}
	addr := port + host
	log.Printf("Create Server Addr: %s", addr)

	return &serverEnvs{serverAddr: addr}
}

func CreateEngine() *http.Server {
	env := newServerEnvConfig()
	handler := newHandlers()
	return &http.Server{
		Addr:    env.serverAddr,
		Handler: handler,
	}
}
