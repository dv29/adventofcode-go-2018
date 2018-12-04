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

func scan(file *os.File, freq int, freqMap map[int]int) (frequency int, fm map[int]int, err error) {
	scanner := bufio.NewScanner(file)
	err = nil
	for scanner.Scan() {
		str := scanner.Text()
		value, err := strconv.Atoi(str)
		check(err)
		freq += value
		if _, ok := freqMap[freq]; ok {
			fmt.Printf("duplicate %d\n", freq)
			return frequency, freqMap, err
		}
		freqMap[freq] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return frequency, freqMap, err
	}
	file.Seek(0, 0)
	return scan(file, freq, freqMap)
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	freq := make(map[int]int)
	frequency := 0
	freq[0] = 1
	frequency, freq, err = scan(file, frequency, freq)

	fmt.Printf("frequency %d\n", frequency)
}
