package services

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/mattn/go-sqlite3"

	"github.com/pmilenov-bg/dentist/models"
	"github.com/pmilenov-bg/dentist/screens"
)

func Search(db *sql.DB) ([]models.Patient, error) {

	var patients []models.Patient
	var err error
	choice := screens.Prompt(`type first letters of the name to search by name
	type the first few digits of the EGN to search by egn
	type the - and the last two digits of the year to search for patients not seen that year
`)

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
		patients, err = searchByInitials(db, choice)
	case yearRegex.MatchString(choice):
		// Search by year
		fmt.Println("shear by year")
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
	fmt.Println("Entering a search by initials")
	rows, err := db.Query("SELECT * FROM patients WHERE CONCAT(SUBSTR(fname, 1, 1), SUBSTR(lname, 1, 1)) LIKE ?", choice+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		// scan the patient fields here
		patients = append(patients, patient)
	}
	fmt.Println("end of search by INIT")
	fmt.Println(patients)
	return patients, nil
}

func searchByNumb(db *sql.DB, choice string) ([]models.Patient, error) {
	fmt.Println("Entering a search by EGN")
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
	}
	fmt.Println("end of search by EGN")
	fmt.Println(patients)
	return patients, nil
}

func searchByYear(db *sql.DB, choice string) ([]models.Patient, error) {
	fmt.Println("Entering a search by Year")
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
	}
	// for patients = range.
	fmt.Println("end of search func")
	fmt.Println(patients)
	return patients, nil

}
