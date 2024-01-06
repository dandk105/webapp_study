package internal

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var existsEnvs = &dbEnvs{dbUser: "myuser", dbName: "mydatabase", dbPassword: "mypassword", dbHost: "localhost"}
var defaultEnvs = &dbEnvs{dbUser: "default", dbName: "default", dbPassword: "default", dbHost: "localhost"}

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

	a := *createDBConfigEnvs()

	assert.Equal(t, tc.dbUser, a.dbUser)
	assert.Equal(t, tc.dbName, a.dbName)
	assert.Equal(t, tc.dbPassword, a.dbPassword)
	assert.Equal(t, tc.dbHost, a.dbHost)
}

func TestEmptyNewDBConfigEnvs(t *testing.T) {
	SetupEmptyEnv()
	tc := *defaultEnvs

	a := *createDBConfigEnvs()

	assert.Equal(t, tc.dbUser, a.dbUser)
	assert.Equal(t, tc.dbName, a.dbName)
	assert.Equal(t, tc.dbPassword, a.dbPassword)
	assert.Equal(t, tc.dbHost, a.dbHost)
}

func TestNewDBSourceName(t *testing.T) {
	SetupExistsEnv()
	tc := *existsEnvs
	d := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", tc.dbUser, tc.dbName, tc.dbPassword, tc.dbHost)

	name := createDataSourceName()
	dsn := name.DBSourceName

	assert.Equal(t, d, dsn)
}

// TODO: Flakyなテストではなく必ず成功する様にテストを書く

func TestCreateConnection(t *testing.T) {
	SetupExistsEnv()

	c := createDataSourceName()
	client := Client{}
	// Databaseが起動している時にはエラーが発生しないが、
	// Databaseが起動していない時は必ずエラーが発生する
	e := client.createConnection(c.DBSourceName)
	assert.NoError(t, e)
}

// TODO: Flakyを直す
// Databaseを起動していないといけない
func TestSetDatabaseClient(t *testing.T) {
	SetupExistsEnv()
	a := CreateConnectedDatabaseClient()

	assert.NotNil(t, a)
}

func TestSetDBClientLogger(t *testing.T) {
	// Set up test environment variables
	dbc := &Client{}
	dbc.SetDataBaseClientLogger()
	assert.NotNil(t, dbc.Log)
	dbc.Log.Printf("Sucess")
}
