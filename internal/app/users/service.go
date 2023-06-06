package users

import (
	"fmt"
	"projectONE/internal/abstraction"
	"projectONE/internal/dto"
	"projectONE/internal/factory"
	"projectONE/internal/repository"
	"projectONE/pkg/util/trxmanager"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type service struct {
	Repository repository.Users
	Db         *gorm.DB
}

var err error

func NewService(f *factory.Factory) *service {
	repository := f.UsersRepository
	db := f.Db

	return &service{
		repository,
		db,
	}
}

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.UsersRegistrationRequest) (*dto.UsersRegistrationResponse, error)
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.UsersRegistrationRequest) (*dto.UsersRegistrationResponse, error) {
	var result dto.UsersRegistrationResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		fmt.Println(*payload)
		fmt.Println("lalalalalala")
		return nil
	}); err != nil {
		logrus.Error("Error in Create service:", err) // Tambahkan log error di sini
		return nil, err
	}

	result = dto.UsersRegistrationResponse{
		Message: "success create user!",
	}

	return &result, nil
}
