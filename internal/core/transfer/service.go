package transfer

import (
	"context"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SaveTransfer(ctx context.Context, t *TransferHistory) error {
	return s.repo.SaveTransfer(ctx, t)
}

func (s *service) GetTransfersByUserID(ctx context.Context, userID string) ([]TransferHistory, error) {
	return s.repo.GetTransfersByUserID(ctx, userID)
}
