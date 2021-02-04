package validator

import (
	"fmt"
	"promo/domain"
	"promo/infrastructure/promoProductApi"
	"regexp"
)

type CreateProductValidator interface {
	Validate(product domain.Product)
	GetErrors() map[string]string
}

type createProduct struct {
	errors          map[string]string
	client promoProductApi.Client
}

func NewCreateProductValidator(client promoProductApi.Client) CreateProductValidator {
	return &createProduct{
		errors: make(map[string]string),
		client: client,
	}
}

func (validator *createProduct) Validate(product domain.Product) {
	if product.ID == "" {
		validator.errors["id"] = "Product ID must be provided."
	}
	if product.ID != "" {
		_, err := validator.client.GetProductById(product.ID)
		if nil != err {
			validator.errors["id"] = fmt.Sprintf("Promo product with ID %v not found.", product.ID)
		}
	}
	if product.Name == "" {
		validator.errors["name"] = "Product name must be provided."
	}
	if len(product.Name) > 255 {
		validator.errors["name"] = "Product name can be only up to 255 characters long."
	}
	if product.Price == "" {
		validator.errors["price"] = "Price must be provided."
	}
	r, _ := regexp.Compile("([0-9]+) ([A-Z]{3})")
	if product.Price != "" && false == r.MatchString(product.Price) {
		validator.errors["price"] = fmt.Sprintf("Price must match following regex: %v", r.String())
	}
}

func (validator createProduct) GetErrors() map[string]string {
	return validator.errors
}