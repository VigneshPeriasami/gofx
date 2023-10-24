package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func GetEnvArg(name string) string {
	return fmt.Sprintf("%s=%s", name, os.Getenv(name))
}

func GetDockerEnvArgs() []string {
	godotenv.Load("../.env")
	return []string{
		GetEnvArg("MYSQL_DATABASE"),
		GetEnvArg("MYSQL_USER"),
		GetEnvArg("MYSQL_PASSWORD"),
		GetEnvArg("MYSQL_ROOT_PASSWORD"),
		GetEnvArg("MYSQL_PORT"),
	}
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Second * 40

	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}
	err = pool.Client.Ping()

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.BuildAndRun(
		"mysql-test",
		"../Dockerfile",
		GetDockerEnvArgs(),
	)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		dsn := fmt.Sprintf(
			"%s:%s@(localhost:%s)/transactions",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"), resource.GetPort("3306/tcp"))
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Println("Waiting for docker database to start..")
			return err
		}
		return DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge database resource: %s", err)
	}

	os.Exit(code)
}

func TestTransactionsTable(t *testing.T) {
	require.NoError(t, DB.Ping(), "Db ping failed")

	rows, err := DB.Query("select count(*) from companies")
	require.NoError(t, err, "Failed querying companies table count")
	rows.Next()
	defer rows.Close()
	var count int
	rows.Scan(&count)
	assert.Equal(t, 1, count)
}
