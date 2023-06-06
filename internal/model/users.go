package model

import "projectONE/internal/abstraction"

type UsersModel struct {
	UsersId       int    `json:"users_id" form:"users_id"`
	UsersName     string `json:"users_name" form:"users_name"`
	UsersEmail    string `json:"users_email" form:"users_email"`
	UsersPassword string `json:"users_password" form:"users_password"`

	Context *abstraction.Context `json:"-" gorm:"-"`
}
