package resource

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ory/dockertest"
)

type DbResource struct {
	Resource *dockertest.Resource
	Db       *sql.DB
	Conn     string
	Pool     *dockertest.Pool
}

func (r *DbResource) Purge() {
	if err := r.Pool.Purge(r.Resource); err != nil {
		log.Fatalf("Could not purge database resource: %s", err)
	}
}

var Resource DbResource

func GetEnvArg(name string) string {
	return fmt.Sprintf("%s=%s", name, os.Getenv(name))
}

func getDockerEnvArgs() []string {
	return []string{
		GetEnvArg("MYSQL_DATABASE"),
		GetEnvArg("MYSQL_USER"),
		GetEnvArg("MYSQL_PASSWORD"),
		GetEnvArg("MYSQL_ROOT_PASSWORD"),
		GetEnvArg("MYSQL_PORT"),
	}
}

func CreateNewSqlDb(name string, dockerFilePath string) DbResource {
	godotenv.Load("../.env")
	var db *sql.DB
	var conn string
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
		name,
		dockerFilePath,
		getDockerEnvArgs(),
	)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		conn = fmt.Sprintf(
			"%s:%s@(localhost:%s)/transactions",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"), resource.GetPort("3306/tcp"))
		db, err = sql.Open("mysql", conn)
		if err != nil {
			log.Println("Waiting for docker database to start..")
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return DbResource{
		Db:       db,
		Resource: resource,
		Pool:     pool,
		Conn:     conn,
	}
}
