package main

import (
	"strconv"
	"strings"
)

// FieldParser decodes cron expression
func fieldParser(value string, fieldRange []int) []int {
	var parsedField []int

	if value == "*" {
		for i := fieldRange[0]; i <= fieldRange[1]; i++ {
			parsedField = append(parsedField, i)
		}
	} else if strings.Contains(value, "/") {
		strArr := strings.Split(value, "/")
		stepSize, err := strconv.Atoi(strArr[1])
		if err != nil {

		}
		for i := fieldRange[0]; i <= fieldRange[1]; i += stepSize {
			parsedField = append(parsedField, i)
		}
	} else if strings.Contains(value, "-") {
		strArr := strings.Split(value, "-")
		initVal, err := strconv.Atoi(strArr[0])
		if err != nil {
			//return error
		}
		finVal, err := strconv.Atoi(strArr[1])
		if err != nil {
		}
		if (initVal < fieldRange[0] || initVal > fieldRange[1]) || (finVal < fieldRange[0] || finVal > fieldRange[1]) {
			//return error
		}
		for i := initVal; i <= finVal; i++ {
			parsedField = append(parsedField, i)
		}
	} else if strings.Contains(value, ",") {
		strArr := strings.Split(value, ",")
		for _, str := range strArr {
			val, err := strconv.Atoi(str)
			if err != nil {

			}
			parsedField = append(parsedField, val)
		}
	} else {
		val, _ := strconv.Atoi(value)
		parsedField = append(parsedField, val)
	}

	return parsedField
}
