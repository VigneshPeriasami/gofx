package sqlhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	Hello string
}

func TestInsertQuery(t *testing.T) {
	queryBuilder := insertQuery("dummy", []string{"columnOne", "columnTwo"})
	assert.Equal(t, "INSERT into dummy (columnOne,columnTwo) VALUES (?,?),(?,?)", queryBuilder(2))
}

func TestInsertQueryPlaces(t *testing.T) {
	assert.Equal(t, "(?,?)", queryPlaces(1, 2))
	assert.Equal(t, "(?,?,?,?),(?,?,?,?)", queryPlaces(2, 4))
}

func TestFlatten(t *testing.T) {
	rows := []Dummy{{"hello"}, {"There"}, {"!!"}}
	flattend := Flatten[Dummy](rows, func(d Dummy) []interface{} {
		return []interface{}{d.Hello}
	})

	assert.Equal(t, []interface{}{"hello", "There", "!!"}, flattend)
}
