package main

import (
	"strconv"
	"strings"
)

// fieldParser decodes cron expression
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

// parseField expands cron field expressions (*, */x, x-y, x,y,...)
//func parseField(field string, fieldRange []int) []int {
//	var result []int
//
//	if field == "*" {
//		for i := fieldRange[0]; i <= fieldRange[1]; i++ {
//			result = append(result, i)
//		}
//	} else if strings.Contains(field, "/") {
//		parts := strings.Split(field, "/")
//		step, _ := strconv.Atoi(parts[1])
//		for i := fieldRange[0]; i <= fieldRange[1]; i += step {
//			result = append(result, i)
//		}
//	} else if strings.Contains(field, "-") {
//		parts := strings.Split(field, "-")
//		start, _ := strconv.Atoi(parts[0])
//		end, _ := strconv.Atoi(parts[1])
//		for i := start; i <= end; i++ {
//			result = append(result, i)
//		}
//	} else if strings.Contains(field, ",") {
//		values := strings.Split(field, ",")
//		for _, v := range values {
//			val, _ := strconv.Atoi(v)
//			result = append(result, val)
//		}
//	} else {
//		val, _ := strconv.Atoi(field)
//		result = append(result, val)
//	}
//
//	return result
//}
