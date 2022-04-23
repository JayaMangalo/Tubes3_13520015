package connector

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InsertDataPenyakit(db *sql.DB, nama_penyakit string, dna_squence string) error {
	db, err2 := sql.Open("mysql", "root:rootstima@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	if err2 != nil {
		log.Fatal("Unable to open connection to db")
	}
	q := "INSERT INTO penyakit VALUES (?,?)"
	insert, err := db.Prepare(q)

	if err != nil {
		log.Fatal(err)
	}

	_, err = insert.Exec(nama_penyakit, dna_squence)

	if err != nil {
		return err
	}

	return err
}
