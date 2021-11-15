package main

import (
	"fmt"
	//"math"
)

//Функция возвращает мапу из кодов букв и частоты их появления в сообщении
func calculateFrequancy(message string) (map[rune]int, int) {

	characterFrequency := make(map[rune]int)
	countWords := 0

	for _, char := range message {
		if _, ok := characterFrequency[char]; ok {
			characterFrequency[char] += 1
			countWords += 1
		} else {
			characterFrequency[char] = 1
		}
	}

	fmt.Println(characterFrequency, " мапа сообщения")
	fmt.Println(countWords, " Всего ьукв алфавит")
	calculateProbability(characterFrequency, message)

	return characterFrequency, countWords
}

//Разбивает мапу на 2 массива и сортирует значения в порядке возрастания
func calculateProbability(characterFrequency map[rune]int, message string) /*([]int, []rune)*/ {

	lengthOfMessage := len(message)

	charMas := make([]rune, 0, lengthOfMessage)
	freqMas := make([]int, 0, lengthOfMessage)

	for char, freq := range characterFrequency {
		charMas = append(charMas, char)
		freqMas = append(freqMas, freq)
	}

	fmt.Println(charMas, " Это массив чаров")
	fmt.Println(freqMas, " Это массив частот")

}

func main() {
	var message string
	fmt.Scanln(&message)
	calculateFrequancy(message)

}
