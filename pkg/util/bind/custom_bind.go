package bind

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

func Bind(payload interface{}, c echo.Context, field string) error {
	db := new(echo.DefaultBinder)

	if err := db.Bind(payload, c); err != nil {
		return err
	}

	if err := BindStringToStruct(payload, c, field); err != nil {
		return err
	}

	return nil
}

func BindStringToStruct(payload interface{}, c echo.Context, field string) error {
	typ, check := reflect.TypeOf(payload).Elem().FieldByName(field)
	if !check {
		return errors.New("Field Not Found!")
	}

	val := reflect.ValueOf(payload).Elem().FieldByName(field)

	refValue := reflect.New(typ.Type).Elem()
	var arrInterface []map[string]string
	json.Unmarshal([]byte(c.FormValue(strings.ToLower(field))), &arrInterface)

	for _, row := range arrInterface {
		elemValue := reflect.New(typ.Type.Elem()).Elem()
		elemType := elemValue.Type()
		for index2 := 0; index2 < elemType.NumField(); index2++ {
			typeField := elemType.Field(index2)
			elemValue.Field(index2).SetString(row[typeField.Tag.Get("form")])
		}

		refValue = reflect.Append(refValue, elemValue)
	}

	val.Set(refValue)

	return nil
}
