package cli

import (
	"fmt"

	"github.com/kalinskilk/arquitetura-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enable", res.GetName())

	case "disable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s, Name: %s, Price: %f, status:%s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
