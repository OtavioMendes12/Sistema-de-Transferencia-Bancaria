package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectDatabase conecta ao MongoDB
func ConnectDatabase(uri string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o MongoDB: %v", err)
	}

	log.Println("Conexão com o MongoDB estabelecida com sucesso!")
	MongoClient = client
}

func GetCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
