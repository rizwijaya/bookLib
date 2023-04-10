package database

import (
	"bookLib/modules/v1/book/domain"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres() (*gorm.DB, error) {
	// config, err := config.New()
	// if err != nil {
	// 	return nil, err
	// }
	db_name := os.Getenv("DB_NAME")
	db_username := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	dsn := "host=" + db_host + " user=" + db_username + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=Asia/Jakarta"
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

	err = db.AutoMigrate(&domain.Book{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	fmt.Println("Migration successfully done")

	return db
}
