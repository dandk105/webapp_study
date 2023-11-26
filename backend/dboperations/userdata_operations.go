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

func (accesser *UserDataAccesser) GetUserdata() {
	accesser.dbclient.CreateConnection()
	db := accesser.dbclient.DataBaseConnection
	db.Query("SELECT * FROM users where id = %d", accesser.user.Id)
}
func (accesser *UserDataAccesser) CreateUserdata() {
	accesser.dbclient.CreateConnection()
	db := accesser.dbclient.DataBaseConnection
	db.Query("SELECT * FROM users where id = %d", accesser.user.Id)
	db.Query("INSERT INTO users (name, birthday) VALUES (%s, %s)", accesser.user.Name, accesser.user.Birthday)
}
func (user *UserDataAccesser) DeleteUserdata() {}
func (user *UserDataAccesser) UpdateUserdata() {}
