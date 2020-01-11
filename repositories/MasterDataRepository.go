package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

type MasterDataRepositoryWithCircuitBreaker struct {
	MasterDataRepository interfaces.IMasterDataRepository
}

func (repository *MasterDataRepositoryWithCircuitBreaker) GetCustomerData(query string, value string, collection string) (models.CustomerModel, error) {
	output := make(chan models.CustomerModel, 1)
	hystrix.ConfigureCommand("get_user_by_username_and_password", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_by_username_and_password", func() error {
		user, _ := repository.MasterDataRepository.GetCustomerData(query, value, collection)
		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.CustomerModel{}, err
	}
}

type MasterDataRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *MasterDataRepository) GetCustomerData(query string, value string, collection string) (models.CustomerModel, error) {
	filter := bson.M{"userName": query, "password": value}
	row, err := repository.FindOne(filter, "users", "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return models.CustomerModel{}, nil
	}

	var user models.CustomerModel
	row.DecodeResults(&user)
	return user, nil
}
