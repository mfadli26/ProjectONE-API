package dto

type UsersRegistrationRequest struct {
	UsersName     string `json:"users_name"`
	UsersEmail    string `json:"users_email"`
	UsersPassword string `json:"users_password"`
}

type UsersRegistrationResponse struct {
	Message string `json:"message"`
}
