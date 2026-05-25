package main

import (
	"fmt"
)

func nmber(prompt string) float64 {
	var num float64

	fmt.Print(prompt)
	fmt.Scan(&num)
	return num
}

func operator() string {
	var operator string

	fmt.Print("กรอกเครื่องหมาย (+ - * /): ")
	fmt.Scan(&operator)
	return operator
}

func calculate(num1 float64, operator string, num2 float64) (float64, string) {
	if operator == "+" {
		return num1 + num2, ""
	} else if operator == "-" {
		return num1 - num2, ""
	} else if operator == "*" {
		return num1 * num2, ""
	} else if operator == "/" {
		if num2 == 0 {
			return 0, "หาร 0 ไม่ได้"
		}

		return num1 / num2, ""

	} else {
		return 0, "เครื่องหมายไม่ถูกต้อง"
	}
}

func display(result float64, err string) {
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("ผลลัพธ์ =", result)
	}
}

func main() {
	num1 := nmber("กรอกเลขตัวที่ 1: ")
	operator := operator()
	num2 := nmber("กรอกเลขตัวที่ 2: ")

	result, err := calculate(num1, operator, num2)
	display(result, err)
}