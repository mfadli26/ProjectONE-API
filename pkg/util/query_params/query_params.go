package query_params

import (
	"encoding/json"
	"fmt"
	"projectONE/internal/abstraction"
	"projectONE/pkg/util/general"
	"strconv"
	"strings"
	"time"
)

type QueryParam struct {
	Key              string
	Value            string
	Type             string
	IsDate           bool
	ChangeFormatDate bool
	ToFormatDate     string
	FromFormatDate   string
}

func CheckAndProcessQueryParams(ctx *abstraction.Context, orderParams map[string]string) string {
	if ctx.QueryParam("page") != "" {
		ctx.QueryParams().Add("page", "1")
	}
	if ctx.QueryParam("pageSize") != "" {
		ctx.QueryParams().Add("page", "10")
	}
	// if ctx.QueryParam("orderName") != "" {
	// 	ctx.QueryParams().Add("orderName", "created_date")
	// }
	var orderType string
	if ctx.QueryParam("orderType") == "descend" || ctx.QueryParam("orderType") == "DESC" {
		ctx.QueryParams().Set("orderType", "DESC")
		orderType = "DESC"
	} else {
		ctx.QueryParams().Set("orderType", "ASC")
		orderType = "ASC"
	}
	offset := (general.GetInt(ctx.QueryParam("page")) - 1) * general.GetInt(ctx.QueryParam("pageSize"))
	ctx.QueryParams().Add("offset", strconv.Itoa(offset))

	orderName := ctx.QueryParam("orderName")

	if value, ok := orderParams[orderName]; ok {
		order := value + " " + orderType
		return order
	}

	return ""
}

func StringToArraysQueryParams(ctx *abstraction.Context, name string) interface{} {
	if ctx.QueryParam(name) != "" {
		dataString := ctx.QueryParam(name)
		dataMap := []map[string]string{}
		err := json.Unmarshal([]byte(dataString), &dataMap)
		if err == nil {
			data := []string{}
			for _, value := range dataMap {
				data = append(data, value["value"])
			}
			return data
		}
		return nil
	}
	return nil
}

func CheckAndProcessQueryParams2(ctx *abstraction.Context, queryParam []QueryParam) string {
	if ctx.QueryParam("page") != "" {
		ctx.QueryParams().Add("page", "1")
	}
	if ctx.QueryParam("pageSize") != "" {
		ctx.QueryParams().Add("page", "10")
	}
	// if ctx.QueryParam("orderName") != "" {
	// 	ctx.QueryParams().Add("orderName", "created_date")
	// }

	var orderType string
	if ctx.QueryParam("orderType") == "descend" || ctx.QueryParam("orderType") == "DESC" {
		ctx.QueryParams().Set("orderType", "DESC")
		orderType = "DESC"
	} else {
		ctx.QueryParams().Set("orderType", "ASC")
		orderType = "ASC"
	}

	offset := (general.GetInt(ctx.QueryParam("page")) - 1) * general.GetInt(ctx.QueryParam("pageSize"))
	ctx.QueryParams().Add("offset", strconv.Itoa(offset))

	orderName := ctx.QueryParam("orderName")

	for _, row := range queryParam {
		if row.Key == orderName {
			value := strings.Split(row.Value, "::")
			order := value[0] + " " + orderType
			return order
		}
	}

	return ""
}

func GetWhereParams(ctx *abstraction.Context, queryParams []QueryParam, isNew bool) string {
	var where string = ""
	for _, row := range queryParams {
		if ctx.QueryParam(row.Key) != "" {
			where += fmt.Sprintf(" AND %s %s @%s", row.Value, row.Type, row.Key)
		}
	}
	if isNew {
		where = where[4:]
	}
	return where
}

