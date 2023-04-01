package database

import (
	"bookLib/infrastructures/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres() (*gorm.DB, error) {
	config, err := config.New()
	if err != nil {
		return nil, err
	}
	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return Db, nil
}

func NewDatabases() *gorm.DB {
	db, err := Postgres()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	fmt.Println("Successfully connected!")

	// err = db.AutoMigrate(&domain.Book{})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return nil
	// }
	// fmt.Println("Migration successfully done")

	return db
}
