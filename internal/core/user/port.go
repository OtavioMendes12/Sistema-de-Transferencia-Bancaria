package user

import "context"

type Service interface {
	GetUsers() ([]User, error)
	AddUser(user *User) error
	Transfer(ctx context.Context, transfer TransferDTO) error
}

// Repository define as operações no banco de dados
type Repository interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
}
