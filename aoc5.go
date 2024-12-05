package main

import (
	"aoc2024/m/v2/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Turns out the rule are NOT recursive.
// Wrote code that was WAY harder then it needed to be. Fun stuff.

type MOrder struct { L int64; R int64 }
type MBook []int64

// Precedent map says that for page LHS, must have come before page RHS.
type MPrecedentMap map[int64][]int64


type QEle struct { val int64; depth int }
func QueuePop(s []QEle) (*QEle, []QEle, bool) {
    if len(s) == 0 {
		return nil, s, false
    }
    return &s[len(s)-1], s[:len(s)-1], true
}

func QueueContains(s []QEle, t int64) bool {
	for _, v := range s {
		if v.val == t{
			return true
		}
	}
	return false
}
// For a given page A, get the set of all pages that must NOT come before it.
// This is recursive, If A << B, and B << C, then A << C.
// If P(A) = the set of all pages one level below A,
//    then this func returns F(A) = Σ[1,len(P(A)](F(P(A)_i))
func GetPrecedingSet(m MPrecedentMap, root int64) []int64 {
	var precSetVals []int64
	var precSet []QEle
	var queue []QEle
	var next *QEle = &QEle{root, 0}
	var hasMore bool = true
	var iter int = 0

	for {

		// fmt.Printf("\n\n[%d]\nQUEUE = %v\nnext = %v\nPRECSET=%v\nLEAVES=%v\n", iter, queue, next, precSet, m[next.val])
		iter++
		if QueueContains(precSet, next.val) {
			continue
		}

		if next.val != root {
			precSet = append(precSet, QEle{next.val, next.depth})
			precSetVals = append(precSetVals, next.val)
		}

		if next.depth == 0 {
			for _, val := range m[next.val] {
				if !QueueContains(precSet, val) && !QueueContains(queue, val) {
					queue = append(queue, QEle{val, next.depth + 1})
				}
			}
		}

		next, queue, hasMore = QueuePop(queue)
		if !hasMore { return precSetVals}
	}
}

// For a given manual, for each page:
//   On the left get the pages that have come before this one
//   On the right, use GetPrecedingSet to get the pages that must NOT have come before this one
//   If left and right intersect this is an invalid manual.
func InvalidIndex(m MBook, p MPrecedentMap) int {
	var prevPage []int64
	for i, v := range m {
		prec := GetPrecedingSet(p, v)
		intesects := util.SliceIntersects[int64](prevPage, prec)
		// fmt.Printf("Page %d: %v, Intersect = %v, Prev = %v, precedents = %v\n", i, v, intesects, prevPage, prec)
		if intesects { return i }
		prevPage = append(prevPage, v)
	}
	return -1
}


type Manual struct {
	precedents map[int64][]int64
	order []MOrder;
	pages []MBook;
}

func (m *Manual) initMap() {
	for _, v := range m.order{
		m.precedents[v.L] = append(m.precedents[v.L], v.R)
	}
}

func NewManual(f string) *Manual {

	precedents := make(map[int64][]int64)
	newMan := &Manual{
		precedents,
		[]MOrder{},
		[]MBook{},
	}


	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()

	var isManuals bool = false

    reader := bufio.NewReader(file)
	for {
        buf, _, err := reader.ReadLine()
		if err != nil { break; }

		line := string(buf)


		if strings.Trim(string(line), " \t") == "" {
			isManuals = true
			fmt.Println("OK")
		} else if isManuals {
			pages := strings.Split(line, ",")
			var newBook MBook

			for _, val := range pages {
				intPage, err := strconv.ParseInt(val, 10, 64)
				if err != nil { log.Panic(err) }
				newBook = append(newBook, intPage)
			}
			newMan.pages = append(newMan.pages, newBook)

		} else {
			ords := strings.Split(line, "|")
			l, err := strconv.ParseInt(ords[0], 10, 64)
			if err != nil { log.Panic(err) }

			r, err := strconv.ParseInt(ords[1], 10, 64)
			if err != nil { log.Panic(err) }

			newMan.order = append(newMan.order, MOrder{L: l, R: r})
		}
	}
	newMan.initMap()
	return newMan
}


func main() {
	// man := NewManual("./aoc5.sample.data")
	man := NewManual("./aoc5.data")
	var sum int64 = 0
	var sum2 int64 = 0
	var indexForViolation int = 0
	for _, m := range man.pages {
		fmt.Println("Check Manuals")
		indexForViolation = InvalidIndex(m, man.precedents)
		if indexForViolation == -1 {
			centerValue := m[len(m) / 2]
			sum += centerValue
			continue
		}

		for {
			m[indexForViolation], m[indexForViolation-1] = m[indexForViolation-1], m[indexForViolation]
			indexForViolation = InvalidIndex(m, man.precedents)
			if indexForViolation == -1 {
				centerValue := m[len(m) / 2]
				sum2 += centerValue
				break
			}
		}
	}


	fmt.Println(sum)
	fmt.Println(sum2)
}