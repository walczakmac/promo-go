package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListProducts interface {
	Handle() ([]map[string]interface{}, error)
}

type listProducts struct {
	db *mongo.Database
}

func NewListProducts(db *mongo.Database) ListProducts {
	return listProducts{db: db}
}

func (handler listProducts) Handle() ([]map[string]interface{}, error) {
	result, err := handler.db.Collection("products").Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}

	var list []map[string]interface{}

	for result.Next(nil) {
		product := make(map[string]interface{})
		_ = result.Decode(&product)
		list = append(list, product)
	}

	return list, nil
}
