package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type calc struct {
	value1 float64
	value2 float64
}

func (calc calc) operation(value1 int, value2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return value1 + value2, nil
	case "-":
		return value1 - value2, nil
	case "*":
		return value1 * value2, nil
	case "/":
		return value1 / value2, nil
	default:
		return 0, errors.New("Invalid operator")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("An error occurred ")
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func parser(value string) int {
	castValue, error := strconv.Atoi(value)
	checkError(error)
	return castValue
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// func main() {

// 	values := readInput()
// 	regex := regexp.MustCompile(`(\d{1,})([+-/*]{1})(\d{1,})`)

// 	if regex.Match([]byte(values)) {
// 		value1 := regex.ReplaceAllString(values, "$1")
// 		operator := regex.ReplaceAllString(values, "$2")
// 		value2 := regex.ReplaceAllString(values, "$3")
// 		num1 := parser(value1)
// 		num2 := parser(value2)

// 		calc := calc{}
// 		response, _ := calc.operation(num1, num2, operator)

// 		println("El resultado: ", response)
// 	}

// }
