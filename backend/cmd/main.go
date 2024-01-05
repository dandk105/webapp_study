package main

import (
	engine "github.com/dandk105/webapp_study/backend/server"
	"log"
)

func main() {
	server := engine.CreateEngine()
	log.Fatal(server.ListenAndServe())

}
