package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("กรอกเลขตัวที่ 1: ")
	fmt.Scan(&num1)

	fmt.Print("กรอกเครื่องหมาย (+ - * /): ")
	fmt.Scan(&operator)

	fmt.Print("กรอกเลขตัวที่ 2: ")
	fmt.Scan(&num2)

	if (operator =="+") {
		fmt.Println("ผลลัพธ์ =", num1+num2)
	}else if (operator == "-") {
		fmt.Println("ผลลัพธ์ =", num1-num2)
	}else if (operator == "*") {
		fmt.Println("ผลลัพธ์ =", num1*num2)
	}else if (operator == "/") {
		if num2 == 0 {
			fmt.Println("หาร 0 ไม่ได้")
		} else {
			fmt.Println("ผลลัพธ์ =", num1/num2)
		}
	}else {
		fmt.Println("เครื่องหมายไม่ถูกต้อง")
	}
}