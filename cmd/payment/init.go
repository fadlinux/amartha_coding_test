package payment

import (
	"database/sql"

	configCmd "github.com/fadlinux/amartha_coding_test/cmd/config"

	httpAdmin "github.com/fadlinux/amartha_coding_test/internal/delivery/payment/http"
	loanRepo "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
	paymentRepo "github.com/fadlinux/amartha_coding_test/internal/repository/payment"

	mysqlLoanRepository "github.com/fadlinux/amartha_coding_test/internal/repository/loan/mysql"
	mysqlPaymentRepository "github.com/fadlinux/amartha_coding_test/internal/repository/payment/mysql"

	mysql "github.com/fadlinux/amartha_coding_test/common/mysql"
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
	uPayment "github.com/fadlinux/amartha_coding_test/internal/usecase/payment"
)

var (
	//init construct var mysql
	mySqlDB          *sql.DB
	mysqlpaymentRepo paymentRepo.Repository
	paymentUsecase   uPayment.Usecase

	mysqlLoanRepo loanRepo.Repository
	loanUsecase   uLoan.Usecase
	//HTTPDelivery http handler payment
	HTTPDelivery httpAdmin.Delivery
)

func Initialize() {
	mySqlDB = mysql.NewDBConnection("mysql", configCmd.MysqlDBConnection)
	mysqlpaymentRepo = mysqlPaymentRepository.NewMySQLPaymentRepo(mySqlDB)
	paymentUsecase = uPayment.NewPaymentUsecase(mysqlpaymentRepo)

	mysqlLoanRepo = mysqlLoanRepository.NewMySQLLoanRepo(mySqlDB)
	loanUsecase = uLoan.NewLoanUsecase(mysqlLoanRepo)

	HTTPDelivery = httpAdmin.NewPaymentHTTP(paymentUsecase, loanUsecase)
}
