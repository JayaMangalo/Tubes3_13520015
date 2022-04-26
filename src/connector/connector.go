package connector

import (
	"database/sql"
	"log"
	"main/src/algorithm"

	_ "github.com/go-sql-driver/mysql"
)

type Penyakit struct {
	Nama string `json:"nama"`
	DNA  string `json:"dna"`
}
type Data struct {
	Tanggal   string `json:"tanggal"`
	Nama      string `json:"nama"`
	Penyakit  string `json:"penyakit"`
	Diagnosis string `json:"diagnosis"`
}

type Result struct {
	Hasil []Data
}

func InsertDataPenyakit(nama_penyakit string, dna_squence string) error {
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

func InsertDataDiagnosis(nama_penyakit string, nama_pengguna string, hasil_diagnosis string, tanggal string) error {
	db, err2 := sql.Open("mysql", "root:rootstima@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	if err2 != nil {
		log.Fatal("Unable to open connection to db")
	}
	q := "INSERT INTO diagnosis_penyakit VALUES (?,?,?,?)"
	insert, err := db.Prepare(q)

	if err != nil {
		log.Fatal(err)
	}

	_, err = insert.Exec(tanggal, nama_pengguna, nama_penyakit, hasil_diagnosis)

	if err != nil {
		return err
	}

	return err
}

func GetDataOrang(patternRegex string) Result {
	res1 := Result{}
	res := []Data{}
	db, err2 := sql.Open("mysql", "root:rootstima@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	if err2 != nil {
		log.Fatal("Unable to open connection to db")
	}

	results, err := db.Query("SELECT * FROM diagnosis_penyakit")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(algorithm.ReadRegex(patternRegex))

	for results.Next() {
		var data Data
		// for each row, scan the result into our tag composite object
		err = results.Scan(&data.Tanggal, &data.Nama, &data.Penyakit, &data.Diagnosis)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_info := data.Tanggal + " " + data.Penyakit
		regex_info := algorithm.ReadRegex(data_info)
		regex_search := algorithm.ReadRegex(patternRegex)

		log.Println(regex_info)
		log.Println(regex_search)

		// bandinginnya gimana?

		res = append(res, data)
		// and then print out the tag's Name attribute
	}
	res1.Hasil = res
	return res1
}

func GetSequencePenyakit(penyakit string) string {
	db, err2 := sql.Open("mysql", "root:rootstima@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	res := ""

	if err2 != nil {
		log.Fatal("Unable to open connection to db")
	}

	results, err := db.Query("SELECT * FROM penyakit")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var data Penyakit
		// for each row, scan the result into our tag composite object
		err = results.Scan(&data.Nama, &data.DNA)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if data.Nama == penyakit {
			res = data.DNA
			break
		}
		// and then print out the tag's Name attribute
	}
	return res
}
