package repository

import (
	"context"
	"geo.org/gin-members/config"
	"geo.org/gin-members/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "members"

func FindAllMembers() (members []models.Member) {
	cursor, err := config.GetMongoClient().Collection(collectionName).Find(context.TODO(), bson.D{})

	if err != nil {
		return []models.Member{}
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error while closing cursor %v", err)
		}
	}(cursor, context.TODO())

	for cursor.Next(context.TODO()) {
		var member models.Member
		err := cursor.Decode(&member)

		if err != nil {
			log.Fatalf("Error while decoding %v", err)
		}
		members = append(members, member)
	}

	return members
}

func InsertMember(member models.Member) (newMember models.Member) {
	result, err := config.GetMongoClient().Collection(collectionName).InsertOne(context.TODO(), member)

	if err != nil {
		log.Printf("Error while inserting member %v", err)
		return newMember
	}

	member.ID = result.InsertedID.(primitive.ObjectID)

	return member
}
