package customer

type (
	AddCustomerRequest struct {
		Fullname string `json:"fullname"`
		KTP      string `json:"ktp_no"`
		Address  string `json:"address"`
	}

	AddCustomerResponse struct {
		Message string `json:"message"`
		CifID   int64  `json:"cif_id,omitempty"`
	}
)
