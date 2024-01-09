package EngineConfig

import (
	"log"
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
