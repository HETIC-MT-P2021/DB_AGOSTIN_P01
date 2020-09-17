package Models

import (
	"database/sql"
	"log"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

type Customer struct {
	customerNumber         int
	customerName           string
	contactLastName        sql.NullString
	contactFirstName       sql.NullString
	phone                  sql.NullString
	addressLine1           sql.NullString
	addressLine2           sql.NullString
	city                   sql.NullString
	state                  sql.NullString
	postalCode             sql.NullString
	country                sql.NullString
	salesRepEmployeeNumber sql.NullInt64
	creditLimit            sql.NullFloat64
}

func (repository *Repository) GetUser() (*Customer, error) {
	rows, _ := repository.Conn.Query("SELECT customerNumber, customerName, contactLastName, contactFirstName, phone, addressLine1, addressLine2, city, state, postalCode, country, salesRepEmployeeNumber, creditLimit FROM customers")

	var (
		customerNumber         int
		customerName           string
		contactLastName        sql.NullString
		contactFirstName       sql.NullString
		phone                  sql.NullString
		addressLine1           sql.NullString
		addressLine2           sql.NullString
		city                   sql.NullString
		state                  sql.NullString
		postalCode             sql.NullString
		country                sql.NullString
		salesRepEmployeeNumber sql.NullInt64
		creditLimit            sql.NullFloat64
	)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1, &addressLine2, &city, &state, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit)
		if err != nil {
			log.Fatal(err)
		}
		switch err := rows.Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1, &addressLine2, &city, &state, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			customer := Customer{
				customerNumber:         customerNumber,
				customerName:           customerName,
				contactLastName:        contactLastName,
				contactFirstName:       contactFirstName,
				phone:                  phone,
				addressLine1:           addressLine1,
				addressLine2:           addressLine2,
				city:                   city,
				state:                  state,
				postalCode:             postalCode,
				country:                country,
				salesRepEmployeeNumber: salesRepEmployeeNumber,
				creditLimit:            creditLimit,
			}
			return &customer, nil
		default:
			return nil, err
		}
	}
	return nil, nil
}
