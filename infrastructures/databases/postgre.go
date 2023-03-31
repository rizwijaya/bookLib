package database

import (
	"bookLib/infrastructures/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func MigrateDB(db *sql.DB) error {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS books(
		id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		title varchar(100) NOT NULL,
		author varchar(100) NOT NULL,
		description varchar(400) NOT NULL
	);
	`
	_, err := db.Exec(sqlTable)
	if err != nil {
		return err
	}
	return nil
}

func Postgres() (*sql.DB, error) {
	config, err := config.New()
	if err != nil {
		return nil, err
	}
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDatabases() *sql.DB {
	db, err := Postgres()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	fmt.Println("Successfully connected!")

	// err = MigrateDB(db)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return nil
	// }
	// fmt.Println("Migration successfully done")

	return db
}
