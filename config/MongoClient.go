package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetMongoClient() *mongo.Database {
	database := getEnv("MONGO_DB_NAME", "members-gin")

	return GetMongoClientByDatabaseName(database)
}

func GetMongoClientByDatabaseName(database string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	host := getEnv("MONGODB_HOST", "localhost")
	port := getEnv("MONGODB_PORT", "27017")

	uri := "mongodb://" + host + ":" + port

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return nil
	}

	return client.Database(database)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
