package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)



type Block struct {
	id int;
	index int;
	fNum int;
	data bool;
	size int;
	nxt *Block
	prv *Block
}

func (b *Block) ToString() string {
	return fmt.Sprintf("[%d: id=%d data=%v fnum=%v size=%v]", b.index, b.id, b.data, b.fNum, b.size)
}

func (b *Block) ToLongString() string {
	var s string = ""
	s += "PREV = "
	if b.prv != nil {
		s += b.prv.ToString()
	}
	s += "\nTHIS = "
	s += b.ToString()
	s += "\nNEXT = "
	if b.nxt != nil {
		s += b.nxt.ToString()
	}
	s += "\n"
	return s
}

func (b *Block) PrintList(rev bool) {
	var node *Block = b
	for {
		if node == nil {
			fmt.Println()
			return
		}
		fmt.Printf("%s\n", node.ToString())
		if rev {
			node = node.prv
		} else {
			node = node.nxt
		}
	}
}

func Swap(l, r *Block) {
	l.id,    r.id   = r.id,    l.id
	l.fNum,  r.fNum = r.fNum,  l.fNum
	l.data,  r.data = r.data,  l.data
	l.size,  r.size = r.size,  l.size
}

var id int = 0
func (b *Block) Split() *Block {

	// fmt.Printf("Splitting Block\n%s\n", b.ToLongString())
	n := &Block{
		id: id,
		fNum: 0,
		data: false,
		size: 0,
		nxt: nil,
		prv: nil,
	}
	id++

	r := b.nxt

	n.nxt = r
	n.prv = b
	b.nxt = n
	r.prv = n

	// Reindex for debugging
	for i := b.nxt; i != nil; i = i.nxt {
		i.index = i.prv.index + 1
	}

	return n
}

func SwapAndSplit(empty, full *Block) {

	extra := empty.Split()

	usedSpace := full.size
	extraSpace := empty.size - full.size

	empty.size = usedSpace
	extra.size = extraSpace

	empty.data = true
	extra.data = false
	full.data = false

	empty.fNum = full.fNum
	extra.fNum = 0
	full.fNum = 0
}


func NewBlockLL(f string) (*Block, *Block) {
	var level int = 0
	var index int = 0
	var head *Block = nil
	var tail *Block = nil
	var hasData bool = true

	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()

    reader := bufio.NewReader(file)

	for {
        char, _, err := reader.ReadRune()
        if err != nil { break }

		value, err := strconv.ParseInt(string(char), 10, 60)
        if err != nil { break }


		nextBlock := &Block{
			id: id,
			index: level,
			fNum: 0,
			data: hasData,
			size: int(value),
			nxt: nil,
			prv: nil,
		}

		if hasData {
			nextBlock.fNum = index
			index++
		}

		if head == nil {
			head = nextBlock
			tail = nextBlock
		} else {
			tail.nxt = nextBlock
			nextBlock.prv = tail
			tail = nextBlock
		}

		id++
		level++
		hasData = !hasData
	}
	return head, tail
}

func ValueList(head *Block) int {
	var sum int = 0
	var idx int = 0

	for b := head; b != nil; b = b.nxt {
		for i := 0; i < b.size; i++ {
			sum += idx * b.fNum
			idx++
		}
	}
	return sum
}

func main() {
	head, tail := NewBlockLL("./aoc9.data")
	// head.PrintList(false)
	// tail.PrintList(true)

	// t := head.nxt.nxt
	// fmt.Printf("Split %v\n", t.id)
	// t.Split()
	// head.PrintList(false)

	// k := head.nxt.nxt.nxt.nxt
	// fmt.Printf("Swap %v %v\n", t.id, k.id)
	// Swap(t, k)
	// head.PrintList(false)

	// j := head.nxt.nxt.nxt.nxt.nxt
	// fmt.Printf("SwapAndSplit %v %v\n", t.id, j.id)
	// SwapAndSplit(t, j)
	// head.PrintList(false)


	for i := tail; i != nil; i = i.prv {
		// if !i.data { continue }

		// fmt.Printf("Sorting: %s\n", i.ToString())
		for j := head; j != nil && j != i; j = j.nxt {
			// fmt.Printf("\tChecking %s\n", j.ToString())
			// If filled, no
			if j.data { continue }
			// If too small, no
			if j.size < i.size { continue }
			// If perfect fit, swap
			if j.size == i.size {
				Swap(j, i)
				break
			}
			// If too big, split
			if j.size > i.size {
				SwapAndSplit(j, i)
				break
			}
		}
	}

	fmt.Printf("Value=%d\n", ValueList(head))

}