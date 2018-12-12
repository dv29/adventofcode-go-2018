package main

import (
	"bufio"
	"fmt"
	// "io"
	"log"
	"os"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isUpper(char string) bool {
	s := []byte(char)[0]
	return s > 64 && s < (64+27)
}

func isLower(char string) bool {
	s := []byte(char)[0]
	return s > 96 && s < (96+27)
}

func react(str *string, char string) {
	*str += char
	for i := len(*str); i > 1; i-- {
		lastChar := string((*str)[i-1])
		last2ndChar := string((*str)[i-2])
		if strings.ToLower(lastChar) == strings.ToLower(last2ndChar) && ((isLower(lastChar) && isUpper(last2ndChar)) || (isUpper(lastChar) && isLower(last2ndChar))) {
			*str = (*str)[:i-2]
			i--
		} else {
			break
		}
	}
}

func reactMapper(str string, cb func(*string, string)) int {
	var resultString string
	for _, s := range str {
		react(&resultString, string(s))
	}
	return len(resultString)
}

func removeUnitReact(str string) {
	max := 0
	for i := 0; i < 26; i++ {
		r := regexp.MustCompile(fmt.Sprintf("(%s|%s)", string(i+65), strings.ToLower(string(i+65))))
		s := r.ReplaceAllString(str, "")
		x := reactMapper(s, react)
		if max > x || max == 0 {
			max = x
		}
	}
	fmt.Printf("%d\n", max)
}

func main() {
	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isString := scanner.Scan()
	if isString {
		str := scanner.Text()
		removeUnitReact(str[:len(str)])
	}
}
