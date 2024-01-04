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

	name := CreateDataSourceName()
	dsn := name.DBSourceName

	assert.Equal(t, d, dsn)
}

// TODO: Flakyなテストではなく必ず成功する様にテストを書く

func TestCreateConnection(t *testing.T) {
	SetupExistsEnv()

	c := CreateDataSourceName()
	client := Client{}
	// Databaseが起動している時にはエラーが発生しないが、
	// Databaseが起動していない時は必ずエラーが発生する
	e := client.CreateConnection(c.DBSourceName)
	assert.NoError(t, e)
}

// TODO: Flakyを直す
// Databaseを起動していないといけない
func TestSetDatabaseClient(t *testing.T) {
	SetupExistsEnv()
	a := SetDatabaseClient()

	assert.NotNil(t, a)
}

func TestSetDBClientLogger(t *testing.T) {
	// Set up test environment variables
	dbc := &Client{}
	dbc.SetDataBaseClientLogger()
	assert.NotNil(t, dbc.Log)
	dbc.Log.Printf("Sucess")
}
