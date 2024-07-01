package loan

import (
	"database/sql"
	"fmt"
	"os"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/loan/http"
	loanRepo "github.com/fadlinux/amartha_coding_test/internal/repository/loan"

	mysqlLoanRepository "github.com/fadlinux/amartha_coding_test/internal/repository/loan/mysql"

	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"

	mysqlCustomerRepository "github.com/fadlinux/amartha_coding_test/internal/repository/customer/mysql"

	mysql "github.com/fadlinux/amartha_coding_test/common/mysql"
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
)

var (
	//init construct var mysql
	mySqlDB       *sql.DB
	mysqlloanRepo loanRepo.Repository
	loanUsecase   uLoan.Usecase

	customerUsecase   uCustomer.Usecase
	mysqlCustomerRepo customerRepo.Repository

	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	mySqlDB = mysql.NewDBConnection("mysql", dsn)

	//mySqlDB = mysql.NewDBConnection("mysql", configCmd.MysqlDBConnection)
	mysqlloanRepo = mysqlLoanRepository.NewMySQLLoanRepo(mySqlDB)
	loanUsecase = uLoan.NewLoanUsecase(mysqlloanRepo)

	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLCustomerRepo(mySqlDB)
	customerUsecase = uCustomer.NewCustomerUsecase(mysqlCustomerRepo)

	HTTPDelivery = httpAdmin.NewLoanHTTP(loanUsecase, customerUsecase)
}
