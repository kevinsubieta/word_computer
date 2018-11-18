package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"os"
)

type Channel struct {
	name string
}

var alphabetic = make(map[string]map[string]int)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	var contentFile = readTxtAndReturnContentAsString(os.Args[1])
	sortAndCountWordsInContent(contentFile)
	printResultInConsole()
}

func readTxtAndReturnContentAsString(filePath string) string {
	contentAsString, error := ioutil.ReadFile(filePath)
	checkIfExistAnyErrorOnOpenFile(error)
	return string(contentAsString)
}

func checkIfExistAnyErrorOnOpenFile(e error) {
	if e != nil {
		panic(e)
	}
}

func sortAndCountWordsInContent(contentFile string) {
	var newWord = ""
	var wordMap = make(map[string]int)
	for i := 0; i < len(contentFile); i++ {
		character := string(contentFile[i])
		if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(character) {
			newWord += string(character)
		} else if len(newWord) > 0 {
			newWord = strings.ToLower(newWord)
			var index = strings.ToLower(string(newWord[0]))
			wordMap = alphabetic[index]
			if wordMap == nil {
				wordMap = make(map[string]int)
			}
			wordMap[newWord] = wordMap[newWord] + 1
			alphabetic[index] = wordMap
			newWord = ""
		}
	}
}

func printResultInConsole() {
	var lettersAlpabeticOrdered = sortLettersAlpabetically(alphabetic)
	for i := 0; i < len(lettersAlpabeticOrdered); i++ {
		var arrayOfWords = alphabetic[lettersAlpabeticOrdered[i]]
		var arrayOfWordsSorted = sortCollectionAlpabetically(arrayOfWords)
		for j := 0; j < len(arrayOfWordsSorted); j++ {
			var orderedWord = arrayOfWordsSorted[j]
			var valueOfTheWord = arrayOfWords[orderedWord]
			fmt.Println(orderedWord + " " + strconv.Itoa(valueOfTheWord))
		}
	}
}


func sortLettersAlpabetically(mapOfLetters map[string]map[string]int) []string {
	keys := make([]string, 0, len(mapOfLetters))
	for k := range mapOfLetters {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}


func sortCollectionAlpabetically(mapOfWords map[string]int) []string {
	keys := make([]string, 0, len(mapOfWords))
	for k := range mapOfWords {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
