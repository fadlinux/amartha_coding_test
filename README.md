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
![Screen Shot 2024-07-02 at 09 03 26](https://github.com/fadlinux/amartha_coding_test/assets/416498/84cd197f-e311-407d-9c0d-b9d6f415fad8)



Create Loan POST /loans
```
curl --location 'http://localhost:8080/loans' \
--form 'cif_id="3"' \
--form 'total_amount="500000"' \
--form 'interest_rate="10"' \
--form 'total_weeks="50"'
```

![Screen Shot 2024-07-01 at 16 15 57](https://github.com/fadlinux/amartha_coding_test/assets/416498/f9b5e7df-db95-4bba-8c5c-0e8d74316a02)

Get Loan GET /loans
```
curl --location --request GET 'http://localhost:8080/loans?loan_id=4' 
```

![Screen Shot 2024-07-01 at 16 16 26](https://github.com/fadlinux/amartha_coding_test/assets/416498/2c2714bc-e530-4d0b-9ac2-eaa948fdadfa)


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
![Screen Shot 2024-07-01 at 16 16 40](https://github.com/fadlinux/amartha_coding_test/assets/416498/4d242623-6abc-41f4-b079-619976ae44a6)


## Import Postman Collection

```
Amartha_postman_collection.json
```
