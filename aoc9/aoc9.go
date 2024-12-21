package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Decode(s string) []int {
	data := []int{}
	for i, v := range []rune(s) {
		num, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil { log.Panic(err)}
		var val int
		if i % 2 == 1 {
			val = -1
		} else {
			val = (i + 1) / 2
		}

		for k := 0; k < int(num); k++ {
			data = append(data, val)
		}
	}
	return data
}

func AttemptCopy(s *[]int, at, value, size int) {

	for x := 0; x < size; x}}

}

func main () {
	file, err := os.Open("./aoc9.data")
    if err != nil { log.Panic(err) }
    defer file.Close()

	t := bufio.NewScanner(file)
	t.Scan()

	orig := t.Text()
	s := Decode(orig)


	var i int = 0
	var j int = len(s) - 1

	for {
		if i == j {
			fmt.Printf("END\n")
			break
		}

		if s[i] != -1 {
			i++
			continue
		}

		if s[j] == -1 {
			j--
			continue
		}

		s[i], s[j] = s[j], s[i]
	}

	var chkSum int = 0

	for i, v := range s {
		if v == -1 { continue}
		if err != nil { log.Panicln(err) }
		chkSum += i * s[i]
	}

	fmt.Println(chkSum)

	s = Decode(orig)
	i = 0
	j = len(s) - 1

	var bVal int = -1
	var bSize int = 0
	chkSum = 0

	for {
		if i == j {
			break
		}

		if s[i] != -1 {
			i++
			continue
		}

		if s[j] == -1 && bVal != -1 {
			AttemptCopy(&s, i, bSize, bVal)
			bVal = -1
			bSize = 0
		}

		if s[j] == -1 && bVal == -1 {
			j--
			continue

		}

		if s[j] == bVal {
			bSize++
			j--
			continue
		}

		if s[j] != -1 {
			bVal = s[j]
			bSize = 1
			j--
			continue
		}

		s[i], s[j] = s[j], s[i]
	}


	for i, v := range s {
		if v == -1 { continue}
		if err != nil { log.Panicln(err) }
		chkSum += i * s[i]
	}

	fmt.Println(chkSum)







}