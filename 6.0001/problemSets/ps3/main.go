package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type WordScore struct {
	Word string
	Hand int
}

// 加载单词
func loadWords(path string) (strSlice []string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	strSlice = strings.Split(string(buffer), " ")

	return
}

// 启动游戏
func initGame() {
	fmt.Println("Loading word list from file...")
	strSlice := loadWords("./words.txt")
	strSliceLen := len(strSlice)

	if strSliceLen == 0 {
		fmt.Println("no word is avaliable...")
		return
	}

	fmt.Printf("%d words loaded.\n", strSliceLen)
	fmt.Println("play_game not yet implemented.")
}

// 是否元音
func isVowels(letter string) bool {
	return baseHelper(letter, "aeiou")
}

func baseHelper(letter string, list string) bool {

	for _, l := range list {
		if letter == string(l) {
			return true
		}
	}
	return false
}

var scrabbleLetterValue = map[string]int{
	"a": 1, "b": 3, "c": 3, "d": 2, "e": 1, "f": 4, "g": 2, "h": 4, "i": 1, "j": 8, "k": 5, "l": 1, "m": 3, "n": 1, "o": 1, "p": 3, "q": 10, "r": 1, "s": 1, "t": 1, "u": 1, "v": 4, "w": 4, "x": 8, "y": 4, "z": 10,
}

// 乘积，英语称作 product
func GetWordScore(wordType WordScore) int {
	var wordLength = len(wordType.Word)
	var first, second int
	var word = strings.ToLower(wordType.Word)

	for _, letter := range word {
		_, ok := scrabbleLetterValue[string(letter)]
		if ok {
			first += scrabbleLetterValue[string(letter)]
		}
	}

	second = 7*wordLength - 3*(wordType.Hand-wordLength)

	if second < 1 {
		second = 1
	}
	return first * second

}

func GetFrequencyDict(word string) map[string]int {
	var dict = make(map[string]int)
	word = strings.ToLower(word)

	for _, letter := range word {
		_, ok := dict[string(letter)] /*如果确定是真实的,则存在,否则不存在 */
		if !ok {
			dict[string(letter)] = 1
		} else {
			dict[string(letter)]++
		}
	}

	return dict
}

func DisplayHand(dict map[string]int) (hand string) {
	var tempSlice = make([]string, 0)
	for k, v := range dict {
		for ; v > 0; v-- {
			tempSlice = append(tempSlice, k)
		}
	}

	// sort.Strings(tempSlice) 无需排序

	hand = strings.Join(tempSlice, " ")
	return
}

func UpdateHand(dict map[string]int, word string) map[string]int {
	hand := make(map[string]int, 0)
	// copy a map
	for k, v := range dict {
		hand[k] = v
	}
	tempDict := GetFrequencyDict(word)
	for k, v := range dict {
		_, ok := tempDict[k]
		if ok {
			newVal := v - tempDict[k]
			if newVal <= 0 {
				delete(hand, k)
			} else {
				hand[k] = newVal
			}
		}
	}
	return hand
}

var wordList = []string{"apple", "hello"}

func isValidWord(dict map[string]int) bool {

	for _, word := range wordList {
		isValid := reflect.DeepEqual(dict, GetFrequencyDict(word))
		if isValid {
			return true
		}
	}

	return false

}

func main() {
	initGame()
}
