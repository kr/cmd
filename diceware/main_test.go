package main

import "testing"

func TestB6digits(t *testing.T) {
	cases := []struct {
		n, w int
	}{
		{1, 1},
		{5, 1},
		{6, 2},
		{7775, 5},
		{7776, 6},
	}
	for _, test := range cases {
		g := b6digits(test.n)
		if g != test.w {
			t.Errorf("b6digits(%d) = %d want %d", test.n, g, test.w)
		}
	}
}

func TestDicenum(t *testing.T) {
	cases := []struct {
		n, b int
		w    string
	}{
		{0, 1, "1"},
		{5, 1, "6"},
		{6, 2, "21"},
		{7775, 5, "66666"},
		{7776, 6, "211111"},
	}
	for _, test := range cases {
		g := dicenum(test.n, test.b)
		if g != test.w {
			t.Errorf("dicenum(%d, %d) = %s want %s", test.n, test.b, g, test.w)
		}
	}
}
