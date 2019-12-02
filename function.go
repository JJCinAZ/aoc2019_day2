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
	noun := 0
	verb := 0
	for ; noun <= 99; noun++ {
		for ; verb <= 99; verb++ {
			pgmCopy := append([]int(nil), pgm...)
			pgmCopy[1] = noun
			pgmCopy[2] = verb
			execPgm(pgmCopy)
			if pgmCopy[0] == 19690720 {
				break
			}
		}
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("noun=%d, verb=%d, result=%d",
		noun, verb, 100*noun+verb)))
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
