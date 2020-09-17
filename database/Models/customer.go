package Models

import (
	"database/sql"
	"log"
	"time"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

type Customer struct {
	CustomerNumber         int     `json:"customerNumber"`
	CustomerName           string  `json:"customerName"`
	ContactLastName        string  `json:"contactLastName"`
	ContactFirstName       string  `json:"contactFirstName"`
	Phone                  string  `json:"phone"`
	AddressLine1           string  `json:"addressLine1"`
	AddressLine2           string  `json:"addressLine2"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	PostalCode             string  `json:"postalCode"`
	Country                string  `json:"country"`
	SalesRepEmployeeNumber int64   `json:"salesRepEmployeeNumber"`
	CreditLimit            float64 `json:"creditLimit"`
}

type Order struct {
	OrderNumber    int       `json:"orderNumber"`
	OrderDate      time.Time `json:"orderDate"`
	RequiredDate   time.Time `json:"requiredDate"`
	ShippedDate    time.Time `json:"shippedDate"`
	Status         string    `json:"status"`
	Comments       string    `json:"comments"`
	CustomerNumber int       `json:"customerNumber"`
}

func (repository *Repository) GetAllCustomers() ([]Customer, error) {
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

	var customersList []Customer

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
				CustomerNumber:         customerNumber,
				CustomerName:           customerName,
				ContactLastName:        contactLastName.String,
				ContactFirstName:       contactFirstName.String,
				Phone:                  phone.String,
				AddressLine1:           addressLine1.String,
				AddressLine2:           addressLine2.String,
				City:                   city.String,
				State:                  state.String,
				PostalCode:             postalCode.String,
				Country:                country.String,
				SalesRepEmployeeNumber: salesRepEmployeeNumber.Int64,
				CreditLimit:            creditLimit.Float64,
			}
			customersList = append(customersList, customer)
		default:
			return nil, err
		}
	}
	return customersList, nil
}

func (repository *Repository) GetCustomer(id int64) (*Customer, error) {

	sqlStmt := "SELECT customerNumber, customerName, contactLastName, contactFirstName, phone, addressLine1, addressLine2, city, state, postalCode, country, salesRepEmployeeNumber, creditLimit FROM customers WHERE customerNumber = ?"
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
	stmt, err := repository.Conn.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1, &addressLine2, &city, &state, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit)
	if err != nil {
		log.Fatal(err)
	}
	customer := Customer{
		CustomerNumber:         customerNumber,
		CustomerName:           customerName,
		ContactLastName:        contactLastName.String,
		ContactFirstName:       contactFirstName.String,
		Phone:                  phone.String,
		AddressLine1:           addressLine1.String,
		AddressLine2:           addressLine2.String,
		City:                   city.String,
		State:                  state.String,
		PostalCode:             postalCode.String,
		Country:                country.String,
		SalesRepEmployeeNumber: salesRepEmployeeNumber.Int64,
		CreditLimit:            creditLimit.Float64,
	}
	return &customer, nil
}

/*func (repository *Repository) GetOrderByCustomer(id int64) ([]Customer, error) {
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

	var customersList []Customer

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
				CustomerNumber:         customerNumber,
				CustomerName:           customerName,
				ContactLastName:        contactLastName.String,
				ContactFirstName:       contactFirstName.String,
				Phone:                  phone.String,
				AddressLine1:           addressLine1.String,
				AddressLine2:           addressLine2.String,
				City:                   city.String,
				State:                  state.String,
				PostalCode:             postalCode.String,
				Country:                country.String,
				SalesRepEmployeeNumber: salesRepEmployeeNumber.Int64,
				CreditLimit:            creditLimit.Float64,
			}
			customersList = append(customersList, customer)
		default:
			return nil, err
		}
	}
	return customersList, nil
}
*/
