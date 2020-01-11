package v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/djamboe/mtools-master-data-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type masterDataServiceServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewMasterDataServiceServer() v1.MasterDataServiceServer {
	return &masterDataServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *masterDataServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	} else {
		return status.Errorf(codes.Unimplemented,
			"please input your api version")
	}
	return nil
}

func (s *masterDataServiceServer) GetCustomerData(ctx context.Context, req *v1.GetCustomerDataRequest) (*v1.GetCustomerDataResponse, error) {
	message := "Successfully get customer data"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	masterDataController := ServiceContainer().InjectMasterDataController()
	response, err := masterDataController.GetCustomerData(req.Query, req.Value, req.Collection)

	if err != nil {
		message = "Failed retreive customer data"
		errorStatus = true
	}

	customerDataSlice := make([]*v1.Customer, len(response))

	for i, value := range response {
		customerDataSlice[i] = new(v1.Customer)
		customerDataSlice[i].XId = value.Id
		customerDataSlice[i].Customername = value.CustomerName
		customerDataSlice[i].Customeraddress = value.CustomerAddress
		customerDataSlice[i].Customeremail = value.CustomerEmail
		customerDataSlice[i].Customerphone = value.CustomerPhone
		customerDataSlice[i].Status = value.Status
	}

	return &v1.GetCustomerDataResponse{
		Api:      apiVersion,
		Message:  message,
		Error:    errorStatus,
		Customer: customerDataSlice,
	}, nil
}

func (s *masterDataServiceServer) GetProductData(ctx context.Context, req *v1.GetProductDataRequest) (*v1.GetProductDataResponse, error) {
	message := "Successfully get product data"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	masterDataController := ServiceContainer().InjectMasterDataController()
	response, err := masterDataController.GetProductData(req.Query, req.Value, req.Collection)

	if err != nil {
		message = "Failed retreive product data"
		errorStatus = true
	}

	productDataSlice := make([]*v1.Product, len(response))

	for i, value := range response {
		productDataSlice[i] = new(v1.Product)
		productDataSlice[i].XId = value.Id
		productDataSlice[i].Productname = value.ProductName
		productDataSlice[i].Description = value.Description
		productDataSlice[i].Status = value.Status
	}

	return &v1.GetProductDataResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
		Product: productDataSlice,
	}, nil
}
