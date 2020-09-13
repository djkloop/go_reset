package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"111222333", 2},
		{"pwwkewa", 4},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"一二三一二三三二一", 3},
		{"黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞；灰化灰黑化肥会挥发发黑灰为讳飞花回化为灰", 9},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s: "+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞灰化"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s: "+"expected %d", actual, s, ans)
		}
	}

}
