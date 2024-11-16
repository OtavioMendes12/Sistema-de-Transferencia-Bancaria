package currency

import "context"

type Repository interface {
	GetRate(ctx context.Context, fromCurrency, toCurrency string) (float64, error)
}

type Service interface {
	Convert(ctx context.Context, amount float64, fromCurrency, toCurrency string) (float64, error)
}
