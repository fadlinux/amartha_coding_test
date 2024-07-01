package loan

import (
	"database/sql"

	configCmd "github.com/fadlinux/amartha_coding_test/cmd/config"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/loan/http"
	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
	loanRepo "github.com/fadlinux/amartha_coding_test/internal/repository/loan"

	mysqlCustomerRepository "github.com/fadlinux/amartha_coding_test/internal/repository/customer/mysql"
	mysqlLoanRepository "github.com/fadlinux/amartha_coding_test/internal/repository/loan/mysql"

	mysql "github.com/fadlinux/amartha_coding_test/common/mysql"
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
)

var (
	//init construct var mysql
	mySqlDB       *sql.DB
	mysqlloanRepo loanRepo.Repository
	loanUsecase   uLoan.Usecase

	mysqlCustomerRepo customerRepo.Repository
	customerUsecase   uCustomer.Usecase
	//HTTPDelivery http handler loan
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	mySqlDB = mysql.NewDBConnection("mysql", configCmd.MysqlDBConnection)
	mysqlloanRepo = mysqlLoanRepository.NewMySQLLoanRepo(mySqlDB)
	loanUsecase = uLoan.NewLoanUsecase(mysqlloanRepo)

	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLCustomerRepo(mySqlDB)
	customerUsecase = uCustomer.NewCustomerUsecase(mysqlCustomerRepo)

	HTTPDelivery = httpAdmin.NewLoanHTTP(loanUsecase, customerUsecase)
}
