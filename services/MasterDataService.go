package services

import (
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
)

type MasterDataService struct {
	interfaces.IMasterDataRepository
}

func (service *MasterDataService) getCustomerData(query string, value string, collection string) (models.CustomerModel, error) {
	var customer models.CustomerModel
	customer, err := service.GetCustomerData(query, value, collection)
	if err != nil {
		panic(err)
	}
	return customer, nil
}
