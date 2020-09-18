package Models

import (
	"database/sql"
	"log"
)

type Employees struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
	City           string `json:"city"`
}

type EmployeeDetails struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
	City           string `json:"city"`
	Phone          string `json:"phone"`
	AddressLine1   string `json:"adressLine1"`
	AddressLine2   string `json:"adressLine2"`
	State          string `json:"state"`
	Country        string `json:"country"`
	PostalCode     string `json:"postalCode"`
	Territory      string `json:"territory"`
}

func (repository *Repository) GetAllEmployees() ([]Employees, error) {
	stmtEmp := "SELECT employeeNumber, lastName, firstName, extension, email, jobTitle, city FROM employees INNER JOIN offices ON employees.officeCode = offices.officeCode"

	rows, _ := repository.Conn.Query(stmtEmp)
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Extension      string
		Email          string
		JobTitle       string
		City           string
	)

	employeesList := make([]Employees, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&EmployeeNumber, &LastName, &FirstName, &Extension, &Email, &JobTitle, &City); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			employee := Employees{
				EmployeeNumber: EmployeeNumber,
				LastName:       LastName,
				FirstName:      FirstName,
				Extension:      Extension,
				Email:          Email,
				JobTitle:       JobTitle,
				City:           City,
			}
			employeesList = append(employeesList, employee)
		default:
			return nil, err
		}
	}
	return employeesList, nil
}

func (repository *Repository) GetEmployeeAction(id uint64) (EmployeeDetails, error) {
	sqlStmt := "SELECT employeeNumber, lastName, firstName, email, jobTitle, city, phone, addressLine1, addressLine2, state, country, postalCode, territory FROM employees INNER JOIN offices ON employees.officeCode = offices.officeCode WHERE employeeNumber = ?"
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Email          string
		JobTitle       string
		City           string
		Phone          string
		AddressLine1   sql.NullString
		AddressLine2   string
		State          sql.NullString
		Country        string
		PostalCode     string
		Territory      string
	)
	stmt, err := repository.Conn.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&EmployeeNumber, &LastName, &FirstName, &Email, &JobTitle, &City, &Phone, &AddressLine1, &AddressLine2, &State, &Country, &PostalCode, &Territory)
	if err != nil {
		log.Fatal(err)
	}
	customer := EmployeeDetails{
		EmployeeNumber: EmployeeNumber,
		LastName:       LastName,
		FirstName:      FirstName,
		Email:          Email,
		JobTitle:       JobTitle,
		City:           City,
		Phone:          Phone,
		AddressLine1:   AddressLine1.String,
		AddressLine2:   AddressLine2,
		State:          State.String,
		Country:        Country,
		PostalCode:     PostalCode,
		Territory:      Territory,
	}
	return customer, nil
}
