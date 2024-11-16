package transfer

import (
	"context"
)

type Service interface {
	SaveTransfer(ctx context.Context, t *TransferHistory) error
	GetTransfersByUserID(ctx context.Context, userID string) ([]TransferHistory, error)
}

type Repository interface {
	SaveTransfer(ctx context.Context, t *TransferHistory) error
	GetTransfersByUserID(ctx context.Context, userID string) ([]TransferHistory, error)
}
