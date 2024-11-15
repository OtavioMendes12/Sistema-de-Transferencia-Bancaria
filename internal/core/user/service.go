package user

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repo Repository
}

// NewService cria uma nova instância do serviço e retorna como Service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// GetUsers retorna todos os usuários
func (s *service) GetUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

// AddUser adiciona um novo usuário
func (s *service) AddUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *service) Transfer(ctx context.Context, transfer TransferDTO) error {
	log.Printf("Transferência: remetente=%s, destinatário=%s, valor=%f",
		transfer.FromID, transfer.ToID, transfer.Amount)
	if transfer.Amount <= 0 {
		return errors.New("valor da transferência deve ser maior que zero")
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
	err = s.repo.Update(ctx, fromUser)
	if err != nil {
		return errors.New("erro ao atualizar saldo do remetente")
	}

	err = s.repo.Update(ctx, toUser)
	if err != nil {
		return errors.New("erro ao atualizar saldo do destinatário")
	}

	return nil
}
