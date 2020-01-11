package controllers

import (
	"github.com/djamboe/mtools-master-data-service/interfaces"
	"github.com/djamboe/mtools-master-data-service/models"
)

type MasterDataController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(username string, password string) (models.UserModel, error) {
	login, err := controller.DoLogin(username, password)
	if err != nil {
		panic(err)
	}
	return login, nil
}
