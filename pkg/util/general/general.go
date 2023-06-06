package general

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"mime"
	"os"
	"regexp"

	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func ArraysIntContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetString(name int) string {
	return strconv.Itoa(name)
}

func GetInt(name string) int {
	i, err := strconv.Atoi(name)
	if err != nil {
		return 0
	}
	return i
}

func GetFloat(name string) float64 {
	i, err := strconv.ParseFloat(name, 64)
	if nil != err {
		return 0
	}
	return i
}

func MakeFolder(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.Mkdir(name, os.ModeDir)
		if err != nil {
			logrus.Println("Directory creation failed with error: " + err.Error())
		}
	}
}

func StringInSlice(text string, data []string) bool {
	for _, row := range data {
		if row == text {
			return true
		}
	}
	return false
}

func ArrayInterfaceToArrayString(data []interface{}) []string {
	arrayString := make([]string, len(data))
	for index, row := range data {
		arrayString[index] = row.(string)
	}
	return arrayString
}

func StringInSliceInterface(text string, data []interface{}) bool {
	for _, row := range data {
		if text == row.(string) {
			return true
		}
	}
	return false
}

func StringInSliceMapInterface(text string, data []map[string]interface{}, key string) bool {
	for _, row := range data {
		if text == row[key].(string) {
			return true
		}
	}
	return false
}

func StringInSliceMapInterfaceWithIndex(text string, data []map[string]interface{}, key string) (bool, int) {
	for index, row := range data {
		if text == row[key].(string) {
			return true, index
		}
	}
	return false, -1
}

func RandSeq(n int) string {
	var letters = []rune("123456789abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandSeqCapital(n int) string {
	var letters = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetStringInBetween(text string, start string, end string) (result string) {
	return strings.TrimLeft(strings.TrimRight(text, end), start)
}

func GetExtensionByDataMimeTypes(DataMimeType string) (result string) {
	mimeTypes := strings.Split(DataMimeType, "data:")[1]
	extensionsData, err := mime.ExtensionsByType(mimeTypes)
	if err != nil {
		return err.Error()
	}
	return extensionsData[0]
}

func PadNumberWithZero(value int, number string) string {
	return fmt.Sprintf("%0"+number+"d", value)
}

func formatCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase, minLowerCase int) string {
	var password strings.Builder
	var lowerCharSet string = "abcdedfghijklmnopqrstuvwxyz"
	var upperCharSet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var specialCharSet string = "!@#$%&*"
	var numberSet string = "0123456789"
	var allCharSet string = lowerCharSet + upperCharSet + specialCharSet + numberSet

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase - minLowerCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func SplitWithError(s, sep string) ([]string, error) {
	result := strings.Split(s, sep)

	if len(result) == 1 {
		return nil, errors.New("delimiter not found")
	}

	return result, nil
}

func GetStringInSliceBySubString(text string, data []string) (bool, string) {
	for _, row := range data {
		if strings.Contains(row, text) {
			return true, row
		}
	}
	return false, ""
}

func StructToJsonByte(payload interface{}) ([]byte, error) {
	var queryParams map[string]interface{}
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytePayload, &queryParams)
	for key, value := range queryParams {
		switch value.(type) {
		case string:
			queryParams[key] = "%" + value.(string) + "%"
		default:
			queryParams[key] = value
		}
	}

	jsonString, err := json.Marshal(queryParams)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func GetAgreementValue(aggreementSymbol string, cpfNilaiperjanjian float64) string {
	printer := message.NewPrinter(language.Indonesian)
	nilaiPerjanjian := strconv.FormatFloat(cpfNilaiperjanjian, 'f', -1, 64)
	aggreementValue := printer.Sprintf("%s %s", aggreementSymbol, nilaiPerjanjian)
	return aggreementValue
}
