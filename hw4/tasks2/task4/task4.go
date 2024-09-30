package main

import "fmt"

func main() {
	fmt.Println(getAgeLimit(18))
}

func getAgeLimit(age int) string {
	if age >= 0 && age <= 8 {
		return "Ребенок"
	} else if age >= 9 && age <= 18 {
		return "Подросток"
	} else if age >= 19 && age <= 65 {
		return "Взрослый"
	} else if age > 65 {
		return "Пожилой"
	}
	return "Ошибка"
}
