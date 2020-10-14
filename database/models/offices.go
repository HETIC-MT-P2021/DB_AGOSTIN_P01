package models

import (
	"database/sql"
)

type Offices struct {
	OfficeCode  uint       `json:"officeCode"`
	City        string     `json:"city"`
	Phone       string     `json:"phone"`
	AdressLine1 string     `json:"adressLine1"`
	AdressLine2 string     `json:"adressLine2"`
	State       string     `json:"state"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postalCode"`
	Territory   string     `json:"territory"`
	Employees   []Employee `json:"employee"`
}

type Employee struct {
	EmployeeNumber uint   `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
	City           string `json:"city"`
}

func (repository *Repository) GetOfficesAction() ([]Offices, error) {
	stmtOffices := "SELECT officeCode, city, phone, addressLine1, addressLine2, state, country, postalCode, territory FROM offices"

	rows, _ := repository.Conn.Query(stmtOffices)
	var (
		OfficeCode  uint
		City        string
		Phone       string
		AdressLine1 string
		AdressLine2 sql.NullString
		State       sql.NullString
		Country     string
		PostalCode  string
		Territory   string
	)

	officesList := make([]Offices, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&OfficeCode, &City, &Phone, &AdressLine1, &AdressLine2, &State, &Country, &PostalCode, &Territory); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			employees, _ := repository.GetEmployeeByOffice(OfficeCode)
			office := Offices{
				OfficeCode:  OfficeCode,
				City:        City,
				Phone:       Phone,
				AdressLine1: AdressLine1,
				AdressLine2: AdressLine2.String,
				State:       State.String,
				Country:     Country,
				Employees:   employees,
			}
			officesList = append(officesList, office)
		default:
			return nil, err
		}
	}
	return officesList, nil
}

func (repository *Repository) GetOfficeAction(id uint64) ([]Employee, error) {
	sqlStmt := "SELECT employeeNumber, lastName, firstName, email, jobTitle FROM employees WHERE officeCode = ?"
	var (
		EmployeeNumber uint
		LastName       string
		FirstName      string
		Email          string
		JobTitle       string
	)

	rows, _ := repository.Conn.Query(sqlStmt, id)
	employeeList := make([]Employee, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&EmployeeNumber, &LastName, &FirstName, &Email, &JobTitle); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			employee := Employee{
				EmployeeNumber: EmployeeNumber,
				LastName:       LastName,
				FirstName:      FirstName,
				Email:          Email,
				JobTitle:       JobTitle,
			}
			employeeList = append(employeeList, employee)
		default:
			return nil, err
		}
	}
	return employeeList, nil
}
