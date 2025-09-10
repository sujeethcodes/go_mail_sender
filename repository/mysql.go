package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/matryer/resync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var onceMysql resync.Once
var mysqlCon *gorm.DB

type MysqlCon struct {
	Connection *gorm.DB
}

func SingleTonPattern() *MysqlCon {
	onceMysql.Do(func() {

		userName := os.Getenv("USER_NAME")
		password := os.Getenv("PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		// DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
			userName, password, dbHost, dbName,
		)

		// using gorm to open mysql connection
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not create MySQL connection", err)
		}

		fmt.Println("DB connection successfull")
		mysqlCon = db

	})
	return &MysqlCon{Connection: mysqlCon}

}
