package errors

import (
	"net/http"

	e "delivery/pkg/errors"
)

var (
	ErrCustomerNotExists      = e.NewError(http.StatusNotFound, "customer not exists")
	ErrAccountNotExists       = e.NewError(http.StatusBadRequest, "account not exists")
	ErrCustomerAlreadyExists  = e.NewError(http.StatusBadRequest, "customer with this phone number already exists")
	ErrAccountAlreadyExists   = e.NewError(http.StatusBadRequest, "account with this info already exists")
	ErrInvalidTransactionType = e.NewError(http.StatusBadRequest, "invalid transaction type")

	ErrClinicAlreadyExists     = e.NewError(http.StatusBadRequest, "clinic already exists")
	ErrRegistrAlreadyExists    = e.NewError(http.StatusBadRequest, "registr already exists")
	ErrSubRegistrAlreadyExists = e.NewError(http.StatusBadRequest, "sub registr already exists")
	ErrClinicNotFound          = e.NewError(http.StatusNotFound, "clinic not exists")
	ErrRegistrNotFound         = e.NewError(http.StatusNotFound, "registr not exists")
	ErrSubRegistrNotExists     = e.NewError(http.StatusBadRequest, "sub registr not exists")

	ErrInvalidInput = e.NewError(http.StatusBadRequest, "invalid input")
)
