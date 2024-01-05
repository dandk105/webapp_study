package domainmodel

import (
	"context"
	schema "github.com/dandk105/webapp_study/backend/schema"

	dbclient "github.com/dandk105/webapp_study/backend/internal"
)

type UserDataAccesses struct {
	client *dbclient.Client
	schema *schema.User
	ctx    context.Context
}

/*
func (user *UserDataAccesses) GetUserdata() *schema.User{
	user.client = dbclient.CreateDatabaseClient()
	db := user.client.DataBaseConnection
	rows, e := db.Query("SELECT * FROM users where id = %d", user.schema.Id)
	if e != nil {
		user.client.FatalSQLF(e)
	}
	 := rows.Scan(&user.schema.Id, &user.schema.Name, &user.schema.Birthday)
	return
}

func (user *UserDataAccesses) CreateUserdata() {
	user.client = dbclient.CreateDatabaseClient()
	db := user.client.DataBaseConnection
	rows, e := db.Query("INSERT INTO users(%d,)", user.schema.Id)
	if e != nil {
		user.client.FatalSQLF(e)
	}
}
func (user *UserDataAccesses) DeleteUserdata() {}
func (user *UserDataAccesses) UpdateUserdata() {}
*/
