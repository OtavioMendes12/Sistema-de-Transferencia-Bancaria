package user

type User struct {
	ID      string  `bson:"_id,omitempty" json:"id"`
	Name    string  `bson:"name" json:"name"`
	Email   string  `bson:"email" json:"email"`
	Balance float64 `bson:"balance" json:"balance"`
}

type TransferHistory struct {
	ID        string  `bson:"_id,omitempty" json:"id"`
	FromID    string  `json:"from_id"`
	ToID      string  `json:"to_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
}
