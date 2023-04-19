package models

type Patient struct {
	FirstName  string
	MiddleName string
	LastName   string
	Egn        int64
	NzokId     int64
	Gender     string
}

// is it possible to complete the PatientStatus from a table
type PatientStatus struct {
	Id int
}

// func ReturnPatientList(db *sql.DB) ([]Patient, error) {
// 	// func searchByNumb(db *sql.DB, choice string) ([]Patient, error) {
// 	choice := "1"
// 	rows, err := db.Query("SELECT * FROM patients WHERE egn LIKE ?", choice+"%") // this is just one of the options
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var patients []Patient
// 	for rows.Next() {
// 		var patient Patient
// 		// scan the patient fields here
// 		patients = append(patients, patient)
// 	}
// 	return patients, nil
// }
