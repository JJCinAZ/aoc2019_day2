package function

import (
	"testing"
)

func Test_execPgm(t *testing.T) {
	tests := []struct {
		name    string
		pgm []int
		want []int
		wantErr bool
	}{
		{ "test1", []int{1,0,0,0,99}, []int{2,0,0,0,99}, false},
		{ "test2", []int{2,3,0,3,99}, []int{2,3,0,6,99}, false},
		{ "test3", []int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}, false},
		{ "test4", []int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}, false},
		{ "test5", []int{1,10,0,0,99}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execPgm(tt.pgm); (err != nil) != tt.wantErr {
				t.Errorf("execPgm() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil {
				if !slicesEqual(tt.pgm, tt.want) {
					t.Errorf("execPgm() didn't return expected values")
				}
			}
		})
	}
}

func slicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func Test_findSolution(t *testing.T) {
	input := `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,6,19,23,2,6,23,27,1,5,27,31,2,31,9,35,1,35,5,39,1,39,5,43,1,43,10,47,2,6,47,51,1,51,5,55,2,55,6,59,1,5,59,63,2,63,6,67,1,5,67,71,1,71,6,75,2,75,10,79,1,79,5,83,2,83,6,87,1,87,5,91,2,9,91,95,1,95,6,99,2,9,99,103,2,9,103,107,1,5,107,111,1,111,5,115,1,115,13,119,1,13,119,123,2,6,123,127,1,5,127,131,1,9,131,135,1,135,9,139,2,139,6,143,1,143,5,147,2,147,6,151,1,5,151,155,2,6,155,159,1,159,2,163,1,9,163,0,99,2,0,14,0`
	pgm := parseInput(input)
	result := findSolution(pgm)
	if result == 10100 {
		t.Errorf("findSolution() failed")
	}
}