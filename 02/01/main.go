package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	letterCount := make(map[int]int)

	for scanner.Scan() {
		str := scanner.Text()
		check(err)
		charCount := make(map[string]int)
		for _, char := range strings.Split(str, "") {
			charCount[char] += 1
		}
		// fmt.Printf("charCount: %v\n", charCount)
		uniqueCount := make(map[int]int)
		for _, v := range charCount {
			if _, ok := uniqueCount[v]; !ok {
				uniqueCount[v] = 1
				letterCount[v] += 1
			}
		}
		// fmt.Printf("uniqueCount: %v\n", uniqueCount)
	}

	total := 1
	for k, v := range letterCount {
		if k != 1 {
			total *= v
		}
	}

	fmt.Printf("letterCount: %v\n", letterCount)
	fmt.Printf("total: %v\n", total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
