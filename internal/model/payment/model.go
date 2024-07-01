package payment

type (
	AddPaymentRequest struct {
		PaymentID   string `json:"payment_id"`
		LoanID      string `json:"loan_id"`
		Amount      int64  `json:"amount"`
		PaymentDate string `json:"payment_date"`
	}

	AddPaymentResponse struct {
		Message   string          `json:"message"`
		PaymentID string          `json:"payment_id"`
		Payment   PaymentResponse `json:"loan,omitempty"`
	}

	PaymentResponse struct {
		TotalAmount   int64     `json:"total_amount"`
		InterestRate  int64     `json:"interest_rate"`
		TotalWeeks    int64     `json:"total_weeks"`
		WeeklyPayment float32   `json:"weekly_payment"`
		WeeksPayment  []float64 `json:"weeks_payment"`
		OutStanding   int64     `json:"outstanding"`
		Status        string    `json:"status"`
	}

	TotalPaymentResponse struct {
		TotalPayment int64 `json:"total_payment"`
	}
)
