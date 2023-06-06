package users

import (
	"fmt"
	"projectONE/internal/abstraction"
	"projectONE/internal/dto"
	"projectONE/internal/factory"
	res "projectONE/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Create(c echo.Context) error {
	cc := abstraction.Context{}

	fmt.Println("1")
	payload := new(dto.UsersRegistrationRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	fmt.Println("2")

	result, err := h.service.Create(&cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
