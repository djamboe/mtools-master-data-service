package repositories

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type MasterDataRepositoryWithCircuitBreaker struct {
	MasterDataRepository interfaces.IMasterDataRepository
}

func (repository *MasterDataRepositoryWithCircuitBreaker) GetCustomerData(query string, value string, collection string) ([]*models.CustomerModel, error) {
	output := make(chan []*models.CustomerModel, 1)
	hystrix.ConfigureCommand("get_customer_data", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_customer_data", func() error {
		user, _ := repository.MasterDataRepository.GetCustomerData(query, value, collection)
		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []*models.CustomerModel{}, err
	}
}

func (repository *MasterDataRepositoryWithCircuitBreaker) GetProductData(query string, value string, collection string) ([]*models.ProductModel, error) {
	output := make(chan []*models.ProductModel, 1)
	hystrix.ConfigureCommand("get_product_data", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_product_data", func() error {
		product, _ := repository.MasterDataRepository.GetProductData(query, value, collection)
		output <- product
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []*models.ProductModel{}, err
	}
}

type MasterDataRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *MasterDataRepository) GetCustomerData(query string, value string, collection string) ([]*models.CustomerModel, error) {
	filter := bson.M{query: value}
	row, err := repository.Find(filter, collection, "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return []*models.CustomerModel{}, nil
	}

	var listCostumerData []*models.CustomerModel
	for row.Next(context.TODO()) {
		var data models.CustomerModel
		err = row.Decode(&data)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listCostumerData = append(listCostumerData, &data)
	}
	return listCostumerData, nil
}

func (repository *MasterDataRepository) GetProductData(query string, value string, collection string) ([]*models.ProductModel, error) {
	filter := bson.M{query: value}
	row, err := repository.Find(filter, collection, "maroon_martools")

	if err != nil {
		panic(err)
	}

	if row == nil {
		return []*models.ProductModel{}, nil
	}

	var listProductData []*models.ProductModel
	for row.Next(context.TODO()) {
		var data models.ProductModel
		err = row.Decode(&data)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listProductData = append(listProductData, &data)
	}
	return listProductData, nil
}
