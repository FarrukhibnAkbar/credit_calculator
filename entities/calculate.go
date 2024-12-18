package entities

import "github.com/google/uuid"

type CalculateModel struct {
	ID         uuid.UUID `json:"-"`
	Principal  float64   `json:"principal" validate:"required,gt=0"`
	AnnualRate float64   `json:"annual_rate" validate:"required,gt=0"`
	Months     uint8     `json:"months" validate:"required,gt=0"`
	CreditType string    `json:"credit_type" validate:"required"`
}

type Payment struct {
	Number        int     `json:"number"`
	Date          string  `json:"date"`
	Interest      float64 `json:"interest"`
	Principal     float64 `json:"principal"`
	TotalPayment  float64 `json:"total_payment"`
	RemainingDebt float64 `json:"remaining_debt"`
}
