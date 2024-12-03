package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
type MulPair struct { l int64; r int64 }
type Parse struct {
	MulPair []*MulPair
	Total int64
	Char rune
	MulEnabled bool
	ActionTable ActionTable
	LookAheadTable LookAheadTable
	Iteration int
	NxtState MulState
	CurState MulState
	LBuffer string
	RBuffer string
}
type MulState string
type ActionTable map[MulState]func(*Parse, rune)
type LookAheadTable map[MulState]map[string]MulState

const GLOBAL string = ":global:"
const DIGIT string = ":digit:"
const EMPTY string = ":empty:"
const (
	E MulState = ""
	SCAN MulState = "scan"
	ON_M MulState = "on_m"
	ON_D MulState = "on_d"
	ON_O MulState = "on_o"
	ON_U MulState = "on_u"
	ON_L MulState = "on_l"
	ON_DO_OPEN MulState = "on_do_open"
	ON_DO_CLOSE MulState = "on_do_close"
	ON_N MulState = "on_n"
	ON_APOS MulState = "on_apos"
	ON_T MulState = "on_t"
	ON_DONT_OPEN = "on_done_open"
	ON_DONT_CLOSE = "on_dont_close"
	ON_MUL_OPEN MulState = "on_mul_open"
	ON_MUL_NUM_1_1 MulState = "on_num_1_1"
	ON_NUM_1_2 MulState = "on_num_1_2"
	ON_NUM_1_3 MulState = "on_num_1_3"
	ON_COMMA MulState = "on_comma"
	ON_NUM_2_1 MulState = "on_num_2_1"
	ON_NUM_2_2 MulState = "on_num_2_2"
	ON_NUM_2_3 MulState = "on_num_2_3"
	ON_CLOSE_PAR MulState = "on_close_par"
)

func isDigit(c rune) bool {
	if c == '0' { return true }
	if c == '1' { return true }
	if c == '2' { return true }
	if c == '3' { return true }
	if c == '4' { return true }
	if c == '5' { return true }
	if c == '6' { return true }
	if c == '7' { return true }
	if c == '8' { return true }
	if c == '9' { return true }
	return false
}
var lookAheadTable LookAheadTable = LookAheadTable{
	E: { "m": ON_M, "d": ON_D, },
	SCAN: { },
	ON_M: { "u": ON_U, },
	ON_U: { "l": ON_L, },
	ON_L: { "(": ON_MUL_OPEN, },
	ON_MUL_OPEN: { DIGIT: ON_MUL_NUM_1_1, },
	ON_MUL_NUM_1_1: { DIGIT: ON_NUM_1_2, ",": ON_COMMA, },
	ON_NUM_1_2: { DIGIT: ON_NUM_1_3, ",": ON_COMMA, },
	ON_NUM_1_3: { ",": ON_COMMA, },
	ON_COMMA: { DIGIT: ON_NUM_2_1, },
	ON_NUM_2_1: { DIGIT: ON_NUM_2_2, ")": ON_CLOSE_PAR, },
	ON_NUM_2_2: { DIGIT: ON_NUM_2_3, ")": ON_CLOSE_PAR, },
	ON_NUM_2_3: { ")": ON_CLOSE_PAR, },
	ON_CLOSE_PAR: { },
	ON_D: { "o": ON_O, },
	ON_O: { "n": ON_N, "(": ON_DO_OPEN, },
	ON_DO_OPEN: { ")": ON_DO_CLOSE, },
	ON_DO_CLOSE: { },
	ON_N: { "'": ON_APOS, },
	ON_APOS: { "t": ON_T, },
	ON_T: { "(": ON_DONT_OPEN, },
	ON_DONT_OPEN: { ")": ON_DONT_CLOSE, },
	ON_DONT_CLOSE: { },
}

var actionTable ActionTable = ActionTable{
	ON_MUL_OPEN: initBuffers,
	ON_MUL_NUM_1_1: pushLeft,
	ON_NUM_1_2: pushLeft,
	ON_NUM_1_3: pushLeft,
	ON_NUM_2_1: pushRight,
	ON_NUM_2_2: pushRight,
	ON_NUM_2_3: pushRight,
	ON_CLOSE_PAR: closeBuffers,
	ON_DO_CLOSE: enableMult,
	ON_DONT_CLOSE: disableMult,
}

func enableMult(p *Parse, c rune) {
	if p.MulEnabled { return }
	p.MulEnabled = true
	fmt.Printf("\n>>>ENABLE mult\n")
}
func disableMult(p *Parse, c rune) {
	if (!p.MulEnabled) { return }
	p.MulEnabled = false
	fmt.Printf("\n>>>DISABLE mult\n")
}
func initBuffers(p *Parse, c rune) {
	p.LBuffer = ""
	p.RBuffer = ""
}

func pushLeft(p *Parse, c rune) {
	p.LBuffer = p.LBuffer + string(c)
}

func pushRight(p *Parse, c rune) {
	p.RBuffer = p.RBuffer + string(c)
}

func closeBuffers(p *Parse, c rune) {
	if !p.MulEnabled {
		return
	}
	l, err := strconv.ParseInt(p.LBuffer, 10, 64)

	if err != nil { log.Panicln(err); }

	r, err := strconv.ParseInt(p.RBuffer, 10, 64)
	if err != nil { log.Panicln(err); }

	p.MulPair = append(p.MulPair, &MulPair{l:l, r:r})

	p.Total += l * r
	p.__printLongState()
}
func (p *Parse) __printState() {
    fmt.Printf("\n>>>char=%v, %s->%s, I=%d, LB=%s, RB=%s\n", string(p.Char), p.CurState, p.NxtState, p.Iteration, p.LBuffer, p.RBuffer)
}

func (p *Parse) __printLongState() {
    fmt.Printf("\n>>>\tMulPair=%+v\n>>>\tTotal=%v\n", p.MulPair[len(p.MulPair) - 1], p.Total)
}

func (p *Parse)step(next rune) {
	fmt.Printf("%s", string(next))

	p.Char = next
	p.NxtState = SCAN
	if val, ok := p.LookAheadTable[E][string(next)]; ok {
		p.NxtState = val
	} else if val, ok := p.LookAheadTable[p.CurState][string(next)]; ok {
		p.NxtState = val
	} else if val, ok := p.LookAheadTable[p.CurState][DIGIT]; ok && isDigit(next) {
		p.NxtState = val
	}

	if action, ok := p.ActionTable[p.NxtState]; ok {
		action(p, next);
	}

	p.Iteration++
	p.CurState = p.NxtState
	p.NxtState = E
}

func NewParser(l LookAheadTable, a ActionTable) *Parse {
	return &Parse{
		ActionTable: actionTable,
		LookAheadTable: lookAheadTable,
		MulEnabled: true,
		NxtState: SCAN,
		CurState: SCAN,
		LBuffer: "",
		RBuffer: "",
	}
}

func main() {

    file, err := os.Open("./aoc3.data")
    if err != nil { log.Panic(err) }
    defer file.Close()

	parser := NewParser(lookAheadTable, actionTable)
    reader := bufio.NewReader(file)
	for {
        char, _, err := reader.ReadRune()
        if err != nil { break }
		parser.step(char)
    }
}