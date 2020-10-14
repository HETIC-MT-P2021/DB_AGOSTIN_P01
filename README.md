<h1 align="center">Dashboard SHOP API</h1>


<p align="center">
    <a href="https://travis-ci.org/schollz/croc"><img
    src="https://img.shields.io/travis/schollz/croc.svg?style=flat-square" alt="Build
    Status"></a> 
</p>
      
<p align="center">
  <a href="#about">About</a> •
  <a href="#lipsum2">How To use ?</a> •
  <a href="#features">Features</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#author">Author</a> •
  <a href="#support">Support</a> •
  <a href="#donate">Donate</a> •
  <a href="#license">License</a>
</p>

---

## About

<table>
<tr>
<td>
  
  This is a simple API rest for serving a dashboard in Front. 
  
</td>
</tr>
</table>

## How to use ?

Before cloning make sure you have docker installed. Then clone the project and follow the steps:

##### Run the project

```bash
docker-compose up
```


## Features

```http request
Endpoint V1
GET    /api/v1/
```
```http request
Get all customers
GET    /api/v1/customers/
```

```http request
Get customer by ID
GET    /api/v1/customers/:id 
```

```http request
Get order by customer
GET    /api/v1/customers/:id/orders 
```

```http request
Get all orders
GET    /api/v1/orders/              
```

```http request
Get all employees
GET    /api/v1/employees/           
```

```http request
Get employees detail by ID
GET    /api/v1/employees/:id        
```

```http request
Get offices with employees
GET    /api/v1/offices/             
```

```http request
Get employees by office code
GET    /api/v1/offices/:id
```

## Contributing

Got **something interesting** you'd like to **share**? Learn about [contributing](https://github.com/HETIC-MT-P2021/DB_AGOSTIN_P01/blob/master/contributing.md).

## Author

| [![Jibe](https://i.kym-cdn.com/photos/images/newsfeed/001/196/011/332.jpg)](https://www.linkedin.com/in/jbagostin/) 	|
|:---------------------------------------------------------------------------------------------------------:	|
|                                            **Agostin Jean-baptiste**                                            	|

## Support

Reach out to me at one of the following places:

- E-Mail: **Jbagostin@gmail.com**

## Donate

[![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg)](YOUR_EMAIL_CODE)

## License

[![License: CC BY-NC-SA 4.0](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://creativecommons.org/licenses/by-nc-sa/4.0/)

- Copyright © [Jibe]().