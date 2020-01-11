package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoDBHandler interface {
	FindOne(filter bson.M, collectionName string, databaseName string) (IRowMongoDB, error)
	Find(filter bson.M, collectionName string, databaseName string) (*mongo.Cursor, error)
}
type IRowMongoDB interface {
	Next() bool
	DecodeResults(v interface{}) error
}
