# Amartha Coding Test



## How To run

```
go mod tidy
go mod vendor
```

1. Run docker-compose

```
docker compose up --build -d  

```

2. Stop and remove docker-compose volume


```
docker compose down --volumes

```

## Endpoint API

Create Customer POST /customer
```
curl --location 'http://localhost:8080/customer' \
--form 'fullname="fadli"' \
--form 'ktp_no="123456789"' \
--form 'address="jakarta"'
```


Create Loan POST /loans
```
curl --location 'http://localhost:8080/loans' \
--form 'cif_id="3"' \
--form 'total_amount="500000"' \
--form 'interest_rate="10"' \
--form 'total_weeks="50"'
```

Get Loan GET /loans
```
curl --location --request GET 'http://localhost:8080/loans?loan_id=4' 
```

Update Loan delinquent PUT /loans
```
curl --location --request PUT 'http://localhost:8080/loans' \
--form 'loan_id="4"' \
--form 'delinquent="2"'
```

Create Payment POST /payment
```

curl --location 'http://localhost:8080/payment' \
--form 'loan_id="4"' \
--form 'amount="11000"
```


## Import Postman Collection

```
Amartha_postman_collection.json
```
