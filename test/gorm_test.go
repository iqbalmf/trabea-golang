package test

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:password@tcp(localhost:3306)/trabeadb?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSql(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?,?)", "1", "Iqbal").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "2", "Fauzan").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "3", "Fauzan").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "4", "Fauzan").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestQuerySql(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.Id)
	assert.Equal(t, "Iqbal", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()
	var samples []Sample
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			id, name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}
