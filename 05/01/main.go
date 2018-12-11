package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	// "strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildString(str *string, char string) {
	// fmt.Printf("Str: %s  char: %s\n", *str, char)
	*str += char
	strLen := len(*str)
	// fmt.Printf("%d \n", strLen)
	for i := strLen; i > 1; i-- {
		lastChar := rune(((*str)[i-1]))
		last2ndChar := rune(((*str)[i-2]))
		fmt.Printf("%d %d\t", lastChar, last2ndChar)
	}
}

func main() {
	file, err := os.Open("./input1.txt")
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
}
