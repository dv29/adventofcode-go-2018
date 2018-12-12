package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func buildString(str *string, char string) {
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

func main() {
	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewReader(file)
	var resultString string

	for {
		if c, _, err := scanner.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// fmt.Printf("%s\n", string(c))
			buildString(&resultString, string(c))
		}
	}

	// reducing by 1 for \n
	fmt.Printf("%s\t%d\n", resultString, len(resultString)-1)
}
