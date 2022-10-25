package mysql

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/jmoiron/sqlx"
)

// test function that get list of lof file in n minutes
func TestGet(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	client := NewClient()
	query := &utils.Query{
		Filter: map[string]interface{}{
			"id$eq!": 1,
		},
	}
	var data entity.Images

	rows := sqlmock.NewRows([]string{"image_url"}).
		AddRow("abc")
	mock.ExpectQuery(regexp.QuoteMeta(QueryFindImage)).
		WillReturnRows(rows)

	err = client.Get(context.Background(), sqlxDB, &data, query, QueryFindImage)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

func TestSelect(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	client := NewClient()
	query := &utils.Query{
		Sort: "upload_time",
	}
	var data []entity.Images

	rows := sqlmock.NewRows([]string{"image_url"}).
		AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta(QueryFindImage)).
		WillReturnRows(rows)

	err = client.Select(context.Background(), sqlxDB, &data, query, QueryFindImage)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

func TestCreateOrUpdate(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	// mock.ExpectBegin()
	client := NewClient()
	data := &entity.Images{
		Id:         1,
		ImageURL:   "test",
		UploadTime: time.Now(),
	}

	mock.ExpectQuery(QueryInsertImage)

	_, err = client.CreateOrUpdate(context.Background(), sqlxDB, &data, QueryInsertImage)
	if err == nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}
