package db

import (
	"bank-transfer-system/internal/core/user"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

// NewUserRepository cria um repositório para usuários
func NewUserRepository(database, collection string) user.Repository {
	return &userRepository{
		collection: GetCollection(database, collection),
	}
}

func (r *userRepository) GetAllUsers() ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []user.User
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Erro ao buscar usuários: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var u user.User
		if err := cursor.Decode(&u); err != nil {
			log.Printf("Erro ao decodificar usuário: %v", err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) CreateUser(u *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Remova qualquer ID do usuário para evitar duplicatas
	u.ID = ""

	_, err := r.collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	var u user.User
	// Converter o ID para ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	log.Printf("Atualizando usuário com ID: %s", u.ID)
	log.Printf("Novo saldo: %f", u.Balance)

	objectID, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return errors.New("ID inválido")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"balance": u.Balance}}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Erro ao atualizar usuário: %v", err)
		return errors.New("erro ao atualizar o usuário")
	}
	log.Println("Atualização bem-sucedida")
	return nil
}
