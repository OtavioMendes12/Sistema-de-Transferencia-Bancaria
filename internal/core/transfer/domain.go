package transfer

type TransferHistory struct {
	ID        string  `bson:"_id,omitempty" json:"id"`
	FromID    string  `bson:"from_id" json:"from_id"`
	ToID      string  `bson:"to_id" json:"to_id"`
	Amount    float64 `bson:"amount" json:"amount"`
	Currency  string  `bson:"currency" json:"currency"`
	CreatedAt string  `bson:"created_at" json:"created_at"`
}
