package payment

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/fadlinux/amartha_coding_test/common/postgre"
	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/payment/http"
	loanRepo "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
	paymentRepo "github.com/fadlinux/amartha_coding_test/internal/repository/payment"

	mysqlLoanRepository "github.com/fadlinux/amartha_coding_test/internal/repository/loan/mysql"
	mysqlPaymentRepository "github.com/fadlinux/amartha_coding_test/internal/repository/payment/mysql"

	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
	uPayment "github.com/fadlinux/amartha_coding_test/internal/usecase/payment"
)

var (
	//init construct var mysql
	postgreDb        *sql.DB
	mysqlpaymentRepo paymentRepo.Repository
	paymentUsecase   uPayment.Usecase

	mysqlLoanRepo loanRepo.Repository
	loanUsecase   uLoan.Usecase
	//HTTPDelivery http handler payment
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	dbDsn := os.Getenv("DATABASE_URL")
	postgreDb = postgre.NewDBConnection("postgres", dbDsn)

	mysqlpaymentRepo = mysqlPaymentRepository.NewMySQLPaymentRepo(postgreDb)
	paymentUsecase = uPayment.NewPaymentUsecase(mysqlpaymentRepo)

	mysqlLoanRepo = mysqlLoanRepository.NewMySQLLoanRepo(postgreDb)
	loanUsecase = uLoan.NewLoanUsecase(mysqlLoanRepo)

	HTTPDelivery = httpAdmin.NewPaymentHTTP(paymentUsecase, loanUsecase)
}
