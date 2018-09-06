package gopwg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	minFragmentLength = 3
	maxFragmentLength = 6
	numFragments      = 2
	minSpecialChars   = 1
	maxSpecialChars   = 2
	minNumbers        = 1
	maxNumbers        = 3
	asciiAlphaRegex   = "^[a-z]+$"
)

var (
	specialChars = [...]rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')'}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func randInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func generateSpecialCharacters() string {
	var specials bytes.Buffer
	numSpecials := randInt(minSpecialChars, maxSpecialChars+1)
	for i := 0; i < numSpecials; i++ {
		specials.WriteRune(specialChars[randInt(0, len(specialChars))])
	}
	return specials.String()
}

func generateNumbers() string {
	var numbers bytes.Buffer
	numNumbers := randInt(minNumbers, maxNumbers+1)
	for i := 0; i < numNumbers; i++ {
		numbers.WriteString(strconv.Itoa(randInt(0, 10)))
	}
	return numbers.String()
}

func isAsciiAlpha(s string) bool {
	match, err := regexp.MatchString(asciiAlphaRegex, s)
	if err != nil {
		return false
	}
	return match
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func generateWords(wordList io.ReadSeeker, wordListSize int64) ([]string, error) {
	words := make([]string, numFragments)
	bufReader := bufio.NewReader(wordList)
	for i := 0; i < numFragments; i++ {
		for {
			// seek to a random offset in the reader
			newOffset := randInt64(0, wordListSize)
			_, err := wordList.Seek(newOffset, 0)
			if err != nil {
				return nil, err
			}
			bufReader.Reset(wordList)
			// toss the first call to ReadBytes() in case we are mid-line
			_, err = bufReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					continue
				}
				return nil, err
			}
			// now read a real line
			rawLine, err := bufReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					continue
				}
				return nil, err
			}
			word := strings.ToLower(strings.TrimSpace(rawLine))
			if isAsciiAlpha(word) {
				wordLen := len(word)
				if wordLen >= minFragmentLength && wordLen <= maxFragmentLength {
					if (randInt(0, 10)+1)%2 == 0 {
						word = upperFirst(word)
					}
					words[i] = word
					break
				}
			}
		}
	}
	return words, nil
}

type PasswordGenerator struct {
	words     io.ReadSeeker
	wordsSize int64
}

func NewPasswordGenerator(words io.ReadSeeker, wordsSize int64) *PasswordGenerator {
	return &PasswordGenerator{words: words, wordsSize: wordsSize}
}

func (pwg *PasswordGenerator) String() string {
	return fmt.Sprintf("PasswordGenerator{wordsSize:%v}", pwg.wordsSize)
}

func (pwg *PasswordGenerator) Generate(count int) ([]string, error) {
	passwords := make([]string, count)
	for i := 0; i < count; i++ {
		numbers := generateNumbers()
		specials := generateSpecialCharacters()
		words, err := generateWords(pwg.words, pwg.wordsSize)
		if err != nil {
			return nil, err
		}
		passwords[i] = words[0] + numbers + specials + words[1]
	}
	return passwords, nil
}
