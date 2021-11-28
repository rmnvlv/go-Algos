package main

import (
	"fmt"
)

type myShenonFano struct {
	characterFrequency    map[rune]int
	newCharacterFrequency map[rune]string
	encodedMessage        string
}

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

		for i := 0; i < len(charMas); i++ {
			if freqMas[i] >= freq {
				freqMas = append(freqMas[:i], append([]int{freq}, freqMas[i:]...)...)
				charMas = append(charMas[:i], append([]rune{char}, charMas[i:]...)...)
				break
			}
		}
	}

	fmt.Println(charMas, " Это массив чаров")
	fmt.Println(freqMas, " Это массив частот")

	//пустая мапа, для чаров - ключи, на месте значений будут 0101010
	if myShenon.newCharacterFrequency == nil {
		myShenon.newCharacterFrequency = make(map[rune]string)
	}
	for _, char := range charMas {
		myShenon.newCharacterFrequency[char] = ""
	}

	n := len(charMas) / 2
	if len(charMas) != n*2 {
		n += 1
	}
	encodingMessage(charMas[:n], myShenon.newCharacterFrequency, true)
	encodingMessage(charMas[n:], myShenon.newCharacterFrequency, false)

	fmt.Println(myShenon.newCharacterFrequency)

	getMessage(myShenon.newCharacterFrequency, message)
}

//Шифрование всех символов сообщения, минимальная длина символа будет у чаще встречающегося в сообщении символа
func encodingMessage(charMas []rune, mapRune map[rune]string, flag bool) /* map[rune]int */ {

	bit := "0"
	lenghChar := len(charMas)

	if flag {
		bit = "1"
	}

	for _, char := range charMas {
		mapRune[char] += bit
	}

	if lenghChar >= 2 {
		t := lenghChar / 2
		if lenghChar != t*2 {
			t += 1
		}
		encodingMessage(charMas[:t], mapRune, true)
		encodingMessage(charMas[t:], mapRune, false)
	}
}

func getMessage(mapRune map[rune]string, message string) {

	for _, char := range message {
		myShenon.encodedMessage += string(mapRune[char])
		myShenon.encodedMessage += " "
	}
	fmt.Println(myShenon.encodedMessage)
}

var myShenon myShenonFano

func main() {

	var message string
	fmt.Scanln(&message)
	calculateFrequancy(message)

}
