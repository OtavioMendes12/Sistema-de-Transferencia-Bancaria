package user

type CreateUserDTO struct {
	ID      string  `bson:"_id,omitempty" json:"id"` //
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

type TransferDTO struct {
	FromID string  `json:"from_id"` // ID do remetente
	ToID   string  `json:"to_id"`   // ID do destinatário
	Amount float64 `json:"amount"`  // Valor da transferência
}
