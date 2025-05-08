package domainmodel

import (
	dbclient "github.com/dandk105/webapp_study/backend/internal"
	"log"
)

func CheckDatabaseConnection() bool {
	_, e := dbclient.CreateConnectedDatabaseClient()

	if e != nil {
		log.Printf("Failed Create DB Connection on domainmodel")
		return false
	} else {
		log.Printf("Sucess Create DB Connection on domainmodel")
		return true
	}
}
