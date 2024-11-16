package user

import (
	"bank-transfer-system/internal/core/currency"
	"context"
	"errors"
)

type service struct {
	repo            Repository
	currencyService currency.Service
}

// NewService cria um novo serviço de usuários
func NewService(repo Repository, currencyService currency.Service) Service {
	return &service{
		repo:            repo,
		currencyService: currencyService,
	}
}

// GetUsers retorna todos os usuários
func (s *service) GetUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

// AddUser adiciona um novo usuário
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

	// Verificar se as moedas são diferentes
	if transfer.FromCurrency != transfer.ToCurrency {
		convertedAmount, err := s.currencyService.Convert(ctx, transfer.Amount, transfer.FromCurrency, transfer.ToCurrency)
		if err != nil {
			return errors.New("erro na conversão de moeda")
		}

		transfer.Amount = convertedAmount
	}

	// Buscar o remetente
	fromUser, err := s.repo.FindByID(ctx, transfer.FromID)
	if err != nil || fromUser == nil {
		return errors.New("usuário remetente não encontrado")
	}

	// Verificar saldo suficiente
	if fromUser.Balance < transfer.Amount {
		return errors.New("saldo insuficiente")
	}

	// Buscar o destinatário
	toUser, err := s.repo.FindByID(ctx, transfer.ToID)
	if err != nil || toUser == nil {
		return errors.New("usuário destinatário não encontrado")
	}

	// Atualizar os saldos
	fromUser.Balance -= transfer.Amount
	toUser.Balance += transfer.Amount

	// Persistir as alterações
	if err := s.repo.Update(ctx, fromUser); err != nil {
		return errors.New("erro ao atualizar saldo do remetente")
	}

	if err := s.repo.Update(ctx, toUser); err != nil {
		return errors.New("erro ao atualizar saldo do destinatário")
	}

	return nil
}
