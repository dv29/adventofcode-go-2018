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

func (c Coords) PrintCoords() {
	for _, x := range c {
		log.Printf("%v\n", x)
	}
}

func (c Coords) Plotter(grid [][]int) {
	for index, x := range c {
		grid[x.x][x.y] = index
	}
	// for x, _ := range grid {
	// 	for y, _ := range grid[x] {
	// 		fmt.Printf("%d", grid[x][y])
	// 	}
	// 	fmt.Println()
	// }
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

func stringToCoord(str string) Coord {
	var coord Coord
	fmt.Sscanf(str, "%d, %d", &coord.x, &coord.y)
	return coord
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

func (c Coord) GetDistance(cb Coord) float64 {
	return math.Abs(float64(c.x-cb.x)) + math.Abs(float64(c.y-cb.y))
}

type MinCoord struct {
	dist       float64
	coordIndex int
}

func (coords Coords) PopulateDistanceMatrix(grid *[][]int) {
	for x, _ := range *grid {
		for y, _ := range (*grid)[x] {
			if (*grid)[x][y] == 0 {
				min := MinCoord{
					(Coord{x, y}).GetDistance(coords[0]),
					0,
				}
				for index, coord := range coords {
					dist := (Coord{x, y}).GetDistance(coord)
					if min.dist < dist {
						min = MinCoord{
							dist,
							index,
						}
					} else if min.dist == dist {
						(*grid)[x][y] = -1
					}
				}
				(*grid)[x][y] = min.coordIndex
			}
		}
	}
}

func PopulateImage(grid [][]int, dc *gg.Context) {
	for x, _ := range grid {
		for y, _ := range grid[x] {
			mx, my := float64(x), float64(y)
			dc.DrawRectangle(mx*multiplier, my*multiplier, multiplier, multiplier)
			ci := float64(grid[x][y])
			dc.SetRGB(ci, 0, 0)
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

	// mx, my := float64(maxX), float64(maxY)
	dc := gg.NewContext(maxX*multiplier, maxY*multiplier)
	// dc.DrawRectangle(0, 0, mx*multiplier, my*multiplier)
	dc.SetRGB(255, 255, 255)
	dc.Clear()
	coords.ImagePlotter(dc)

	coords.PopulateDistanceMatrix(&grid)
	PopulateImage(grid, dc)
	fmt.Printf("%v\n", grid)
	dc.SavePNG("out.png")
}
