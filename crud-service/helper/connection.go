package helper

import (
	"crud-service/exception"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewConnection() *gorm.DB {
	var db *gorm.DB
	var err error

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("DBSSLMODE")
	timezone := os.Getenv("DBTIMEZONE")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s timezone=%s",user,password,host,port,dbname,sslmode,timezone)
	
	i := 0

	for i != 3 {
		db,err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		exception.InternalServerError(err)
	}

	pooling,err := db.DB()
	exception.InternalServerError(err)
	pooling.SetMaxIdleConns(10)
	pooling.SetMaxOpenConns(100)
	pooling.SetConnMaxIdleTime(5 * time.Minute)
	pooling.SetConnMaxLifetime(2 * time.Hour)

	return db
}
