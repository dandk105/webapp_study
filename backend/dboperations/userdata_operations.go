package main

import (
	"database/sql"
	"github.com/dandk105/webapp_study/backend/internal"
)

type UserDataHandler struct {
	dbconn DBClient
	db     *sql.DB
	user   User
}

func (handler *UserDataHandler) GetUserdata() {
	db := handler.dbconn.CreateConnection()
	db.Query("SELECT * FROM users where id = %d", handler.user.Id)
}
func (user *UserDataHandler) CreateUserdata() {}
func (user *UserDataHandler) DeleteUserdata() {}
func (user *UserDataHandler) UpdateUserdata() {}
