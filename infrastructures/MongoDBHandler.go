package infrastructures

import (
	"context"
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandler struct {
	Conn *mongo.Client
}

func (handler *MongoDBHandler) FindOne(filter bson.M, collectionName string, dbName string) (interfaces.IRowMongoDB, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows := collection.FindOne(context.TODO(), filter)
	row := new(MongoRow)
	row.Rows = rows
	return row, nil
}

type MongoRow struct {
	Rows *mongo.SingleResult
}

func (handler *MongoDBHandler) Find(filter bson.M, collectionName string, dbName string) (*mongo.Cursor, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows, _ := collection.Find(context.TODO(), filter)
	row := new(MongoCursor)
	row.Rows = rows
	return rows, nil
}

type MongoCursor struct {
	Rows *mongo.Cursor
}

func (r MongoCursor) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}

func (r MongoCursor) Next() *mongo.Cursor {
	return r.Next()
}

func (r *MongoRow) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}

func (r MongoRow) Next() bool {
	return r.Next()
}
