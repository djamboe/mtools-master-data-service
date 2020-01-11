package services

import (
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
)

type MasterDataService struct {
	interfaces.IMasterDataRepository
}

func (service *MasterDataService) getCustomerData(query string, value string, collection string) ([]*models.CustomerModel, error) {
	customer, err := service.GetCustomerData(query, value, collection)
	if err != nil {
		panic(err)
	}
	return customer, nil
}

func (service *MasterDataService) getProductData(query string, value string, collection string) ([]*models.ProductModel, error) {
	product, err := service.GetProductData(query, value, collection)
	if err != nil {
		panic(err)
	}
	return product, nil
}
