package Models

import (
	"database/sql"
	"time"
)

type ItemsDetails struct {
	OrderNumber        uint      `json:"orderNumber"`
	ProductCode        string    `json:"productCode"`
	QuantityOrdered    uint      `json:"quantity"`
	PriceEach          float64   `json:"price"`
	OrderLineNumber    uint16    `json:"orderLineNumber"`
	ProductName        string    `json:"productName"`
	ProductLine        string    `json:"productLine"`
	ProductScale       string    `json:"productScale"`
	ProductVendor      string    `json:"productVendor"`
	ProductDescription string    `json:"productDescription"`
	QuantityInStock    uint      `json:"stock"`
	BuyPrice           float64   `json:"buyPrice"`
	MSRP               float64   `json:"msrp"`
	OrderDate          time.Time `json:"orderDate"`
	RequiredDate       time.Time `json:"requiredDate"`
	ShippedDate        time.Time `json:"shippedDate"`
	Status             string    `json:"status"`
	Comments           string    `json:"comments"`
	CustomerNumber     uint      `json:"customerNumber"`
}

func (repository *Repository) GetAllItemsInOrder() ([]ItemsDetails, error) {
	sqlStmt := "SELECT orders.orderNumber, orderDate, requiredDate, shippedDate, status, comments, customerNumber, products.productCode, quantityOrdered, priceEach, orderLineNumber, productName, productLine, productScale, productVendor, productDescription, quantityInStock, buyPrice, MSRP FROM orders INNER JOIN orderdetails ON orders.orderNumber = orderdetails.orderNumber INNER JOIN products ON orderdetails.productCode = products.productCode"
	rows, _ := repository.Conn.Query(sqlStmt)
	var (
		OrderNumber        uint
		ProductCode        string
		QuantityOrdered    uint
		PriceEach          float64
		OrderLineNumber    uint16
		ProductName        string
		ProductLine        string
		ProductScale       string
		ProductVendor      string
		ProductDescription string
		QuantityInStock    uint
		BuyPrice           float64
		MSRP               float64
		OrderDate          time.Time
		RequiredDate       time.Time
		ShippedDate        sql.NullTime
		Status             string
		Comments           sql.NullString
		CustomerNumber     uint
	)

	itemsDetails := make([]ItemsDetails, 0)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&OrderNumber, &OrderDate, &RequiredDate, &ShippedDate, &Status, &Comments, &CustomerNumber, &ProductCode, &QuantityOrdered, &PriceEach, &OrderLineNumber, &ProductName, &ProductLine, &ProductScale, &ProductVendor, &ProductDescription, &QuantityInStock, &BuyPrice, &MSRP); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			details := ItemsDetails{
				OrderNumber:        OrderNumber,
				ProductCode:        ProductCode,
				QuantityOrdered:    QuantityOrdered,
				PriceEach:          PriceEach,
				OrderLineNumber:    OrderLineNumber,
				ProductName:        ProductName,
				ProductLine:        ProductLine,
				ProductScale:       ProductScale,
				ProductVendor:      ProductVendor,
				ProductDescription: ProductDescription,
				QuantityInStock:    QuantityInStock,
				BuyPrice:           BuyPrice,
				MSRP:               MSRP,
				OrderDate:          OrderDate,
				RequiredDate:       RequiredDate,
				ShippedDate:        ShippedDate.Time,
				Status:             Status,
				Comments:           Comments.String,
				CustomerNumber:     CustomerNumber,
			}
			itemsDetails = append(itemsDetails, details)
		default:
			return nil, err
		}
	}
	return itemsDetails, nil
}
