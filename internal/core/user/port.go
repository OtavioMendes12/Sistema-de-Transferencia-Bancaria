package user

import "context"

// Service define as operações do serviço de usuários
type Service interface {
	GetUsers() ([]User, error)                                // Buscar todos os usuários
	AddUser(user *User) error                                 // Adicionar um novo usuário
	Transfer(ctx context.Context, transfer TransferDTO) error // Realizar transferência de saldo
}

// Repository define as operações no banco de dados
type Repository interface {
	GetAllUsers() ([]User, error)                           // Buscar todos os usuários
	CreateUser(user *User) error                            // Criar um novo usuário
	FindByID(ctx context.Context, id string) (*User, error) // Buscar usuário por ID
	Update(ctx context.Context, user *User) error           // Atualizar um usuário
}
