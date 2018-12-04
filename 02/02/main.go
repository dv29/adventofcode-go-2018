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

func compareStrings(a []string, b []string) (errourCount int) {
	errorCount := 0
	for i, val := range a {
		if val != b[i] {
			errorCount++
			if errourCount == 2 {
				break
			}
		}
	}
	return errorCount
}

func findStringWithOneDiff(values [][]string) (a, b []string) {
	for i, val := range values {
		for j := i; j < len(values); j++ {
			a, b = val, values[j]
			if errorCount := compareStrings(a, b); errorCount == 1 {
				return
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values [][]string

	for scanner.Scan() {
		str := scanner.Text()
		check(err)
		values = append(values, strings.Split(str, ""))
	}

	a, b := findStringWithOneDiff(values)

	var str []string
	for i, val := range a {
		if val == b[i] {
			str = append(str, val)
		}
	}
	fmt.Printf("a: %v, b: %v\n", strings.Join(a, ""), strings.Join(b, ""))
	fmt.Printf("str: %s\n", strings.Join(str, ""))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
