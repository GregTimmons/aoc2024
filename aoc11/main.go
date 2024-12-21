package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type StoneRule func(*Stone) bool
type Stone struct {
	a int // Age of the stone.
	v int
	n *Stone
	p *Stone
}

// Split to the left so that it is forward iteration safe.
func (b *Stone) Split(lv, rv int) {
	a := b.p
	n := &Stone { v: 0, n: nil, p: nil }

	b.p = n
	n.n = b

	if a != nil {
		a.n = n
		n.p = a
	}

	n.v = lv
	b.v = rv
}

func StoneZeroAddOne(s *Stone) bool {
	if s.v != 0 { return false }
	s.v = 1
	return true
}

func StoneSplitEvenDigits(s *Stone) bool {
	vs := fmt.Sprintf("%d", s.v)
	if len(vs) % 2 == 1 {
		return false
	}

	half := int(len(vs) / 2)
	lhs := vs[0:half]
	rhs := vs[half:]

	lhsInt, err := strconv.ParseInt(lhs, 10, 64)
	if err != nil { log.Panic(err)}

	rhsInt, err := strconv.ParseInt(rhs, 10, 64)
	if err != nil { log.Panic(err)}

	s.Split(int(lhsInt), int(rhsInt))
	return true
}

func StoneMult2024(s *Stone) bool {
	s.v = s.v * 2024
	return true
}
var Rules []StoneRule = []StoneRule{
	StoneZeroAddOne,
	StoneSplitEvenDigits,
	StoneMult2024,
}

func NewStoneList(f string) *Stone{
	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()
	reader := bufio.NewReader(file)
    buf, _, err := reader.ReadLine()
	if err != nil { log.Panic(err)}
	line := string(buf)

	var head *Stone = nil
	var tail *Stone = nil

	for _, v := range strings.Split(line, " ") {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil { log.Panic(err) }

		ns := &Stone{
			a: 0,
			v: int(vInt),
			n: nil,
			p: nil,
		}

		if head == nil { head = ns }
		if tail == nil { tail = ns }
		if tail != ns {
			tail.n = ns
			ns.p = tail
			tail = ns
		}
	}
	return head
}


func Blink(s *Stone) {

	for i := s; i != nil; i = i.n {
		// fmt.Printf("Blink %+v <- %d -> %+v\n", s.p, s.v, s.n)
		for _, rule := range Rules {
			res := rule(i)
			if res { break }
		}
	}
}



func WriteStones(s *Stone) string {
	var t string = ""
	for i := s; i != nil; i = i.n {
		t += fmt.Sprintf("%d ", i.v)
	}
	return t
}

func CountStones(s *Stone) int {
	var c int = 0
	for i := s; i != nil; i = i.n { c++ }
	return c
}

func main() {

	// s := NewStoneList("sample.data")
	s := NewStoneList("data")

	h := &Stone{}
	h.n = s
	s.p = h

	fmt.Printf("I=0, C=%d, S=%s\n", CountStones(h.n), WriteStones(h.n))

	for i := 1; i <= 75; i++ {
		Blink(h.n)
		// fmt.Printf("I=%d, C=%d, S=%s\n", i, CountStones(h.n), WriteStones(h.n))
		fmt.Printf("I=%d, C=%d \n", i, CountStones(h.n))
	}
}