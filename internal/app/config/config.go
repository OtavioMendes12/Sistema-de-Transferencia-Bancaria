package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI       string
	Database       string
	CollectionName string
	ServerPort     string
	JWTSecret      string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}

	return Config{
		MongoURI:       viper.GetString("MONGO_URI"),
		Database:       viper.GetString("DATABASE_NAME"),
		CollectionName: viper.GetString("USERS_COLLECTION"),
		ServerPort:     viper.GetString("SERVER_PORT"),
		JWTSecret:      viper.GetString("JWT_SECRET"),
	}
}
