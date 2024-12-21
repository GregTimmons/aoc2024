package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ComputeValue(val int, nVal int, allowConcat bool) []int {
	if (!allowConcat) {
		return []int{
			val + nVal,
			val * nVal,
		}
	}
	cVal, err := strconv.ParseInt(fmt.Sprintf("%d%d", val, nVal), 10, 64)
	if err != nil { log.Panic(err); }
	return []int{
		val + nVal,
		val * nVal,
		int(cVal),
	}
}

func ComputeValuesForAllPoss(poss []int, nVal int, allowConcat bool) []int {
	var nValues []int;

	for _, v := range poss {
		for _, c := range ComputeValue(v, nVal, allowConcat) {
			nValues = append(nValues, c)
		}
	}

	return nValues
}


func ProcessLine(s string, allowConcat bool) int {
	args := strings.Split(s, " ")

	targ, err := strconv.ParseInt(strings.Split(args[0], ":")[0], 10, 64)
	if err != nil { log.Panic(err); }

	var ops []int

	for _, opsString := range args[1:] {
		newOp, err := strconv.ParseInt(opsString, 10, 64)
		if err != nil { log.Panic(err); }
		ops = append(ops, int(newOp))
	}

	fmt.Println(ops)
	var possibilites []int = []int{ops[0]}
	ops = ops[1:]
	fmt.Printf("Stating: t=%v, o=%v, p=%v\n", targ, ops, possibilites)
	for {
		possibilites = ComputeValuesForAllPoss(possibilites, ops[0], allowConcat)
		if len(ops) == 1 {
			break
		}
		ops = ops[1:]
	}

	for _, v := range possibilites {
		if v == int(targ) {
			return int(targ)
		}
	}
	return 0
}



func main() {
	f, err := os.Open("./aoc7.data")

	if err != nil {
		log.Panic(err)
	}

	r := bufio.NewScanner(f);


	var total int = 0
	var totalWithConcat int = 0
	for r.Scan() {
		line := r.Text()
		total += ProcessLine(line, false)
		totalWithConcat += ProcessLine(line, true)
	}
	fmt.Println(total)
	fmt.Println(totalWithConcat)

}
