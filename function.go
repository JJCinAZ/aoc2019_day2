// function.go
package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Part1(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Missing input", http.StatusNotAcceptable)
		return
	}
	pgm := parseInput(string(input))
	if len(pgm) > 3 {
		pgm[1] = 12
		pgm[2] = 2
		execPgm(pgm)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(strconv.FormatInt(int64(pgm[0]), 10)))
}

func Part2(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Missing input", http.StatusNotAcceptable)
		return
	}
	pgm := parseInput(string(input))
	result := findSolution(pgm)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("result=%d", result)))
}

func findSolution(pgm []int) int {
	noun := 0
	verb := 0
	pgmCopy := make([]int, len(pgm))
	OUTER:
	for ; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			copy(pgmCopy, pgm)
			pgmCopy[1] = noun
			pgmCopy[2] = verb
			execPgm(pgmCopy)
			if pgmCopy[0] == 19690720 {
				break OUTER
			}
		}
	}
	return noun*100+verb
}
func parseInput(input string) []int {
	a := strings.Split(input, ",")
	pgm := make([]int, len(a))
	for i := range a {
		pgm[i], _ = strconv.Atoi(a[i])
	}
	return pgm
}

func execPgm(pgm []int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	ip := 0
PGMLOOP:
	for {
		switch pgm[ip] {
		case 99:
			break PGMLOOP
		case 1:
			op1 := pgm[ip+1]
			op2 := pgm[ip+2]
			op3 := pgm[ip+3]
			pgm[op3] = pgm[op1] + pgm[op2]
			ip += 4
		case 2:
			op1 := pgm[ip+1]
			op2 := pgm[ip+2]
			op3 := pgm[ip+3]
			pgm[op3] = pgm[op1] * pgm[op2]
			ip += 4
		default:
			panic(fmt.Errorf("illegal opcode at offset %d", ip))
		}
	}
	return nil
}
