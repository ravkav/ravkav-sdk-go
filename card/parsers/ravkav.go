package parsers

import (
	"go-ravkav/card/parsers/dictionaries"
	"strconv"
)

func Operator(hexString string) string {
	operatorID := ParseEn1545Number(hexString)
	operatorIDInt, err := strconv.Atoi(operatorID)
	if err != nil {
		return ""
	}
	if operator, ok := dictionaries.RavkavIssuers[operatorIDInt]; ok {
		return operator
	}
	return operatorID
}
