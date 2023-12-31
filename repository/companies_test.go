package repository

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vigneshperiasami/analytics/database"
	"github.com/vigneshperiasami/analytics/models"
	"github.com/vigneshperiasami/analytics/resource"
)

var DbClient database.DbClient

var FixedTime = time.Date(
	2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

func TestMain(m *testing.M) {
	resource := resource.CreateNewSqlDb("repository_setup", "../database/Dockerfile")
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

func TestInsertCompanies(t *testing.T) {
	client := NewCompanyClient(DbClient)
	companies := []models.Company{
		{
			Id:      12,
			Ibans:   "Idk",
			Name:    "Whatever",
			Address: "Wherever",
		},
	}
	err := client.InsertCompanies(companies)
	require.NoError(t, err, "Error inserting new company")
	newList, err := client.GetAllCompanies()
	require.NoError(t, err, "Couldn't read companies after insert")
	assert.Equal(t, 3, len(newList))

	// remove the new company
	db, err := DbClient.Open()
	defer db.Close()
	_, err = db.Query("DELETE from companies where ID=?", companies[0].Id)
	require.NoError(t, err, "error deleting the inserted record")
}

func TestInsertTransactions(t *testing.T) {
	client := NewCompanyClient(DbClient)
	transactions := []models.Transaction{
		{
			Id:          "hey",
			Amount:      2.1,
			Currency:    "swd",
			Sender:      "random-id-of-the-sender",
			Beneficiary: "random-id-of-the-receiver",
			Timestamp:   FixedTime,
		},
	}
	err := client.InsertTransactions(transactions)
	require.NoError(t, err, "Couldn't insert companies")
}
