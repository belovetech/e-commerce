package utils

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request data")
	ErrInvalidOrderID = errors.New("invalid order ID")

	// cancel order
	ErrOrderAlreadyCancelled = errors.New("order already cancelled")
	ErrOrderNotPending       = errors.New("order not pending")
	ErrOrderNotFound         = errors.New("order not found")

	// create product
	ErrProductExists   = errors.New("product already exists")
	ErrProductNotFound = errors.New("product not found")
)
