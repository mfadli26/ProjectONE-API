package factory

import (
	"projectONE/database"
	"projectONE/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db *gorm.DB

	repository_initiated
}

type repository_initiated struct {
	UsersRepository repository.Users
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()
	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("POSTGRES")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		f.SetupDb()
	}
	f.UsersRepository = *repository.NewUsers(f.Db)
}
