package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var existsEnvs = &Envs{DbUser: "myuser", DbName: "mydatabase", DbPassword: "mypassword", DbHost: "localhost"}
var defaultEnvs = &Envs{DbUser: "default", DbName: "default", DbPassword: "default", DbHost: "localhost"}

func SetupExistsEnv() {
	u := "myuser"
	n := "mydatabase"
	p := "mypassword"
	h := "localhost"

	os.Setenv("DB_USER", u)
	os.Setenv("DB_NAME", n)
	os.Setenv("DB_PASSWORD", p)
	os.Setenv("DB_HOST", h)
}

func SetupEmptyEnv() {
	e := ""

	os.Setenv("DB_USER", e)
	os.Setenv("DB_NAME", e)
	os.Setenv("DB_PASSWORD", e)
	os.Setenv("DB_HOST", e)
}

func TestNewDBConfigEnvs(t *testing.T) {
	SetupExistsEnv()
	tc := *existsEnvs

	a := Envs{}
	a = *NewDBConfigEnvs()

	assert.Equal(t, tc.DbUser, a.DbUser)
	assert.Equal(t, tc.DbName, a.DbName)
	assert.Equal(t, tc.DbPassword, a.DbPassword)
	assert.Equal(t, tc.DbHost, a.DbHost)
}

func TestEmptyNewDBConfigEnvs(t *testing.T) {
	SetupEmptyEnv()
	tc := *defaultEnvs

	a := Envs{}
	a = *NewDBConfigEnvs()

	assert.Equal(t, tc.DbUser, a.DbUser)
	assert.Equal(t, tc.DbName, a.DbName)
	assert.Equal(t, tc.DbPassword, a.DbPassword)
	assert.Equal(t, tc.DbHost, a.DbHost)
}

func TestNewDBSourceName(t *testing.T) {
	SetupExistsEnv()
	tc := *existsEnvs
	d := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", tc.DbUser, tc.DbName, tc.DbPassword, tc.DbHost)

	dbconf := CreateDataSourceName()
	dsn := dbconf.DBSourceName

	assert.Equal(t, d, dsn)
}

// TODO: 例外が発生しないようなテストを書く
// 制約: PostgresDatabaseをlocalでデフォルトの値で動かしている必要がある
// なお、DBを稼働させていなかった場合は、Pingが接続されずに失敗する

func TestCreateConnection(t *testing.T) {
	SetupExistsEnv()

	c := CreateDataSourceName()
	client := Client{}
	client.CreateConnection(c.DBSourceName)
}

func TestSetDBClientLogger(t *testing.T) {
	// Set up test environment variables
	dbc := &Client{}
	dbc.SetDataBaseClientLogger()
	assert.NotNil(t, dbc.logger)
	dbc.logger.Print("Test")
}
