package currency

import "context"

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Convert(ctx context.Context, amount float64, fromCurrency, toCurrency string) (float64, error) {
	rate, err := s.repo.GetRate(ctx, fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}

	if rate == 0 {
		return 0, nil
	}

	return amount * rate, nil
}
