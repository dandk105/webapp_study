package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var existsEnvs = &Envs{DbUser: "testuser", DbName: "testdb", DbPassword: "testpassword", DbHost: "localhost"}
var emptyEnvs = &Envs{DbUser: "", DbName: "", DbPassword: "", DbHost: ""}
var defaultEnvs = &Envs{DbUser: "default", DbName: "default", DbPassword: "default", DbHost: "localhost"}

func SetupExistsEnv() {
	u := "testuser"
	n := "testdb"
	p := "testpassword"
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

	c := DatabaseSourceConfig{}
	dbconf := c.CreateDataSourceName()
	dsn := dbconf.DBSourceName

	assert.Equal(t, d, dsn)
}

// TODO: 例外が発生しないようなテストを書く
// A2023/11/26 15:55:50 Set Dsn user= dbname= password= host= sslmode=disable
// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
func TestCreateConnection(t *testing.T) {
	switch os.Getenv("TEST_TYPE") {
	case "integration":
		// Set up test environment variables for integration testing
		dbc := &Client{}
		// ここのDataBaseへの接続の際に、中で隠蔽しているDSNの作成で空白文字が
		// 生成されてしまうので、テストが失敗してしまう
		dsn := "user=test dbname=test password=test host=test sslmode=disable"
		dbc.CreateConnection(dsn)
		assert.NotNil(t, dbc.DataBaseConnection)
	default:
		// Set up test environment variables
		dbc := &Client{}

		dsn := "user=default dbname=default password=default host=localhost sslmode=disable"
		dbc.CreateConnection(dsn)
		assert.NotNil(t, dbc.DataBaseConnection)
		assert.Nil(t, dbc.DataBaseConnection.Ping())
	}
}

func TestSetDBClientLogger(t *testing.T) {
	// Set up test environment variables
	dbc := &Client{}
	dbc.SetDataBaseClientLogger()
	assert.NotNil(t, dbc.logger)
	dbc.logger.Print("Test")
}
