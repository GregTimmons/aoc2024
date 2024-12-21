package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Grid [][]int
type Coord struct { x int; y int}

func (g Grid) X() int { return len(g[0]) }
func (g Grid) Y() int { return len(g)    }

func NewGrid(f string) Grid {
	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()

    reader := bufio.NewReader(file)

	var grid [][]int
	var line []int
	for {
        char, _, err := reader.ReadRune()
        if err != nil { break }

		if char == '\n' {
			grid = append(grid, line)
			line = []int{}
		} else {
			v := map[rune]int {
				'X': 0,
				'M': 1,
				'A': 2,
				'S': 3,
			}[char]
			line = append(line, v)
		}
	}
	return grid
}

func (g Grid) At (p Coord) int {
	return g[p.y][p.x]
}


func (g Grid) OutOfBounds(p Coord) bool {
	return p.x < 0 || p.x >= g.X() || p.y < 0 || p.y >= g.Y()
}

func SearchType1_Direction(g Grid, pos, step Coord) bool {
	for i := 0; i < 4; i++ {
		newPos := Coord{
			x: (pos.x + (step.x * i)),
			y: (pos.y + (step.y * i)),
		}
		if g.OutOfBounds(newPos) || g.At(newPos) != i {
			return false
		}
	}
	return true
}

// Search in all directions if XMAS is spelled, return the count. 
func SearchType1(g Grid, pos Coord) int {
	var count int = 0
	for xDir := -1; xDir <= 1; xDir++ {
		for yDir := -1; yDir <= 1; yDir++ {
			if SearchType1_Direction(g, pos, Coord{x: xDir,  y: yDir}) {
				count++
			}
		}
	}
	return count
}

// Search for a X made of MAS'es
func SearchType2(g Grid, pos Coord) int {
	if g.At(pos) != 2 {
		return 0
	}

	if (g.OutOfBounds(Coord{pos.x - 1, pos.y - 1}) ||
		g.OutOfBounds(Coord{pos.x + 1, pos.y - 1}) ||
		g.OutOfBounds(Coord{pos.x + 1, pos.y + 1}) ||
		g.OutOfBounds(Coord{pos.x - 1, pos.y + 1})) {
		return 0
	}

	tl := g.At(Coord{pos.x - 1, pos.y - 1})
	tr := g.At(Coord{pos.x + 1, pos.y - 1})
	br := g.At(Coord{pos.x + 1, pos.y + 1})
	bl := g.At(Coord{pos.x - 1, pos.y + 1})

	criss := (tl == 1 || tl == 3) && (br == 1 || br == 3) && tl != br
	cross := (tr == 1 || tr == 3) && (bl == 1 || bl == 3) && tr != bl

	if criss && cross {
		return 1
	}
	return 0
}

// Search the grid with a 'check' function to count matches at each pos.
func (g Grid) Search(check func(Grid, Coord) int) int {
	var count int = 0
	for xInd := 0; xInd < g.X(); xInd++ {
		for yInd := 0; yInd < g.Y(); yInd++ {
			count += check(g, Coord{x: xInd, y: yInd})
		}
	}
	return count
}

func main() {

	grid := NewGrid("./aoc4.data")
	fmt.Printf("Grid[%d, %d]\n", grid.X(), grid.Y())
	fmt.Printf("Count = %d \n", grid.Search(SearchType1))
	fmt.Printf("Count = %d \n", grid.Search(SearchType2))
}


