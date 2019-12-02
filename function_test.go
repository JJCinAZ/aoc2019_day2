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
			} else if err != nil {
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
