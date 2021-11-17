package main

import (
	"fmt"
)

//Функция возвращает мапу из кодов букв и частоты их появления в сообщении
func calculateFrequancy(message string) map[rune]int {

	characterFrequency := make(map[rune]int)

	for _, char := range message {
		if _, ok := characterFrequency[char]; ok {
			characterFrequency[char] += 1
		} else {
			characterFrequency[char] = 1
		}
	}

	fmt.Println(characterFrequency, " мапа сообщения")
	calculateProbability(characterFrequency, message)

	return characterFrequency
}

//Разбивает мапу на 2 массива и сортирует значения в порядке возрастания
func calculateProbability(characterFrequency map[rune]int, message string) /*([]int, []rune)*/ {
	lengthOfMessage := len(message)
	charMas := make([]rune, 0, lengthOfMessage)
	freqMas := make([]int, 0, lengthOfMessage)

	//Разбивает
	for char, freq := range characterFrequency {

		if len(freqMas) == 0 || freqMas[len(freqMas)-1] < freq {
			freqMas = append(freqMas, freq)
			charMas = append(charMas, char)
			continue
		}

		//fmt.Println("Начало цикла сортировки")
		for i := 0; i < len(charMas); i++ {
			if freqMas[i] >= freq {
				freqMas = append(freqMas[:i], append([]int{freq}, freqMas[i:]...)...)
				charMas = append(charMas[:i], append([]rune{char}, charMas[i:]...)...)
				break
			} /*else {
				freqMas = append(freqMas, freq)
				charMas = append(charMas, char)
				break
			}*/
		}
	}

	fmt.Println(charMas, " Это массив чаров")
	fmt.Println(freqMas, " Это массив частот")

}

func main() {

	var message string
	fmt.Scanln(&message)
	calculateFrequancy(message)

}
