package loan

import (
	"database/sql"
	"os"

	"github.com/fadlinux/amartha_coding_test/common/postgre"
	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/loan/http"
	customerRepo "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
	loanRepo "github.com/fadlinux/amartha_coding_test/internal/repository/loan"

	mysqlCustomerRepository "github.com/fadlinux/amartha_coding_test/internal/repository/customer/mysql"
	mysqlLoanRepository "github.com/fadlinux/amartha_coding_test/internal/repository/loan/mysql"

	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
)

var (
	//init construct var mysql
	postgreDb     *sql.DB
	mysqlloanRepo loanRepo.Repository
	loanUsecase   uLoan.Usecase

	mysqlCustomerRepo customerRepo.Repository
	customerUsecase   uCustomer.Usecase
	//HTTPDelivery http handler loan
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	dbDsn := os.Getenv("DATABASE_URL")
	postgreDb = postgre.NewDBConnection("postgres", dbDsn)

	mysqlloanRepo = mysqlLoanRepository.NewMySQLLoanRepo(postgreDb)
	loanUsecase = uLoan.NewLoanUsecase(mysqlloanRepo)

	mysqlCustomerRepo = mysqlCustomerRepository.NewMySQLCustomerRepo(postgreDb)
	customerUsecase = uCustomer.NewCustomerUsecase(mysqlCustomerRepo)

	HTTPDelivery = httpAdmin.NewLoanHTTP(loanUsecase, customerUsecase)
}
