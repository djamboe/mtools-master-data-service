package controllers

import (
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
)

type MasterDataController struct {
	interfaces.IMasterDataService
}

func (controller *MasterDataController) GetCustomerDataProcess(query string, value string, collection string) ([]*models.CustomerModel, error) {
	customerData, err := controller.GetCustomerData(query, value, collection)
	if err != nil {
		panic(err)
	}
	return customerData, nil
}
