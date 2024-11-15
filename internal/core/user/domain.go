package user

type User struct {
	ID      string  `bson:"_id,omitempty" json:"id"` // "_id" é gerado automaticamente pelo MongoDB
	Name    string  `bson:"name" json:"name"`
	Email   string  `bson:"email" json:"email"`
	Balance float64 `bson:"balance" json:"balance"`
}
