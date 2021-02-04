package promoProductApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client interface {
	GetProductById(id string) (PromoProductResponse, error)
}

type promoProductApi struct {
	apiUrl string
	token string
}

func NewPromoProductApi() Client {
	return promoProductApi{
		os.Getenv("PROMO_PRODUCT_API_URL"),
		os.Getenv("PROMO_PRODUCT_API_TOKEN"),
	}
}

func (client promoProductApi) GetProductById(id string) (PromoProductResponse, error) {
	productUrl, _ := url.Parse(fmt.Sprintf("%v/product/%v", client.apiUrl, id))
	request := &http.Request{
		Header: map[string][]string{
			"X-token": {client.token},
		},
		URL: productUrl,
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(request)

	if nil != err {
		return PromoProductResponse{}, err
	}

	product := PromoProductResponse{}
	err = json.NewDecoder(response.Body).Decode(&product)
	if err != nil {
		log.Println("Error loading .env file")
		return PromoProductResponse{}, err
	}

	return product, nil
}
