package databaseClientConfig

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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

	cb := context.Background()
	c := createDataSourceName(cb)
	v := c.Value("DataSourceName")

	assert.Equal(t, d, v)
}
