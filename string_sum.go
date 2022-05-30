package string_sum

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// test
func StringSum(input string) (output string, err error) {
	var errorEmptyInputMy = fmt.Errorf("Error_input_empty: %w", errorEmptyInput)
	var errorPresentLetterMy = fmt.Errorf("Error: letter present ")
	var errorInputMy = fmt.Errorf("Error_incorectly_input: %w", errorNotTwoOperands)

	input = strings.ReplaceAll(input, " ", "") 

	if len(strings.Fields(input)) == 0 {
		return "", errorEmptyInputMy
	}

	isLatter, _ := regexp.MatchString("[A-Za-z]", input)
	if isLatter != false {
		return "", errorPresentLetterMy
	}

	isCorrect, _ := regexp.MatchString("^(\\-|\\+)?\\d+(\\-|\\+)\\d+$", input)
	if isCorrect != true {
		return "", errorInputMy
	}

	input = strings.ReplaceAll(input, "-", " -")
	input = strings.ReplaceAll(input, "+", " +")
	input = strings.TrimSpace(input)
	res := strings.Split(input, " ")
	x, _ := strconv.Atoi(res[0])
	y, _ := strconv.Atoi(res[1])
	output = strconv.Itoa(x + y)


	return output, nil
}
