package repository

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/resource"
)

var DbClient database.DbClient

func TestMain(m *testing.M) {
	resource := resource.CreateNewSqlDb("repository_setup", "../database/Dockerfile")
	log.Println(resource.Conn)
	DbClient = database.NewDbClient(database.DbParams{Conn: resource.Conn})
	code := m.Run()
	resource.Purge()
	os.Exit(code)
}

func TestCompaniesCount(t *testing.T) {
	client := NewCompanyClient(DbClient)
	count, err := client.GetCompanyTotalCount()
	require.NoError(t, err, "Couldn't count companies")
	assert.Equal(t, 2, count)
}

func TestGetCompanies(t *testing.T) {
	client := NewCompanyClient(DbClient)
	companies, err := client.GetAllCompanies()
	require.NoError(t, err, "Couldn't read companies")
	assert.Equal(t, 2, len(companies))
}
