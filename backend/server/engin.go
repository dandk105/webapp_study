package server

import (
	"log"
	"net/http"
	"os"
)

type serverEnvs struct {
	serverAddr string
}

func newServerEnvConfig() *serverEnvs {
	defaultEnv := &serverEnvs{serverAddr: ":5001"}
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	if port == "" || host == "" {
		log.Printf("Create Server Addr: %s", defaultEnv.serverAddr)
		return defaultEnv
	}

	addr := host + port
	log.Printf("Create Server Addr: %s", addr)

	return &serverEnvs{serverAddr: addr}
}

func CreateEngine() *http.Server {
	env := newServerEnvConfig()
	handler := CreateHandler()
	return &http.Server{
		Addr:    env.serverAddr,
		Handler: *handler,
	}
}
