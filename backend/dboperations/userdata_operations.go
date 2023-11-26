package dboperations

import (
	"context"
	"database/sql"

	dbclient "github.com/dandk105/webapp_study/backend/internal"
)

type UserDataAccesser struct {
	dbclient dbclient.Client
	db       *sql.DB
	user     User
	ctx      context.Context
}

func (handler *UserDataAccesser) GetUserdata() {
	handler.dbclient.CreateConnection()
	db := handler.dbclient.DataBaseConnection
	db.Query("SELECT * FROM users where id = %d", handler.user.Id)
}
func (user *UserDataAccesser) CreateUserdata() {}
func (user *UserDataAccesser) DeleteUserdata() {}
func (user *UserDataAccesser) UpdateUserdata() {}
