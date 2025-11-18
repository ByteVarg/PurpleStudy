package main

// < 16 - сильный дефицит массы тела
// 16-18,5 - дефицит массы тела
// 18,5-25 - норма
// 25-30 - избыточная масса тела (предожирение)
// 30-35 - ожирение первой степени
// 35-40 - ожирение второй степени
// > 40 - ожирение третьей степени (морбидное)

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func getUserInput() (float64, float64) {
	var height, weight float64
	fmt.Print("Укажите свой рост в см: ")
	fmt.Scan(&height)
	fmt.Print("Укажите свой вес в кг: ")
	fmt.Scan(&weight)
	return weight, height
}

func calc(height, weight float64) (float64, error) {
	height = height / 100
	if height <= 0 || weight <= 0 {
		return 0, errors.New("error: wrong user parapetres")
	}
	BMI := weight / math.Pow(height, 2)
	return BMI, nil

}

func diagType(result float64) {

	diagnosis := map[int]string{
		1: "сильный дефицит массы тела",
		2: "дефицит массы тела",
		3: "норма",
		4: "избыточная масса тела (предожирение)",
		5: "ожирение первой степени",
		6: "ожирение второй степени",
		7: "ожирение третьей степени (морбидное)",
	}

	switch {
	case result < 16:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[1])
	case result < 18.5:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[2])
	case result < 25:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[3])
	case result < 30:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[4])
	case result < 35:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[5])
	case result < 40:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[6])
	default:
		fmt.Printf("Ваш индекс массы тела: %.0f\nВаш диагноз: %s\n", result, diagnosis[7])

	}

}

func main() {

	for {
		fmt.Println("\nНачинаем расчет индекса массы тела.")
		var choice string
		weight, height := getUserInput()
		result, err := calc(height, weight)
		if err != nil {
			fmt.Println("\nОшибка: введены некоректные параметры рост/вес.")
			continue
		}
		diagType(result)

		fmt.Print("Хотите продолжить? (y/n): ")
		fmt.Scanln(&choice)
		fmt.Scanln(strings.ToLower(choice))
		if choice != "y" {
			break
		} else {
			continue
		}
	}

}
