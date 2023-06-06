package currency

import (
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FloatToCurrency(value float64, symbol string, rounding int, formatter language.Tag) string {
	nilaiPerjanjian := strconv.FormatFloat(value, 'f', rounding, 64)
	printer := message.NewPrinter(formatter)
	aggreementValue := printer.Sprintf("%s %f", symbol, value)
	aggreementValueSplit := strings.Split(aggreementValue, ",")
	if len(aggreementValueSplit) > 1 {
		indexPrecission := strings.Index(nilaiPerjanjian, ".")
		decimalValue := ""
		if indexPrecission != -1 {
			decimalValue = "," + nilaiPerjanjian[indexPrecission+1:]
		}
		aggreementValue = aggreementValueSplit[0] + decimalValue
	}
	return aggreementValue
}
