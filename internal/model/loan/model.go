package loan

type (
	AddLoanRequest struct {
		CifID         string  `json:"cif_id"`
		TotalAmount   int64   `json:"total_amount"`
		InterestRate  int64   `json:"interest_rate"`
		TotalWeeks    int64   `json:"total_weeks"`
		WeeklyPayment float32 `json:"weekly_payment"`
		Status        string  `json:"status"`
	}

	AddLoanResponse struct {
		Message      string       `json:"message"`
		LoanID       string       `json:"loan_id"`
		CustomerData CustomerData `json:"customer_data"`
		Loan         LoanResponse `json:"loan"`
	}

	LoanResponse struct {
		TotalAmount   int64     `json:"total_amount"`
		InterestRate  int64     `json:"interest_rate"`
		TotalWeeks    int64     `json:"total_weeks"`
		WeeklyPayment float32   `json:"weekly_payment"`
		WeeksPayment  []float64 `json:"weeks_payment"`
		OutStanding   int64     `json:"outstanding"`
		Status        string    `json:"status"`
	}

	CustomerData struct {
		FullName string `json:"fullname"`
		KTP      string `json:"ktp_no"`
		Address  string `json:"address"`
	}
)
