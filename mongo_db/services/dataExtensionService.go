package services

import (
	"context"
	"log"

	models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoService struct {
	client *mongo.Client
	dbName string
}

func NewMongoService(client *mongo.Client, dbName string) *MongoService {
	return &MongoService{
		client: client,
		dbName: dbName,
	}
}

func (s *MongoService) StoreDataExtension(ctx context.Context, data *models.DataExtensionMongoDB) (*mongo.InsertOneResult, error) {
	collection := s.client.Database(s.dbName).Collection("DataExtensions")
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Printf("Failed to store data extension: %v", err)
		return nil, err
	}
	return result, nil
}
