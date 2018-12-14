package main

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"math"
	"os"
)

const multiplier = 20

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Coord struct {
	x, y int
}
type Coords []Coord

func stringToCoord(str string) Coord {
	var coord Coord
	fmt.Sscanf(str, "%d, %d", &coord.x, &coord.y)
	return coord
}

func (c Coords) PrintCoords() {
	for _, x := range c {
		log.Printf("%v\n", x)
	}
}

func (c Coords) GetBounds() (maxX, maxY int) {
	maxX = c[0].x
	maxY = c[0].y
	for _, x := range c {
		if maxX < x.x {
			maxX = x.x
		}
		if maxY < x.y {
			maxY = x.y
		}
	}
	return
}

func (c Coords) Plotter(grid [][]int) {
	for index, x := range c {
		grid[x.x][x.y] = index + 1
	}
}

func (c Coord) GetDistance(cb Coord) float64 {
	return math.Abs(float64(c.x-cb.x)) + math.Abs(float64(c.y-cb.y))
}

type MinCoord struct {
	dist       float64
	coordIndex int
}

func (coords Coords) PopulateDistanceMatrix(grid *[][]int) (m map[int]int) {
	m = make(map[int]int)
	for x, _ := range *grid {
		for y, _ := range (*grid)[x] {
			if (*grid)[x][y] == 0 {
				min := MinCoord{
					(Coord{x, y}).GetDistance(coords[0]),
					1,
				}
				for index, coord := range coords {
					if dist := (Coord{x, y}).GetDistance(coord); min.dist > dist {
						min = MinCoord{
							dist,
							index,
						}
					} else if min.dist == dist {
						min = MinCoord{
							dist,
							-1,
						}
					}
				}
				(*grid)[x][y] = min.coordIndex + 1
				if _, ok := m[min.coordIndex]; ok {
					m[min.coordIndex] += 1
				} else {
					m[min.coordIndex] = 1
				}
			}
		}
	}
	return
}

func (c Coords) ImagePlotter(dc *gg.Context) {
	for index, x := range c {
		mx, my := float64(x.x), float64(x.y)
		dc.DrawRectangle(mx*multiplier, my*multiplier, multiplier*2, multiplier*2)
		ci := float64(index)
		dc.SetRGB(ci, ci+mx/255.0, ci+my/255.0)
		dc.Fill()
	}
}

func PopulateImage(grid [][]int, dc *gg.Context) {
	for x, _ := range grid {
		for y, _ := range grid[x] {
			mx, my := float64(x), float64(y)
			dc.DrawRectangle(mx*multiplier, my*multiplier, multiplier, multiplier)
			ci := float64(grid[x][y])
			dc.SetRGB(ci, ci+ci/ci, ci+ci/ci)
			dc.Fill()
		}
	}
}

func matrix(x, y int) [][]int {
	m := make([][]int, x)
	for i := range m {
		m[i] = make([]int, y)
	}
	return m
}

func Max(m map[int]int) (max int) {
	for _, x := range m {
		if max < x {
			max = x
		}
	}
	return
}

func main() {
	file, err := os.Open("./input")
	errorHandler(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var coords Coords

	for scanner.Scan() {
		str := scanner.Text()
		coords = append(coords, stringToCoord(str))
	}

	// coords.PrintCoords()
	maxX, maxY := coords.GetBounds()
	log.Printf("Bounds => x: %d, y: %d", maxX, maxY)

	grid := matrix(maxX+1, maxY+1)
	coords.Plotter(grid)
	m := coords.PopulateDistanceMatrix(&grid)

	max := Max(m)
	log.Printf("%d", max)
	// mx, my := float64(maxX), float64(maxY)
	dc := gg.NewContext(maxX*multiplier, maxY*multiplier)
	// dc.DrawRectangle(0, 0, mx*multiplier, my*multiplier)
	dc.SetRGB(255, 255, 255)
	dc.Clear()
	// coords.ImagePlotter(dc)
	PopulateImage(grid, dc)

	dc.SavePNG("out.png")
}
