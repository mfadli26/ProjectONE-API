package response

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"projectONE/internal/config"
	"projectONE/pkg/log"
	"projectONE/pkg/util/date"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

type errorResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	WrongPassword       Error
	Duplicate           Error
	NotFound            Error
	RouteNotFound       Error
	UnprocessableEntity Error
	Unauthorized        Error
	BadRequest          Error
	Validation          Error
	InternalServerError Error
}

var ErrorConstant errorConstant = errorConstant{
	WrongPassword: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Password not match",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Created value already exists",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Unauthorized, please login",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Bad Request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	Validation: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Something bad happened",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res *Error, message error) *Error {
	res.ErrorMessage = message
	return res
}

func CustomErrorBuilder(code int, err string, message string, Info string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
				Info:    &Info,
			},
			Error: err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(c echo.Context) error {
	var errorMessage string
	id := shortid.MustGenerate()

	if e.ErrorMessage != nil {
		e.Response.Meta.ErrorId = &id
		errorMessage = fmt.Sprintf("%+v", errors.WithStack(e.ErrorMessage))
	}

	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Error(e.ErrorMessage)

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		logrus.Warn("error read body, message : ", e.Error())
	}

	bHeader, err := json.Marshal(c.Request().Header)
	if err != nil {
		logrus.Warn("error read header, message : ", e.Error())
	}

	go func() {
		retries := 3
		logError := log.LogError{
			ID:           id,
			Header:       string(bHeader),
			Body:         string(body),
			URL:          c.Request().URL.Path,
			HttpMethod:   c.Request().Method,
			ErrorMessage: errorMessage,
			Level:        "Error",
			AppName:      config.Get().Server.App,
			Version:      config.Get().Server.Version,
			Env:          os.Getenv("ENV"),
			CreatedAt:    *date.DateTodayLocal(),
		}
		for i := 0; i < retries; i++ {
			err := log.InsertErrorLog(context.Background(), &logError)
			var logrusErr error
			if e.ErrorMessage != nil {
				e.Response.Meta.ErrorId = &id
				logrusErr = log.LogruswriteError(id, e.ErrorMessage.Error())
			}
			if err == nil && logrusErr == nil {
				break
			}
		}
	}()

	return c.JSON(e.Code, e.Response)
}
