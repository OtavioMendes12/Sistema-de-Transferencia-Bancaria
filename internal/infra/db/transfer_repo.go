package db

import (
	"bank-transfer-system/internal/core/transfer"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type transferRepository struct {
	collection *mongo.Collection
}

var _ transfer.Repository = (*transferRepository)(nil)

func NewTransferRepository(database, collection string) transfer.Repository {
	return &transferRepository{
		collection: GetCollection(database, collection),
	}
}

func (r *transferRepository) SaveTransfer(ctx context.Context, t *transfer.TransferHistory) error {
	t.CreatedAt = time.Now().Format(time.RFC3339)
	_, err := r.collection.InsertOne(ctx, t)
	return err
}

func (r *transferRepository) GetTransfersByUserID(ctx context.Context, userID string) ([]transfer.TransferHistory, error) {
	var transfers []transfer.TransferHistory
	filter := bson.M{"$or": []bson.M{
		{"from_id": userID},
		{"to_id": userID},
	}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var t transfer.TransferHistory
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}
		transfers = append(transfers, t)
	}

	return transfers, nil
}
