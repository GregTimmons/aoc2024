package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point2 struct{ x, y int}
type Grid6 struct{
	plc [][]Tile6;
	loc Point2;
	dir Dir6;
	cnt int;
	hst map[string]int
}

type Dir6 int
const (
	UP Dir6 = iota
	DWN
	LEFT
	RIGHT
)

type Tile6 int
const (
	OBSTACLE Tile6 = iota
    OPEN
	PLAYER_UP
	USED
)

func (g *Grid6) SetLoc(p Point2) {
	if g.TileAt(p) == OPEN { g.cnt++ }
	g.plc[p.y][p.x] = USED
	g.loc.x = p.x
	g.loc.y = p.y
}

func (g Grid6) IsInBounds(p Point2) bool {
	return p.y >= 0 && p.x >= 0 && p.y < len(g.plc) && p.x < len(g.plc[0])
}

type Event6 int
const (
	MOVE Event6 = iota
	TURN
	LOOP
	END
)


func (g *Grid6) TileAt(p Point2) Tile6 {
	return g.plc[p.y][p.x]
}

func (g *Grid6) Turn() {
	switch g.dir {
	case UP: g.dir = RIGHT
	case RIGHT: g.dir = DWN
	case DWN: g.dir = LEFT
	case LEFT: g.dir = UP
	}
}

func (g *Grid6) GetNext() Point2 {
	n := Point2{ g.loc.x, g.loc.y}
	switch g.dir {
	case UP: n.y += -1
	case DWN: n.y += 1
	case LEFT: n.x += -1
	case RIGHT: n.x += 1
	}
	return n
}

func (g *Grid6) Step() Event6 {

	stepHash := fmt.Sprintf("%d|%d|%d", g.loc.x, g.loc.y, g.dir)

	if _, ok := g.hst[stepHash]; ok {
		return LOOP
	} else {
		g.hst[stepHash] = 1
	}

	next := g.GetNext()

	if !g.IsInBounds(next) {
		return END
	}

	t := g.TileAt(next)

	if t == OBSTACLE {
		g.Turn()
		return TURN
	}

	g.SetLoc(next)
	return MOVE
}
func NewGrid6(f string) Grid6 {
	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()

    reader := bufio.NewReader(file)

	var grid Grid6
	var line []Tile6
	var y int = 0
	var x int = -1
	for {
        char, _, err := reader.ReadRune()
        if err != nil { break }

		if char == '\n' {
			grid.plc = append(grid.plc, line)
			x = -1
			y++
			line = []Tile6{}
		} else {
			x++
			v := map[rune]Tile6 {
				'.': OPEN,
				'#': OBSTACLE,
				'^': PLAYER_UP,
			}[char]

			if v == PLAYER_UP {
				grid.loc.x = x
				grid.loc.y = y
				grid.dir = UP
				grid.cnt = 1
				line = append(line, USED)
			} else {
				line = append(line, v)
			}
		}
	}
	return grid
}

func (g *Grid6) DrawGrid() {
	for _, l := range g.plc {
		for _, t := range l {
			if t == OPEN {
				fmt.Printf(".")
			}
			if t == USED {
				fmt.Printf("X")
			}
			if t == OBSTACLE{
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
}

func (g *Grid6) Process() Event6{
	for {
		e := g.Step()
		if e == LOOP || e == END {
			return e
		}
	}
}


func main() {
	g := NewGrid6("./aoc6.data");
	g.hst = make(map[string]int)
	g.Process()
	fmt.Printf("First processs takes %d steps\n", g.cnt)

	var loop int = 0

	for y := 0; y < len(g.plc); y++ {
		for x := 0; x < len(g.plc[0]); x++ {
			grid := NewGrid6("./aoc6.data");
			grid.hst = make(map[string]int)
			t := grid.TileAt(Point2{x, y})

			if t == OBSTACLE {
				continue
			}

			if (x == grid.loc.x && y == grid.loc.y) {
				continue
			}

			grid.plc[y][x] = OBSTACLE
			e := grid.Process()

			if e == LOOP {
				fmt.Println("LOOP!")
				loop++
			} else {
				fmt.Println("NO LOOP")

			}

		}
	}
	fmt.Printf("First processs takes %d steps\n", loop)

}