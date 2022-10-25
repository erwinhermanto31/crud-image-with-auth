package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigration() {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user") + ":" + os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_name")) + "?charset=utf8&parseTime=true&loc=UTC"
	db, _ := sql.Open("mysql", urlConnection)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://repo/mysql/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Println(err)
	}

	err = m.Steps(1)
	if err != nil {
		log.Println(err)
	}
	err = m.Steps(2)
	if err != nil {
		log.Println(err)
	}
	err = m.Steps(3)
	if err != nil {
		log.Println(err)
	}
}
