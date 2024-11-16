package user

type CreateUserDTO struct {
	ID      string  `bson:"_id,omitempty" json:"id"` //
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

type TransferDTO struct {
	FromID       string  `json:"from_id"`
	ToID         string  `json:"to_id"`
	Amount       float64 `json:"amount"`
	FromCurrency string  `json:"from_currency"` // Moeda do remetente
	ToCurrency   string  `json:"to_currency"`   // Moeda do destinatário
}
