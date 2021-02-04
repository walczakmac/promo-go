package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"promo/domain"
	"promo/infrastructure/promoProductApi"
)

type CreateProduct interface {
	Handle(product domain.Product) error
}

type createProduct struct {
	db *mongo.Database
	promoProductApiClient promoProductApi.Client
}

func NewCreateProduct(db *mongo.Database, client promoProductApi.Client) CreateProduct {
	return createProduct{db: db, promoProductApiClient: client}
}

func (handler createProduct) Handle(product domain.Product) error {
	promoProduct, _ := handler.promoProductApiClient.GetProductById(product.ID)
	for _, color := range promoProduct.Colors {
		product.Attributes = append(product.Attributes, domain.Attribute{Name: "color", Value: color})
	}
	for _, size := range promoProduct.Sizes {
		product.Attributes = append(product.Attributes, domain.Attribute{Name: "size", Value: size})
	}

	_, err := handler.db.Collection("products").InsertOne(context.TODO(), product)
	return err
}
