package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

var AppDB *sqlx.DB

func InitCon() {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user") + ":" + os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_name")) + "?charset=utf8&parseTime=true&loc=UTC"

	// fmt.Printf("⇨ Connect MYSQL to Server %s ... \n", hostName)
	// log.Println(urlConnection)
	db, err := apmsql.Open(os.Getenv("driver"), urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", os.Getenv("driver"), os.Getenv("db_host"), os.Getenv("db_port"), err.Error())
	}

	dbx := sqlx.NewDb(db, os.Getenv("driver"))
	err = dbx.Ping()
	if err != nil {
		log.Println(err)
	}

	AppDB = dbx
}
