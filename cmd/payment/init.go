package payment

import (
	"database/sql"
	"fmt"
	"os"

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
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST") //if not run by pass to "172.31.0.2"
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	mySqlDB = mysql.NewDBConnection("mysql", dsn)

	//	mySqlDB = mysql.NewDBConnection("mysql", configCmd.MysqlDBConnection)
	mysqlpaymentRepo = mysqlPaymentRepository.NewMySQLPaymentRepo(mySqlDB)
	paymentUsecase = uPayment.NewPaymentUsecase(mysqlpaymentRepo)

	mysqlLoanRepo = mysqlLoanRepository.NewMySQLLoanRepo(mySqlDB)
	loanUsecase = uLoan.NewLoanUsecase(mysqlLoanRepo)

	HTTPDelivery = httpAdmin.NewPaymentHTTP(paymentUsecase, loanUsecase)
}
