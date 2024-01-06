package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	domain "github.com/dandk105/webapp_study/backend/domainmodel"
	schma "github.com/dandk105/webapp_study/backend/schema"
)

func DatabaseStatusCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s request: %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

	result := domain.CheckDatabaseConnection()

	if result == false {
		// ここユースケースで失敗した時のログという風に切り出したい
		// 標準出力にエラーとしてログを出したい
		log.Printf("Cannot Create DB Connection on Handler")
		http.Error(w, "Failed DB Connection Check", http.StatusInternalServerError)
	} else {
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
	}
}
