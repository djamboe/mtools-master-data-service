package interfaces

import "github.com/djamboe/mtools-master-data-service/models"

type ILoginService interface {
	DoLogin(username string, password string) (models.UserModel, error)
}
