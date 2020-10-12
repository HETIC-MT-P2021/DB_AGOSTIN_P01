# Dashboard customer

This project aims to develop a API for a dashboard with #Goland

## Usage

```bash
docker-compose up
```

## Routes

```http request
Endpoint V1
GET    /v1 
```
```http request
Get all customers
GET    /v1/customers/
```

```http request
Get customer by ID
GET    /v1/customers/:id 
```

```http request
Get order by customer
GET    /v1/customers/:id/orders 
```

```http request
Get all orders
GET    /v1/orders/              
```

```http request
Get all employees
GET    /v1/employees/           
```

```http request
Get employees detail by ID
GET    /v1/employees/:id        
```

```http request
Get offices with employees
GET    /v1/offices/             
```

```http request
Get employees by office code
GET    /v1/offices/:id
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)