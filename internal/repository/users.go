package repository

import (
	"projectONE/internal/abstraction"
	"projectONE/internal/dto"
	"projectONE/internal/model"

	"gorm.io/gorm"
)

type Users struct {
	abstraction.Repository
}

func NewUsers(db *gorm.DB) *Users {
	return &Users{
		abstraction.Repository{
			Db: db,
		},
	}
}

type UsersRepository interface {
	Create(ctx *abstraction.Context, data *dto.UsersRegistrationRequest) (*model.UsersModel, error)
}

func (r *Users) Create(ctx *abstraction.Context, data *model.UsersModel) (*model.UsersModel, error) {
	conn := r.CheckTrx(ctx)

	data.Context = ctx

	err := conn.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
