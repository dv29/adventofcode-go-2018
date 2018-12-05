package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ClaimRectangle struct {
	id     uint
	x      uint
	y      uint
	width  uint
	height uint
}

func mapToClaimRectangle(str string, re *regexp.Regexp) (cr ClaimRectangle) {
	value := re.FindStringSubmatch(str)
	id, err := strconv.Atoi(value[1])
	check(err)
	x, err := strconv.Atoi(value[2])
	check(err)
	y, err := strconv.Atoi(value[3])
	check(err)
	width, err := strconv.Atoi(value[4])
	check(err)
	height, err := strconv.Atoi(value[5])
	check(err)

	return ClaimRectangle{
		uint(id),
		uint(x),
		uint(y),
		uint(width),
		uint(height),
	}
}

const FABRIC_SIZE = 1000

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<width>\d+)x(?P<height>\d+)`)
	var fabric [FABRIC_SIZE][FABRIC_SIZE]int
	count := 0
	var claims []ClaimRectangle

	for scanner.Scan() {
		str := scanner.Text()
		check(err)
		claims = append(claims, mapToClaimRectangle(str, regex))
	}

	// part 1
	for _, cr := range claims {
		for w := uint(0); w < cr.width; w++ {
			for h := uint(0); h < cr.height; h++ {
				val := fabric[cr.x+w][cr.y+h]
				if val == 0 {
					fabric[cr.x+w][cr.y+h] = int(cr.id)
				} else if val != -1 {
					fabric[cr.x+w][cr.y+h] = -1
					count++
				}
			}
		}
	}

	// part 2
	for _, cr := range claims {
		isUnique := true
		for w := uint(0); w < cr.width; w++ {
			for h := uint(0); h < cr.height; h++ {
				val := fabric[cr.x+w][cr.y+h]
				if val != int(cr.id) {
					isUnique = false
				}
			}
		}

		if isUnique {
			fmt.Printf("uinque: %d\n", cr.id)
		}
	}

	// for i := 0; i < FABRIC_SIZE; i++ {
	// 	for j := 0; j < FABRIC_SIZE; j++ {
	// 		fmt.Printf("\t%d\t", fabric[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }

	fmt.Printf("count: %d\n", count)

	fmt.Println("")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
