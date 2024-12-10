package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vec2 struct { x, y int }

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{
		v1.x - v2.x,
		v1.y - v2.y,
	}
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		v1.x + v2.x,
		v1.y + v2.y,
	}
}

func (v1 Vec2) Hash() string {
	return fmt.Sprintf("%d|%d", v1.x, v1.y)
}

type Grid8 struct {
	X int
	Y int
	A map[rune][]Vec2 // Locations of antennas. No need for an actual grid in data.
	R map[string]bool 	// Hash of antinode locations -- map to keep this deduped. Hash is "${x}|${y}"
	S map[string]bool 	// Hash of antinode locations -- map to keep this deduped. Hash is "${x}|${y}"
}

func NewGrid8(f string) Grid8 {
	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()

	ng := Grid8{}
	ng.A = make(map[rune][]Vec2)
	ng.R = map[string]bool{}
	ng.S = map[string]bool{}
    reader := bufio.NewReader(file)

	var x int = 0
	var y int = 0
	for {
        char, _, err := reader.ReadRune()
        if err != nil { break }
		if char == '\n' {
			ng.X = x
			x = 0
			y++
			continue
		}
		x++
		if char == '.' {
			continue
		}
		ng.A[char] = append(ng.A[char], Vec2{x-1, y})
	}

	ng.Y = y
	return ng
}

func (g *Grid8) IsInBounds(v Vec2) bool {
	return v.x >= 0 && v.y >= 0 && v.x < g.X && v.y < g.Y
}

func (g *Grid8) ProcessAntenna(name rune) {
	var loc []Vec2 = g.A[name]

	fmt.Println(name)
	fmt.Println(loc)
	for _, lhs := range loc {
		for _, rhs := range loc {
			if lhs.Hash() == rhs.Hash() {
				continue
			}
			var step Vec2 = lhs.Sub(rhs)
			var pos Vec2 = lhs

			fmt.Printf("\t%v => %v = %v ... %v\n", lhs, rhs, pos, g.IsInBounds(pos))

			if (g.IsInBounds(pos.Add(step))) {
				g.R[pos.Add(step).Hash()] = true
			}

			for {
				fmt.Println(pos)
				if !g.IsInBounds(pos) {
					break
				}
				g.S[pos.Hash()] = true
				pos = pos.Add(step)
			}
		}
	}
}

func (g *Grid8) Process() {
	for k := range g.A {
		g.ProcessAntenna(k)
	}
}

func main() {
	g := NewGrid8("./aoc8.data")
	g.Process()
	fmt.Println(g)
	fmt.Println(len(g.R))
	fmt.Println(len(g.S))


}