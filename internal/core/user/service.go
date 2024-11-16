package user

import (
	"bank-transfer-system/internal/core/currency"
	"bank-transfer-system/internal/core/transfer"
	transferCore "bank-transfer-system/internal/core/transfer"

	"context"
	"errors"
)

type service struct {
	repo            Repository
	currencyService currency.Service
	transferRepo    transferCore.Repository
}

func NewService(repo Repository, currencyService currency.Service, transferRepo transfer.Repository) Service {
	return &service{
		repo:            repo,
		currencyService: currencyService,
		transferRepo:    transferRepo,
	}
}

func (s *service) GetUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) AddUser(user *User) error {
	return s.repo.CreateUser(user)
}

const MaxTransferLimit = 10000.00
const TransferFeePercent = 0.01
const FixedTransferFee = 5.00

func (s *service) Transfer(ctx context.Context, transfer TransferDTO) error {
	if transfer.Amount <= 0 {
		return errors.New("valor da transferência deve ser maior que zero")
	}

	fee := transfer.Amount * TransferFeePercent
	if fee < FixedTransferFee {
		fee = FixedTransferFee
	}
	totalAmount := transfer.Amount + fee

	if totalAmount > MaxTransferLimit {
		return errors.New("valor da transferência excede o limite permitido")
	}

	if transfer.FromCurrency != transfer.ToCurrency {
		convertedAmount, err := s.currencyService.Convert(ctx, transfer.Amount, transfer.FromCurrency, transfer.ToCurrency)
		if err != nil {
			return errors.New("erro na conversão de moeda")
		}
		transfer.Amount = convertedAmount
	}

	fromUser, err := s.repo.FindByID(ctx, transfer.FromID)
	if err != nil || fromUser == nil {
		return errors.New("usuário remetente não encontrado")
	}

	if fromUser.Balance < totalAmount {
		return errors.New("saldo insuficiente para cobrir o valor e a taxa")
	}

	toUser, err := s.repo.FindByID(ctx, transfer.ToID)
	if err != nil || toUser == nil {
		return errors.New("usuário destinatário não encontrado")
	}

	fromUser.Balance -= totalAmount
	toUser.Balance += transfer.Amount

	if err := s.repo.Update(ctx, fromUser); err != nil {
		return errors.New("erro ao atualizar saldo do remetente")
	}
	if err := s.repo.Update(ctx, toUser); err != nil {
		return errors.New("erro ao atualizar saldo do destinatário")
	}

	history := transferCore.TransferHistory{
		FromID:   transfer.FromID,
		ToID:     transfer.ToID,
		Amount:   transfer.Amount,
		Currency: transfer.ToCurrency,
	}
	if err := s.transferRepo.SaveTransfer(ctx, &history); err != nil {
		return errors.New("erro ao salvar histórico de transferência")
	}

	return nil
}
