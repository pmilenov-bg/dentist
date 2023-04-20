package services

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/pmilenov-bg/dentist/models"
	"github.com/pmilenov-bg/dentist/screens"
)

func TermsQuery() (choice, string) {
	choice := screens.Prompt(`type first letters of the name to search by name
	type the first few digits of the EGN to search by egn
	type the - and the last two digits of the year to search for patients not seen that year
`)
	fmt.Printf("You entered %s\n", choice)

	return choice
}

func Search(db *sql.DB) ([]models.Patient, error) {
	// take out the QUESTION section
	var patients []models.Patient
	var err error
	choice := screens.Prompt(`type first letters of the name to search by name
	type the first few digits of the EGN to search by egn
	type the - and the last two digits of the year to search for patients not seen that year
`)
	// take out the QUESTION section

	// Use regular expressions to identify the format of the Prompt string
	egnRegex := regexp.MustCompile(`^\d{6}$`) // ^\d{10}$ you can reduce the typing by shortening the inpit digits
	initialsRegex := regexp.MustCompile(`^[A-Za-z]{2}$`)
	yearRegex := regexp.MustCompile(`^-([0-9]{2})$`)

	switch {
	case egnRegex.MatchString(choice):
		// Search by EGN
		fmt.Println("search by EGN")
		patients, err = searchByNumb(db, choice)
	case initialsRegex.MatchString(choice):
		// Search by initials
		fmt.Println("search by initials")
		choice := strings.ToUpper(choice)
		patients, err = searchByInitials(db, choice)
	case yearRegex.MatchString(choice):
		// Search by year
		fmt.Println("search by year")
		patients, err = searchByYear(db, choice)
	default:
		err = fmt.Errorf("invalid input format")
	}

	if err != nil {
		return nil, err
	}
	return patients, nil
}

func searchByInitials(db *sql.DB, choice string) ([]models.Patient, error) {
	fmt.Printf("You typed %s Entering a search by initials\n", choice)
	rows, err := db.Query("SELECT substr(fname, 1, 1) || substr(lname, 1, 1) AS initials FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		// scan the patient fields here
		patients = append(patients, patient)
		fmt.Println(patients)

	}
	fmt.Println("end of search by INIT")
	return patients, nil
}

func searchByNumb(db *sql.DB, choice string) ([]models.Patient, error) {
	fmt.Printf("You typed %s Entering a search by EGN\n", choice)
	// assuming db is a *sql.DB object
	rows, err := db.Query("SELECT * FROM patients WHERE egn LIKE ?", choice+"%")

	// can i substitude the following code and do it as a separate func somewhere else
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		// scan the patient fields here
		patients = append(patients, patient)
		fmt.Println(patient)

	}
	fmt.Println("end of search by EGN")
	return patients, nil
}

func searchByYear(db *sql.DB, choice string) ([]models.Patient, error) {
	fmt.Printf("You typed %s Entering a search by Year\n", choice)
	// assuming db is a *sql.DB object
	lastTwo := choice[1:]
	rows, err := db.Query("SELECT * FROM patients WHERE year LIKE ?", "%"+lastTwo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		// scan the patient fields here
		patients = append(patients, patient)
		fmt.Println(patients)

	}
	// for patients = range.
	fmt.Println("end of search func")
	return patients, nil

}

// func Display(){
// 	fmt.Println("enterint the Display func")

// 	for _, p := range []models.Patient {
//         fmt.Printf("patients search list %s,\n", p)
// 	}
// }
