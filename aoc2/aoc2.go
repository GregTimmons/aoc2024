package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
type State string

const (
	FIRST_READ 	State = "first_read"
	FIRST_STEP 	State = "first_step"
	READ_DESC  	State = "read_desc"
	READ_ASC	State = "read_asc"
)

func checkWithSkip(args []string, skip int) bool {
	var state State = FIRST_READ
	var valid bool = true
	var err error
	var prev int64
	var curr int64
	fmt.Println("NEW")
	for i := 0; i < len(args); i++ {
		if i == skip { continue }
		if state == FIRST_READ {
			prev, err = strconv.ParseInt(args[i], 10, 64)
			if err != nil { log.Panic(err) }
			state = FIRST_STEP
			fmt.Printf("comp %v state=%v prev=%v curr=%v\n", i, state, prev, curr, )
			continue
		}

		curr, err = strconv.ParseInt(args[i], 10, 64)
		if err != nil { log.Panic(err) }

		if state == FIRST_STEP {
			if prev > curr {
				state = READ_DESC
			} else {
				state = READ_ASC
			}
		}

		if state == READ_ASC {
			step := curr - prev
			if step > 3 || step < 1 {
				valid = false
				break
			}
		}

		if state == READ_DESC {
			step := prev - curr
			if step > 3 || step < 1 {
				valid = false
				break
			}
		}



		fmt.Printf("comp %v state=%v prev=%v curr=%v\n", i, state, prev, curr )
		prev = curr
	}
	fmt.Printf("RESULT=%v\n", valid)
	return valid
}

func main() {
	f, err := os.Open("./aoc2.data")

	if err != nil {
		log.Panic(err)
	}

	r := bufio.NewScanner(f);
	var safeLevels int = 0

	for r.Scan() {
		args := strings.Split(r.Text(), " ")
		if checkWithSkip(args, -1) {
			safeLevels++
		}

		for i := 0; i < len(args); i++ {

		}
	}

	fmt.Println(safeLevels)

}