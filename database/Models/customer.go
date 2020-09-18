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
	OrderNumber    uint           `json:"orderNumber"`
	OrderDate      time.Time      `json:"orderDate"`
	RequiredDate   time.Time      `json:"requiredDate"`
	ShippedDate    time.Time      `json:"shippedDate"`
	Status         string         `json:"status"`
	Comments       string         `json:"comments"`
	CustomerNumber uint           `json:"customerNumber"`
	Details        []OrderDetails `json:"details"`
}

type OrderDetails struct {
	OrderNumber        uint    `json:"orderNumber"`
	ProductCode        string  `json:"productCode"`
	QuantityOrdered    uint    `json:"quantity"`
	PriceEach          float64 `json:"price"`
	OrderLineNumber    uint16  `json:"orderLineNumber"`
	ProductName        string  `json:"productName"`
	ProductLine        string  `json:"productLine"`
	ProductScale       string  `json:"productScale"`
	ProductVendor      string  `json:"productVendor"`
	ProductDescription string  `json:"productDescription"`
	QuantityInStock    uint    `json:"stock"`
	BuyPrice           float64 `json:"buyPrice"`
	MSRP               float64 `json:"msrp"`
}

type OrderSummary struct {
	TotalPrice float64 `json:"totalPrice"`
	TotalItems int16   `json:"totalItems"`
	Order      []Order `json:"orders"`
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

func (repository *Repository) getTotalPriceOrder(id int64) (float64, error) {
	var totalPrice float64
	sqlStmt := "SELECT SUM(priceEach) as totalPrice FROM orders INNER JOIN orderdetails ON orders.orderNumber = orderdetails.orderNumber WHERE customerNumber = ?"
	err := repository.Conn.QueryRow(sqlStmt, id).Scan(&totalPrice)
	if err != nil {
		log.Fatal(err)
	}
	return totalPrice, nil
}

func (repository *Repository) getTotalItemsOrder(id int64) (int16, error) {
	var totalItems int16
	sqlStmt := "SELECT COUNT(priceEach) as totalItems FROM orders INNER JOIN orderdetails ON orders.orderNumber = orderdetails.orderNumber WHERE customerNumber = ?"
	err := repository.Conn.QueryRow(sqlStmt, id).Scan(&totalItems)
	if err != nil {
		log.Fatal(err)
	}
	return totalItems, nil
}

func (repository *Repository) getAllOrder(id int64) (int16, error) {
	var totalItems int16
	sqlStmt := "SELECT COUNT(priceEach) as totalItems FROM orders INNER JOIN orderdetails ON orders.orderNumber = orderdetails.orderNumber WHERE customerNumber = ?"
	err := repository.Conn.QueryRow(sqlStmt, id).Scan(&totalItems)
	if err != nil {
		log.Fatal(err)
	}
	return totalItems, nil
}

func (repository *Repository) getOrderDetails(id uint) ([]OrderDetails, error) {
	sqlStmtOrder := "SELECT orderNumber, products.productCode, quantityOrdered, priceEach, orderLineNumber, productName, productLine, productScale, productVendor, productDescription, quantityInStock, buyPrice, MSRP FROM orderdetails INNER JOIN products ON orderdetails.productCode = products.productCode WHERE orderNumber = ? ORDER BY orderLineNumber ASC"
	rows, _ := repository.Conn.Query(sqlStmtOrder, id)
	var (
		OrderNumber        uint
		ProductCode        string
		QuantityOrdered    uint
		PriceEach          float64
		OrderLineNumber    uint16
		productName        string
		productLine        string
		productScale       string
		productVendor      string
		productDescription string
		quantityInStock    uint
		buyPrice           float64
		MSRP               float64
	)

	orderDetails := make([]OrderDetails, 0)

	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&OrderNumber, &ProductCode, &QuantityOrdered, &PriceEach, &OrderLineNumber, &productName, &productLine, &productScale, &productVendor, &productDescription, &quantityInStock, &buyPrice, &MSRP); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			details := OrderDetails{
				OrderNumber:        OrderNumber,
				ProductCode:        ProductCode,
				QuantityOrdered:    QuantityOrdered,
				PriceEach:          PriceEach,
				OrderLineNumber:    OrderLineNumber,
				ProductName:        productName,
				ProductLine:        productLine,
				ProductScale:       productScale,
				ProductVendor:      productVendor,
				ProductDescription: productDescription,
				QuantityInStock:    quantityInStock,
				MSRP:               MSRP,
			}
			orderDetails = append(orderDetails, details)
		default:
			return nil, err
		}
	}
	return orderDetails, nil
}

func (repository *Repository) GetOrderByCustomer(id int64) (*OrderSummary, error) {
	totalPrice, _ := repository.getTotalPriceOrder(id)
	totalItems, _ := repository.getTotalItemsOrder(id)

	sqlStmtOrder := "SELECT orderNumber, orderDate, requiredDate, shippedDate, status, comments, customerNumber FROM orders WHERE customerNumber = ?"
	rows, _ := repository.Conn.Query(sqlStmtOrder, id)
	var (
		OrderNumber    uint
		OrderDate      time.Time
		RequiredDate   time.Time
		ShippedDate    time.Time
		Status         string
		Comments       sql.NullString
		CustomerNumber uint
	)

	orderList := make([]Order, 0)

	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&OrderNumber, &OrderDate, &RequiredDate, &ShippedDate, &Status, &Comments, &CustomerNumber); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			orderDetails, _ := repository.getOrderDetails(OrderNumber)
			order := Order{
				OrderNumber:    OrderNumber,
				OrderDate:      OrderDate,
				RequiredDate:   RequiredDate,
				ShippedDate:    ShippedDate,
				Status:         Status,
				Comments:       Comments.String,
				CustomerNumber: CustomerNumber,
				Details:        orderDetails,
			}
			orderList = append(orderList, order)
		default:
			return nil, err
		}
	}
	summary := OrderSummary{TotalPrice: totalPrice, TotalItems: totalItems, Order: orderList}
	return &summary, nil
}
