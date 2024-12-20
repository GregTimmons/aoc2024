package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


type Point struct {X int; Y int}
type Grid [][]int

func (p Point) L() Point { return Point{X: p.X - 1, Y: p.Y + 0}}
func (p Point) R() Point { return Point{X: p.X + 1, Y: p.Y + 0}}
func (p Point) U() Point { return Point{X: p.X + 0, Y: p.Y - 1}}
func (p Point) D() Point { return Point{X: p.X + 0, Y: p.Y + 1}}
func (p Point) Hash() string { return fmt.Sprintf("%d|%d", p.Y, p.X) }

func (g Grid) X() int { return len(g[0]) }
func (g Grid) Y() int { return len(g) }
func (g Grid) At(p Point) int { if g.In(p) { return g[p.Y][p.X] } else { return -1 }}
func (g Grid) In(p Point) bool { return p.Y >= 0 && p.Y < g.Y() && p.X >= 0 && p.X < g.X() }


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
			line = append(line, int(char) - 48)
		}
	}
	return grid
}

type Score struct {rating int; score int }

func (g Grid) Walk(p Point) Score {
	var cur_height int = 0
	var walker_locs []Point = []Point{p}
	var walker_next []Point = []Point{}
	// fmt.Printf("Strating Walk From %+v\n", p)

	for {
		// fmt.Printf("h=%d, wl=%+v\n", cur_height, walker_locs)
		for _, i := range walker_locs {
			if (g.At(i.L()) == cur_height + 1) {
				walker_next = append(walker_next, i.L())
			}
			if (g.At(i.R()) == cur_height + 1) { walker_next = append(walker_next, i.R()) }
			if (g.At(i.U()) == cur_height + 1) { walker_next = append(walker_next, i.U()) }
			if (g.At(i.D()) == cur_height + 1) { walker_next = append(walker_next, i.D()) }
		}
		cur_height++
		walker_locs = walker_next
		walker_next = []Point{}

		if cur_height == 9 { break }
		if len(walker_locs) == 0 { break }
	}

	var m map[string]bool = make(map[string]bool)
	for _, v := range walker_locs {
		m[v.Hash()] = true
	}

	fmt.Printf("Score = %d, %d\n\t Locs=%v\n", len(walker_locs),  len(m), walker_locs)
	return Score{rating: len(walker_locs), score: len(m)}
}

func main() {
	// g := NewGrid("./sample.data")
	g := NewGrid("./data")

	var rating int = 0
	var score int = 0
	for y, row := range g {
		for x, h := range row {
			if h == 0 {
				s := g.Walk(Point{Y:y, X:x})
				score = score + s.score
				rating = rating + s.rating
			}
		}
	}

	fmt.Println(score)
	fmt.Println(rating)
}