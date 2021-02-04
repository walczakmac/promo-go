package controllers

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"promo/application/handler"
	"promo/application/validator"
	"promo/domain"
	"promo/infrastructure/promoProductApi"
)

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type productController struct {
	db *mongo.Database
	promoProductApiClient promoProductApi.Client
}

func NewProductController(db *mongo.Database, client promoProductApi.Client) ProductController {
	return productController{db: db, promoProductApiClient: client}
}

func (controller productController) Create(w http.ResponseWriter, r *http.Request) {
	product := domain.Product{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProductValidator := validator.NewCreateProductValidator(controller.promoProductApiClient)
	createProductValidator.Validate(product)

	w.Header().Add("Content-Type", "application/json")

	if len(createProductValidator.GetErrors()) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(createProductValidator.GetErrors())
	}

	createProductHandler := handler.NewCreateProduct(controller.db, controller.promoProductApiClient)
	err = createProductHandler.Handle(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func (controller productController) Index(w http.ResponseWriter, r *http.Request) {
	listProductsHandler := handler.NewListProducts(controller.db)
	list, err := listProductsHandler.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(list)
}