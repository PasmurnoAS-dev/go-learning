package main

import "fmt"

func main() {
	// Способы объявления переменных в Go

	// 1. Полное объявление с указанием типа
	var name string = "Саша"
	var age int = 22 // Измени возраст на свой :)
	var isProgrammer bool = true

	fmt.Println("1. Полное объявление:")
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Я программист:", isProgrammer)
	fmt.Println() // Просто пустая строка для красоты

	// 2. Короткое объявление (без указания типа) - самый частый способ!
	// Компилятор сам понимает тип по значению
	language := "Go"
	year := 2025

	fmt.Println("2. Короткое объявление:")
	fmt.Println("Я изучаю:", language)
	fmt.Println("Текущий год:", year)
	fmt.Println()

	// 3. Объявление нескольких переменных сразу
	var (
		city string = "Брянск"
		temp int    = 15
	)

	fmt.Println("3. Объявление нескольких переменных:")
	fmt.Println("Город:", city)
	fmt.Println("Температура:", temp)
	fmt.Println()

	// 4. Изменение значения переменной
	fmt.Println("4. Изменение значения:")
	fmt.Println("Бывший возраст:", age, "года.")
	age = 23 // Меняем значение
	fmt.Println("Мне будет через пол года:", age, "года.")
}
