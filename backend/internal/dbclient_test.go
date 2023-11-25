package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDataSourceName(t *testing.T) {
	// Set up test environment variables
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")

	expectedDSN := "user=testuser dbname=testdb password=testpassword host=localhost sslmode=disable"

	// Call the function being tested
	dbinit := &DBInitializer{}
	dbinit.CreateDataSourceName()

	// Check if the returned DSN matches the expected DSN
	assert.Equal(t, expectedDSN, dbinit.dsn)
}

// TODO: テストの改善を行う
// テストの実施とDebugを行おうとしても、全てプロセスが正常に終了したというメッセージのみで、
// テストの実施結果が表示されず、正しく接続されているのかも分からない
func TestCreateConnection(t *testing.T) {
	switch os.Getenv("TEST_ENV") {
	case "unit":
		// Set up test environment variables
		dbc := &DBClient{}
		dbc.CreateConnection()
		assert.NotNil(t, dbc.DBConneciton)
		assert.Nil(t, dbc.DBConneciton.Ping())
	case "integration":
		// Set up test environment variables for integration testing
		dbc := &DBClient{}
		dbc.CreateConnection()
		assert.NotNil(t, dbc.DBConneciton)
		assert.Nil(t, dbc.DBConneciton.Ping())
	}
}

func TestSetDBClientLogger(t *testing.T) {
	// Set up test environment variables
	dbc := &DBClient{}
	dbc.SetDBClientLogger()
	assert.NotNil(t, dbc.logger)
	dbc.logger.Print("Test")
}
