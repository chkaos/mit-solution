package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	fmt.Println("Welcome to the game Hangman!")
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(strSliceLen)
	randomWord := strSlice[randomIndex]

	Hangman(randomWord, strSlice)
}

// Hangman!!!!!
func Hangman(guessedWord string, hintList []string) {
	var inputReader *bufio.Reader
	var input string
	guessTime := 6
	waringTime := 3
	avaliableLetters := "abcdefghijklmnopqrstuvwxyz"
	guessedWordSlice := make([]string, 0)
	pendingGuessWord := GetWordGuessed(guessedWord, guessedWordSlice)

	fmt.Printf("I am thinking of a word that is %d letters long.\n", len(guessedWord))
	fmt.Printf("You have %d waringTime left.\n", waringTime)

	for guessTime > 0 && !IsWordGuessed(guessedWord, guessedWordSlice) {
		fmt.Println("-------------")
		fmt.Printf("You have %d guesses left.\n", guessTime)
		fmt.Printf("Available letters: %s\n", avaliableLetters)
		fmt.Println(guessedWord)
		inputReader = bufio.NewReader(os.Stdin)
		fmt.Println("Please guess a letter: ")
		input, _ = inputReader.ReadString('\n')
		letter := strings.ToLower(string(input[0]))
		ok := isAvaliableLetter(letter)

		if ok {
			hasGuessed := inLetterGuessed(letter, guessedWordSlice)
			if !hasGuessed {

				guessedWordSlice = append(guessedWordSlice, letter)
				isInWord := inWord(letter, guessedWord)
				if isInWord {
					pendingGuessWord = GetWordGuessed(guessedWord, guessedWordSlice)
					avaliableLetters = GetAvaliableLetters(guessedWordSlice)
					fmt.Printf("Good guess: %s\n", pendingGuessWord)

					if IsWordGuessed(guessedWord, guessedWordSlice) {
						fmt.Println("-------------")
						fmt.Println("Congratulations, you won!")
						fmt.Printf("Your total score for this game is: %d\n", guessTime*GetWordUniqueCount(guessedWord))
					}
				} else {
					isVowel := isVowels(letter)
					fmt.Printf("Oops! That letter is not in my word: %s\n", pendingGuessWord)
					if isVowel {
						guessTime -= 2
					} else {
						guessTime--
					}

				}
			} else {
				fmt.Printf("Oops! You've already guessed that letter. You have %d warnings left: %s\n", waringTime, pendingGuessWord)
				if waringTime == 0 {
					guessTime--
				} else {
					waringTime--
				}
			}
		} else {
			if letter == "*" {
				fmt.Printf("Possible word matches are: %s\n", showPossibleMatches(pendingGuessWord, hintList))
			} else if !ok {
				fmt.Printf("Oops! That is not a valid letter. You have %d warnings left: %s\n", waringTime, pendingGuessWord)
				if waringTime == 0 {
					guessTime--
				} else {
					waringTime--
				}
			}
		}
	}

	if guessTime == 0 {
		fmt.Printf("Sorry, you ran out of guesses. The word was %s.\n", guessedWord)
	}

}

// 是否猜中单词
func IsWordGuessed(selectedWord string, letterGuessed []string) bool {
	for _, letter := range selectedWord {
		if !inLetterGuessed(string(letter), letterGuessed) {
			return false
		}
	}

	return true
}

// 获取当前猜到的单词格式
func GetWordGuessed(selectedWord string, letterGuessed []string) (res string) {
	for _, letter := range selectedWord {
		if inLetterGuessed(string(letter), letterGuessed) {
			res += string(letter)
		} else {
			res += "_"
		}
	}

	return
}

// 获取剩余可猜字母
func GetAvaliableLetters(letterGuessed []string) (res string) {

	remainLetters := "abcdefghijklmnopqrstuvwxyz"

	tempMap := make(map[string]bool)
	for _, letter := range letterGuessed {
		_, ok := tempMap[letter] /*如果确定是真实的,则存在,否则不存在 */
		if !ok {
			tempMap[letter] = true
		}
	}

	for _, letter := range remainLetters {
		_, ok := tempMap[string(letter)]
		if !ok {
			res += string(letter)
		}
	}
	return res
}

// 获取单词里唯一字母个数
func GetWordUniqueCount(wordGuessed string) (count int) {

	tempMap := make(map[string]bool)
	for _, letter := range wordGuessed {
		_, ok := tempMap[string(letter)] /*如果确定是真实的,则存在,否则不存在 */
		if !ok {
			tempMap[string(letter)] = true
			count++
		}
	}

	return
}

// 输入是否字母
func isAvaliableLetter(input string) bool {
	return baseHelper(input, "abcdefghijklmnopqrstuvwxyz")
}

// 是否辅音(暂时不用)
func isConsonants(letter string) bool {
	return baseHelper(letter, "bcdfghjklmnpqrstvwxyz")
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

// 是否在单词里
func inWord(letter string, secrectWord string) bool {

	for _, wordLetter := range secrectWord {
		if letter == string(wordLetter) {
			return true
		}
	}
	return false
}

// 是否在猜过的字母切片里
func inLetterGuessed(letter string, letterGuessed []string) bool {

	for _, wordLetter := range letterGuessed {
		if letter == wordLetter {
			return true
		}
	}
	return false
}

func showPossibleMatches(guessWord string, allWords []string) (res string) {

	for _, word := range allWords {
		if len(guessWord) == len(word) && isMatch(guessWord, word) {
			res += word
			res += " "
		}
	}
	if len(res) == 0 {
		res = "No matches found"
	}
	return
}

func isMatch(guessWord string, word string) bool {
	for index, letter := range guessWord {
		if string(letter) != "_" {
			if string(word[index]) != string(letter) {
				return false
			}
		}
	}

	return true
}

func main() {
	initGame()
}
