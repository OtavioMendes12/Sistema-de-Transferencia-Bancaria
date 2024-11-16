package db

import (
	"bank-transfer-system/internal/core/currency"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type currencyRepository struct {
	collection *mongo.Collection
}

func NewCurrencyRepository(database, collection string) currency.Repository {
	return &currencyRepository{
		collection: GetCollection(database, collection),
	}
}

func (r *currencyRepository) GetRate(ctx context.Context, fromCurrency, toCurrency string) (float64, error) {
	var rate currency.CurrencyRate
	filter := bson.M{
		"from_currency": fromCurrency,
		"to_currency":   toCurrency,
	}

	err := r.collection.FindOne(ctx, filter).Decode(&rate)
	if err == mongo.ErrNoDocuments {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return rate.Rate, nil
}
