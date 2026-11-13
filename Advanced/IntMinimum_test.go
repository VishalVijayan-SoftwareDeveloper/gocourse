package Advanced

import (
	"fmt"
	"testing"
)

/*
func TestIntmin(t *testing.T) {
	ans := Intmin(1, 2)
	if ans != 1 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}

}*/

func TestIntminTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Intmin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
