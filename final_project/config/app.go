package config

import (
	"fmt"
	"log"

	"github.com/MuhGhifari/GolangBootcamp/final-project/models"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// dialect  = "postgres"
	// host     = "localhost"
	// dbPort   = "5432"
	// user     = "postgres"
	// password = "postgres"
	// dbName   = "my_gram"
	// db       *gorm.DB
	// err      error

	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbname   = "my_gram"
	db       *gorm.DB
	err      error
)

func StartDB() {

	// dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// db, err = gorm.Open(dialect, dbURI)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Printf("Listening in port:%v\n", dbPort)
	// }

	// defer db.Close()

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("sukses koneksi ke database")
	// db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.SocialMedia{})
	// db.AutoMigrate(&models.Photo{})
	// db.AutoMigrate(&models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
