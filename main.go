package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	return Expression(&expression)
}

func Expression(expression *string) (float64, error) {
	result, err := Terms(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '+':
			*expression = (*expression)[1:]
			num, err := Terms(expression)
			if err != nil {
				return 0, err
			}
			result += num
		case '-':
			*expression = (*expression)[1:]
			num, err := Terms(expression)
			if err != nil {
				return 0, err
			}
			result -= num
		default:
			return result, nil
		}
	}
	return result, nil
}

func Terms(expression *string) (float64, error) {
	result, err := Brackets(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '*':
			*expression = (*expression)[1:]
			num, err := Brackets(expression)
			if err != nil {
				return 0, err
			}
			result *= num
		case '/':
			*expression = (*expression)[1:]
			num, err := Terms(expression)
			if err != nil {
				return 0, err
			}
			if num == 0 {
				return 0, fmt.Errorf("деление на 0 запрещено")
			}
			result /= num
		default:
			return result, nil
		}
	}
	return result, nil
}

func Brackets(expression *string) (float64, error) {
	if len(*expression) == 0 {
		return 0, fmt.Errorf("неправильный конец выражения")
	}
	if (*expression)[0] == '(' {
		*expression = (*expression)[1:]
		result, err := Expression(expression)
		if err != nil {
			return 0, err
		}
		if len(*expression) == 0 || (*expression)[0] != ')' {
			return 0, fmt.Errorf("пропущены скобки")
		}
		*expression = (*expression)[1:]
		return result, nil
	}
	return Nums(expression)
}

func Nums(expression *string) (float64, error) {
	i := 0
	for i < len(*expression) && (unicode.IsDigit(rune((*expression)[i])) || (*expression)[i] == '.') {
		i++
	}
	if i == 0 {
		return 0, fmt.Errorf("введите число, найдено %c", (*expression)[0])
	}
	num, err := strconv.ParseFloat((*expression)[:i], 64)
	if err != nil {
		return 0, fmt.Errorf("неверное число: %v", num)
	}
	*expression = (*expression)[i:]
	return num, nil
}

func main() {
	tests := []string{
		"3 + 5",
		"(1 + 2) * 3",
		"10 / 2 + (8 - 3)",
		"5 + 3 * (12 / 4)",
		"10 / (5 - 5)", // Ошибка деления на ноль
		"3 + * 5",      // Ошибка неверного синтаксиса
		"3 + 12 / 2)",
	}
	for _, test := range tests {
		result, err := Calc(test)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
