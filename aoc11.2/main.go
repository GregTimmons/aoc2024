package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_AGE int = 10


type Cache map[int]map[int]int

func ShouldSplit(i int) bool {
	return len(fmt.Sprintf("%d", i)) % 2 == 0
}

func SplitDigits(i int) (int, int) {
	vs := fmt.Sprintf("%d", i)
	half := int(len(vs) / 2)
	lhs := vs[0:half]
	rhs := vs[half:]

	l, err := strconv.ParseInt(lhs, 10, 64)
	if err != nil { log.Panic(err)}

	r, err := strconv.ParseInt(rhs, 10, 64)
	if err != nil { log.Panic(err)}

	return int(l), int(r)
}

func NewStoneList(f string) []int {
	file, err := os.Open(f)
    if err != nil { log.Panic(err) }
    defer file.Close()
	reader := bufio.NewReader(file)
    buf, _, err := reader.ReadLine()
	if err != nil { log.Panic(err)}
	line := string(buf)
	var stones []int = nil
	for _, v := range strings.Split(line, " ") {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil { log.Panic(err) }
		stones = append(stones, int(vInt))
	}
	return stones
}

func (c Cache) Lookup(value, steps int) (int, bool) {
	if _, ok := c[value]; !ok {
		return 0, false
	}
	v, ok := c[value][steps]
	return v, ok
}

func (c Cache) Store(value, steps, res int) int {
	if _, ok := c[value]; !ok {
		c[value] = make(map[int]int)
	}

	c[value][steps] = res
	return res

}

func (c Cache) ComputeValue(value, steps int) int {
	if steps == 0 {
		return 1
	}

	if value == 0 {
		return c.GetValue(1, steps - 1)
	}

	if ShouldSplit(value) {
		l, r := SplitDigits(value)
		lv := c.GetValue(l, steps - 1)
		rv := c.GetValue(r, steps - 1)
		return lv + rv
	}

	return c.GetValue(value * 2024, steps - 1)
}


func (c Cache) GetValue(value, steps int) int {

	if v, ok := c.Lookup(value, steps); ok {
		return v
	}

	res := c.ComputeValue(value, steps)
	return c.Store(value, steps, res)
}

func main() {
	// var s []int = NewStoneList("sample.data")
	var s []int = NewStoneList("data")
	var c Cache = make(Cache)

	var sum6 int = 0
	var sum25 int = 0
	var sum75 int = 0
	for _, v := range s {
		sum6  += c.GetValue(v, 6)
		sum25 += c.GetValue(v, 25)
		sum75 += c.GetValue(v, 75)
	}

	fmt.Println(sum6)
	fmt.Println(sum25)
	fmt.Println(sum75)






}