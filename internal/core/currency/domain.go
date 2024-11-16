package currency

type CurrencyRate struct {
	ID           string  `bson:"_id,omitempty" json:"id"`
	FromCurrency string  `bson:"from_currency" json:"from_currency"`
	ToCurrency   string  `bson:"to_currency" json:"to_currency"`
	Rate         float64 `bson:"rate" json:"rate"`
}