func ProcessQueryParams(ctx *abstraction.Context, queryParam []QueryParam) (map[string]interface{}, int, int) {
	params := make(map[string]interface{})
	pageSize := general.GetInt(ctx.QueryParam("pageSize"))
	offset := general.GetInt(ctx.QueryParam("offset"))
	for _, row := range queryParam {
		if row.Type == "ILIKE" {
			params[row.Key] = "%" + ctx.QueryParam(row.Key) + "%"
		} else if row.Type == "IN" {
			value := StringToArraysQueryParams(ctx, row.Key)
			params[row.Key] = value
		} else if row.ChangeFormatDate == true {
			dt, _ := time.Parse(row.FromFormatDate, ctx.QueryParam(row.Key))
			dateTime := dt.Format(row.ToFormatDate)
			params[row.Key] = dateTime
		} else {
			params[row.Key] = ctx.QueryParam(row.Key)
		}
	}
	return params, pageSize, offset
}

func CpfQueryParamsProcess(ctx *abstraction.Context, where string) (whereQuery string, whereParams map[string]interface{}, order string, pageSize int, offset int) {
	orderParams := []QueryParam{}
	orderParams = append(orderParams, QueryParam{
		Key:   "cpf_number",
		Value: "cpf_number",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "tipekontrak_name",
		Value: "tipekontrak_name",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "cpf_namaproject",
		Value: "cpf_namaproject",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "users_name",
		Value: "users_name",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "divisi_name",
		Value: "divisi_name",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:              "created_date",
		Value:            "created_date::date",
		Type:             "=",
		IsDate:           true,
		ChangeFormatDate: true,
		ToFormatDate:     "2006-01-02",
		FromFormatDate:   "02-01-2006",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "pic_users_name",
		Value: "pic_users_name",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "cpf_namarekanan",
		Value: "cpf_namarekanan",
		Type:  "ILIKE",
	})
	orderParams = append(orderParams, QueryParam{
		Key:              "date_start",
		Value:            "created_date::date",
		Type:             ">=",
		IsDate:           true,
		ChangeFormatDate: true,
		ToFormatDate:     "2006-01-02",
		FromFormatDate:   "02/01/2006",
	})
	orderParams = append(orderParams, QueryParam{
		Key:              "date_end",
		Value:            "created_date::date",
		Type:             "<=",
		IsDate:           true,
		ChangeFormatDate: true,
		ToFormatDate:     "2006-01-02",
		FromFormatDate:   "02/01/2006",
	})
	orderParams = append(orderParams, QueryParam{
		Key:              "approval_date",
		Value:            "approval_date::date",
		Type:             "=",
		IsDate:           true,
		ChangeFormatDate: true,
		ToFormatDate:     "2006-01-02",
		FromFormatDate:   "02-01-2006",
	})
	orderParams = append(orderParams, QueryParam{
		Key:   "statuscpf_name",
		Value: "statuscpf_name",
		Type:  "IN",
	})

	order = CheckAndProcessQueryParams2(ctx, orderParams)

	params, pageSize, offset := ProcessQueryParams(ctx, orderParams)

	isNew := false
	if where == "" {
		isNew = true
	}

	whereQuery = GetWhereParams(ctx, orderParams, isNew)
	where = where + whereQuery

	whereParams = map[string]interface{}{
		"date_start":            params["date_start"],
		"date_end":              params["date_end"],
		"created_date":          params["created_date"],
		"cpf_number":            params["cpf_number"],
		"approval_date":         params["approval_date"],
		"tipekontrak_name":      params["tipekontrak_name"],
		"cpf_namaproject":       params["cpf_namaproject"],
		"users_name":            params["users_name"],
		"divisi_name":           params["divisi_name"],
		"pic_users_name":        params["pic_users_name"],
		"cpf_namarekanan":       params["cpf_namarekanan"],
		"statuscpf_name":        params["statuscpf_name"],
		"reject_head_divisi":    "f28d07e3-e972-4ec1-8c4a-038849e2d8c9",
		"reject_head_divisi_ho": "342346f1-f311-4c08-80f1-d52eae1e50ba",
		"reject_direktur":       "27f5a0fe-ae8e-4b67-b1da-9729d796785c",
		"reject_losd":           "4bb449e8-3fa8-4785-9b2e-7c75d8f3e10c",
		"false":                 false,
		"true":                  true,
	}
	return where, whereParams, order, pageSize, offset
}
