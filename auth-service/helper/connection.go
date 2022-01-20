package helper

import (
	"auth-service/exception"
	"auth-service/model/app"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewConnection() *gorm.DB {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("DBSSLMODE")
	timezone := os.Getenv("DBTIMEZONE")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s timezone=%s",user,password,host,port,dbname,sslmode,timezone)
	db,err := gorm.Open(postgres.Open(dsn))
	exception.InternalServerError(err)

	pooling,err := db.DB()
	exception.InternalServerError(err)
	pooling.SetMaxIdleConns(10)
	pooling.SetMaxOpenConns(100)
	pooling.SetConnMaxIdleTime(5 * time.Minute)
	pooling.SetConnMaxLifetime(2 * time.Hour)

	return db
}

func initDataConnection() *gorm.DB {
	var err error

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := "postgres"
	sslmode := os.Getenv("DBSSLMODE")
	timezone := os.Getenv("DBTIMEZONE")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s timezone=%s",user,password,host,port,dbname,sslmode,timezone)

	i := 0

	for i != 3 {
		db,err := gorm.Open(postgres.Open(dsn))
		if err == nil {
			return db
		}
		i += 1
		time.Sleep(3 * time.Second)
	}
	exception.InternalServerError(err)
	return nil
}

func CreateDatabase() {
	db := initDataConnection()
	result := db.Exec("CREATE DATABASE sejutacita")
	fmt.Println(result.Error)
}

func SetupDatabase(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	Migrate(db)
	db.Exec("DELETE FROM users")
	init := []app.User{
		{
			Username:     "admin",
			NamaDepan:    "admin",
			NamaBelakang: "admin lagi",
			Level:        "admin",
			Password:     GeneratePassword("rahasia"),
		},
		{
			Username:     "user",
			NamaDepan:    "user",
			NamaBelakang: "user lagi",
			Level:        "user",
			Password:     GeneratePassword("rahasia"),
		},
	}
	db.Create(&init)
}
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&app.User{})
}