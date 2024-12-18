package admin

import (
	"context"
	"delivery/configs"
	"delivery/entities"
	"delivery/logger"
	"math"
	"time"
)

type AdminController interface {
	CalculateCredit(ctx context.Context, req entities.CalculateModel) ([]entities.Payment, error)
}

type adminController struct {
	log logger.LoggerI
	cfg *configs.Configuration
}

func NewAdminController(log logger.LoggerI) AdminController {
	return adminController{
		log: log,
		cfg: configs.Config(),
	}
}

func (a adminController) CalculateCredit(ctx context.Context, req entities.CalculateModel) ([]entities.Payment, error) {
	a.log.Info("CalculateCredit started: ")

	var payments []entities.Payment

	monthlyRate := req.AnnualRate / 12 / 100 // Oylik foiz stavkasi
	remainingDebt := req.Principal           // Qolgan qarz
	var totalPayment float64

	for i := 1; i <= int(req.Months); i++ {
		var interest, principalPayment, payment float64

		if req.CreditType == "annuitet" {
			annuityCoefficient := (monthlyRate * math.Pow(1+monthlyRate, float64(req.Months))) / (math.Pow(1+monthlyRate, float64(req.Months)) - 1)
			payment = req.Principal * annuityCoefficient
			interest = remainingDebt * monthlyRate
			principalPayment = payment - interest
		} else if req.CreditType == "differential" {
			// Differensial to'lov
			principalPayment = req.Principal / float64(req.Months)
			interest = remainingDebt * monthlyRate
			payment = principalPayment + interest
		}

		// Qolgan qarzni hisoblash
		remainingDebt -= principalPayment
		if remainingDebt < 0 {
			remainingDebt = 0
		}

		paymentDate := time.Now().AddDate(0, i, 0).Format("2006-01-02")

		// To'lovni qo'shish
		payments = append(payments, entities.Payment{
			Number:        i,
			Date:          paymentDate,
			Principal:     math.Round(principalPayment*100) / 100,
			Interest:      math.Round(interest*100) / 100,
			TotalPayment:  math.Round(payment*100) / 100,
			RemainingDebt: math.Round(remainingDebt*100) / 100,
		})

		totalPayment += payment
	}

	a.log.Info("CalculateCredit finished")

	return payments, nil
}
