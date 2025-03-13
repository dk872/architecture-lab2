package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PrefixToLisp(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("invalid input: empty expression")
	}

	tokens := strings.Fields(input)
	list := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if isValidOp(&token) {
			if len(list) < 2 {
				return "", fmt.Errorf("Incorrect input: invalid expression")
			}
			str1 := list[len(list)-1]
			str2 := list[len(list)-2]
			str := fmt.Sprintf("(%s %s %s)", token, str1, str2)
			list = list[:len(list)-2]
			list = append(list, str)
		} else {
			if !isNum(token) {
				return "", fmt.Errorf("Incorrect input: invalid character in expression")
			}
			list = append(list, token)
		}

	}

	if len(list) != 1 {
		return "", fmt.Errorf("Incorrect input: invalid expression")
	}
	return list[0], nil
}

func isValidOp(op *string) bool {
	switch *op {
	case "+", "-", "*", "/":
		return true
	case "^":
		*op = "pow"
		return true
	default:
		return false
	}
}

func isNum(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return true
	}
	return false
}
