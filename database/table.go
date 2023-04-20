package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pmilenov-bg/dentist/models"
)

// type Patient struct {
// 	FirstName  string
// 	MiddleName string
// 	LastName   string
// 	Egn        int64
// 	NzokId     int64
// 	Gender     string
// }

// // is it possible to complete the PatientStatus from a table
//
//	type PatientStatus struct {
//		Id int
//	}
func OpenDatabase(filepath string) *sql.DB {
	// Open the database file (creating it if it doesn't exist)
	fmt.Printf("Opening database%s\n", filepath)
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var LAST_MIGRATION int = 2

func CreateTable(db *sql.DB) {
	fmt.Printf("Begin creating table\n")

	// createTable := `
	// CREATE TABLE IF NOT EXISTS migrations (
	// 	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 	last_version varchar,
	// `

	// Check if the "patients" table exists would it make a difference?
	// var tableName string
	// err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='patients'").Scan(&tableName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if tableName != "patients" {
	// 	log.Fatal("patients table does not exist")
	// }

	// Create (if not exist) the "patients" table
	// how to check externally whether the table exist?

	createTable := `
		CREATE TABLE IF NOT EXISTS patients (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            fname TEXT NOT NULL,
            mname TEXT,
            lname TEXT NOT NULL,
            egn BIGINT UNSIGNED NOT NULL,
            nzok_id BIGINT UNSIGNED,
            gender TEXT NOT NULL
        )
    `
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// Read the data from the CSV file
	file, err := os.Open("./data/patients.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Begin populating the table")
	// Insert the data into the "patients" table
	for _, record := range records {
		fmt.Println(record)
		patient := models.Patient{
			FirstName:  record[0],
			MiddleName: record[1],
			LastName:   record[2],
			Egn:        parseEgn(record[3]),
			NzokId:     parseNzokId(record[4]),
			Gender:     record[5],
		}

		insert := `
            INSERT INTO patients (fname, mname, lname, egn, nzok_id, gender)
            VALUES (?, ?, ?, ?, ?, ?)
        `
		_, err := db.Exec(insert, patient.FirstName, patient.MiddleName, patient.LastName, patient.Egn, patient.NzokId, patient.Gender)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("creating the table Done")
}

func parseEgn(s string) int64 {
	egn, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return egn
}

func parseNzokId(s string) int64 {
	if s == "" {
		return 0
	}
	nzokId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return nzokId
}
