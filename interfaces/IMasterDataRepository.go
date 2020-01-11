package interfaces

import (
	"github.com/djamboe/mtools-master-data-service/models"
)

type IMasterDataRepository interface {
	GetCustomerData(query string, value string, collection string) ([]*models.CustomerModel, error)
	GetProductData(query string, value string, collection string) ([]*models.ProductModel, error)
}
