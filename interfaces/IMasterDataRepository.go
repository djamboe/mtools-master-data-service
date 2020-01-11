package interfaces

import (
	"github.com/djamboe/mtools-master-data-service/models"
)

type IMasterDataRepository interface {
	GetCustomerData(username string, password string, collection string) (models.CustomerModel, error)
}
