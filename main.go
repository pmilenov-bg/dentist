package main

import (
	"database/sql"
	"fmt"

	"github.com/pmilenov-bg/dentist/database"
	"github.com/pmilenov-bg/dentist/services"
)

func main() {
	fmt.Println("Dentist")
	var db *sql.DB = database.OpenDatabase("./data/patients.db")
	// db := database.OpenDatabase("./data/patients.db")
	// db.createTable
	defer db.Close()
	database.CreateTable(db)

	services.Search(db)
}
