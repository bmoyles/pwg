package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dictFileName    = "/usr/share/dict/words"
	minFragLen      = 3
	maxFragLen      = 6
	minSpecialChars = 1
	maxSpecialChars = 2
	minNumbers      = 1
	maxNumbers      = 3
	specialChars    = `!@#$%^&*()`
)

var passCount int
var workerPoolSize int

func randomWords(wordList []string, size int, count int) []string {
	var words []string
	for i := 0; i < count; i++ {
		index := rand.Intn(size)
		word := wordList[index]
		if (rand.Intn(10)+1)%2 == 0 {
			word = strings.ToUpper(string(word[0])) + string(word[1:])
		}
		words = append(words, word)
	}
	return words
}

func specialCharFragment() string {
	var specialFrag string
	count := rand.Intn(maxSpecialChars) + minSpecialChars
	for i := 0; i < count; i++ {
		index := rand.Intn(len(specialChars))
		specialFrag = specialFrag + string(specialChars[index])
	}
	return specialFrag
}

func numberFragment() string {
	numberFrag := ""
	count := rand.Intn(maxNumbers) + minNumbers
	for i := 0; i < count; i++ {
		num := rand.Intn(10)
		numberFrag = numberFrag + strconv.Itoa(num)
	}
	return numberFrag
}

func allWords() (words []string, err error) {
	dictFile, err := os.Open(dictFileName)
	if err != nil {
		return words, err
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) >= minFragLen && len(word) <= maxFragLen {
			words = append(words, strings.ToLower(scanner.Text()))
		}
	}
	return words, scanner.Err()
}

func generatePassword(wordList []string, wordListSize int, passChan chan<- string, workChan <-chan int) {
	for range workChan {
		wordFrags := randomWords(wordList, wordListSize, 2)
		specialFrag := specialCharFragment()
		numFrag := numberFragment()
		password := wordFrags[0] + numFrag + specialFrag + wordFrags[1]
		passChan <- password
	}
}

func generatePasswords(count int, workerCount int) (passwords []string, err error) {
	passChan := make(chan string)
	workChan := make(chan int, count)

	wordList, err := allWords()
	if err != nil {
		return nil, err
	}
	wordListSize := len(wordList)

	for w := 0; w < workerCount; w++ {
		go generatePassword(wordList, wordListSize, passChan, workChan)
	}

	for i := 0; i < count; i++ {
		workChan <- i
	}
	close(workChan)

	for len(passwords) < count {
		passwords = append(passwords, <-passChan)
	}
	return passwords, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.IntVar(&passCount, "count", 1, "Number of passwords to generate")
	flag.IntVar(&workerPoolSize, "workers", 10, "Number of goroutines to spawn")
}

func main() {
	flag.Parse()
	passwords, err := generatePasswords(passCount, workerPoolSize)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, password := range passwords {
			fmt.Println(password)
		}
	}
}
