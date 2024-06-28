package customer

import (
	"database/sql"

	configCmd "github.com/fadlinux/amartha_coding_test/cmd/config"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/customer/http"
	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"

	mysqlCustomerRepository "flip.id/internal/repository/customer/mysql"

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
	mySqlDB = mysql.NewDBConnection("mysql", configCmd.MysqlDBConnection)
	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLcustomerRepo(mySqlDB)

	customerUsecase = uCustomer.NewcustomerUsecase(mysqlCustomerRepo)

	HTTPDelivery = httpAdmin.NewcustomerHTTP(customerUsecase)
}
