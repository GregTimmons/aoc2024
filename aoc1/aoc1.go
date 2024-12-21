package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Counts struct {
	Left int64
	Right int64
}

func main() {
	valuesMap := make(map[int64]*Counts)
	f, err := os.Open("./aoc1.data")

	if err != nil {
		log.Panic(err)
	}

	r := bufio.NewScanner(f);

	for r.Scan() {
		args := strings.Split(r.Text(), "  ")

		left, err := strconv.ParseInt(strings.Trim(args[0], " "), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		right, err := strconv.ParseInt(strings.Trim(args[1], " "), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("l=%v r=%v\n", left, right)

		if valuesMap[left] == nil {
			valuesMap[left] = &Counts{0, 0}
		}

		valuesMap[left].Left += 1

		if valuesMap[right] == nil {
			valuesMap[right] = &Counts{0, 0}
		}

		valuesMap[right].Right += 1
	}

	var i int64 = 0
	for number, c := range valuesMap {
		if c.Left == 0 { continue }
		i += (number * c.Right)
		fmt.Println(number, c)
	}
	fmt.Println(i)


}