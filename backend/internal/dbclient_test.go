package db

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

	// 構造体を初期化して、ポインタ参照を行っている
	dbinit := &DBInitializer{}
	// 環境変数が取得できなかった時に、ランタイムエラーが発生してしまう。
	dbinit.CreateDataSourceName()

	// Check if the returned DSN matches the expected DSN
	assert.Equal(t, expectedDSN, dbinit.dsn)
}

/*
ここでのテストは、環境変数が取得できなかった時に、デフォルト値が
設定されるかどうかを確認するテスト
*/
func TestCreateDataSourceNameWithEmptyEnv(t *testing.T) {
	// Set up test environment variables
	os.Setenv("DB_USER", "")
	os.Setenv("DB_NAME", "")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_HOST", "")

	expectedDSN := "user=myuser dbname=mydatabase password=mypassword host=localhost sslmode=disable"

	// Call the function being tested
	dbinit := &DBInitializer{}
	// 環境変数が取得できなかった時に、ランタイムエラーが発生してしまう。
	dbinit.CreateDataSourceName()

	// Check if the returned DSN matches the expected DSN
	assert.Equal(t, expectedDSN, dbinit.dsn)
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
		dbc.CreateConnection()
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
