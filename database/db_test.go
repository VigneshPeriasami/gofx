package database

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vigneshperiasami/analytics/resource"
)

var DB *sql.DB

func TestMain(m *testing.M) {
	resource := resource.CreateNewSqlDb("mysql_setup", "./Dockerfile")
	DB = resource.Db
	code := m.Run()

	resource.Purge()
	os.Exit(code)
}

func TestCompaniesTable(t *testing.T) {
	require.NoError(t, DB.Ping(), "Db ping failed")

	rows, err := DB.Query("select count(*) from companies")
	require.NoError(t, err, "Failed querying companies table count")
	rows.Next()
	defer rows.Close()
	var count int
	rows.Scan(&count)
	assert.Equal(t, 2, count)
}
