package main

import (
	"fmt"
)

func main() {
	var a interface{} = 42             // int
	var b interface{} = "artur"        // string
	var c interface{} = true           // bool
	var d interface{} = make(chan int) // channel

	// Определяем типы переменных
	identifyType(a)
	identifyType(b)
	identifyType(c)
	identifyType(d)
}

// Функция для определения типа переменной
func identifyType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("Тип переменной: int, значение: %d\n", v)
	case string:
		fmt.Printf("Тип переменной: string, значение: %s\n", v)
	case bool:
		fmt.Printf("Тип переменной: bool, значение: %t\n", v)
	case chan int:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}
