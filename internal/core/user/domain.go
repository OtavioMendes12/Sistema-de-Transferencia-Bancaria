package user

type User struct {
	ID      string  `bson:"_id,omitempty" json:"id"`
	Name    string  `bson:"name" json:"name"`
	Email   string  `bson:"email" json:"email"`
	Balance float64 `bson:"balance" json:"balance"`
}
