package customer

import (
	"database/sql"
	"os"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/customer/http"
	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"

	mysqlCustomerRepository "github.com/fadlinux/amartha_coding_test/internal/repository/customer/mysql"

	postgre "github.com/fadlinux/amartha_coding_test/common/postgre"
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
)

var (
	//init construct var mysql
	postgreDb         *sql.DB
	mysqlCustomerRepo customerRepo.Repository

	customerUsecase uCustomer.Usecase

	//HTTPDelivery http handler customer
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	dbDsn := os.Getenv("DATABASE_URL")
	postgreDb = postgre.NewDBConnection("postgres", dbDsn)

	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLCustomerRepo(postgreDb)

	customerUsecase = uCustomer.NewCustomerUsecase(mysqlCustomerRepo)
	HTTPDelivery = httpAdmin.NewCustomerHTTP(customerUsecase)
}
