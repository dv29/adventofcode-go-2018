package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	for scanner.Scan() {
		str := scanner.Text()
		check(err)
		//  TODO:  <03-12-18, Vora, Deep> // file read starts here
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
