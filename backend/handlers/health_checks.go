package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/dandk105/webapp_study/backend/internal"
	schma "github.com/dandk105/webapp_study/backend/schema"
)

func DatabaseStatusCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

	client := db.Client{}
	e := client.DataBaseConnection.Ping()
	if e != nil {
		// DBへの接続が成功した時はJSON形式でstatus:OK を返す
		response := schma.StatusResponse{Status: "OK"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error:Json marshal %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		log.Printf("Cannot Create DB Connection on Handler")
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}

}
