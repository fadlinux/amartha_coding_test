package customer

import (
	"database/sql"
	"fmt"
	"os"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/customer/http"
	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"

	mysqlCustomerRepository "github.com/fadlinux/amartha_coding_test/internal/repository/customer/mysql"

	mysql "github.com/fadlinux/amartha_coding_test/common/mysql"
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
)

var (
	//init construct var mysql
	mySqlDB           *sql.DB
	mysqlCustomerRepo customerRepo.Repository

	customerUsecase uCustomer.Usecase

	//HTTPDelivery http handler customer
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	//root@tcp(localhost:3306)/amartha
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	mySqlDB = mysql.NewDBConnection("mysql", dsn)

	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLCustomerRepo(mySqlDB)
	customerUsecase = uCustomer.NewCustomerUsecase(mysqlCustomerRepo)

	HTTPDelivery = httpAdmin.NewCustomerHTTP(customerUsecase)
}
