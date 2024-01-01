package base

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
	"testing"
)

var testIntervals = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string

	//a
	// i_interval_string.Before(x_interval_string)
	i_Before_x bool
	//b
	// x_interval_string.Before(i_interval_string)
	x_Before_i bool
	//c
	// i_interval_string.Cover(x_interval_string)
	i_Cover_x bool
	//d
	// x_interval_string.Cover(i_interval_string)
	x_Cover_i bool

	//e
	// i_interval_string.Intersect(x_interval_string)
	i_Interval_x string
	// x_interval_string.Intersect(i_interval_string)
	// == i_Interval_x
	// f string

	//g,h
	// i_interval_string.Bisect(x_interval_string)
	i_Bisect_x_1, i_Bisect_x_2 string
	//j,k
	// x_interval_string.Bisect(i_interval_string)
	x_Bisect_i_1, x_Bisect_i_2 string

	//l
	// i_interval_string.Adjoin(x_interval_string)
	i_Adjoin_x string
	// x_interval_string.Adjoin(i_interval_string)
	// i_Adjoin_x == m
	// m string

	//o
	// i_interval_string.Encompass(x_interval_string)
	i_Encompass_x string
	// x_interval_string.Encompass(i_interval_string)
	// == i_Encompass_x
	// p string
}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_Before_x:        true,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "=====",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "-------=========",

		i_Adjoin_x: "",

		i_Encompass_x: "================",
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Before_x:        true,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "=====",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "------=========",

		i_Adjoin_x: "",

		i_Encompass_x: "===============",
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Before_x:        true,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "=====",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "-----*=========",

		i_Adjoin_x: "",

		i_Encompass_x: "===============",
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Before_x:        true,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "=====",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "-----=========",

		i_Adjoin_x: "",

		i_Encompass_x: "==============",
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Before_x:        true,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "=====",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*=========",

		i_Adjoin_x: "==============",

		i_Encompass_x: "==============",
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "----=",

		i_Bisect_x_1: "====*",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*========",

		i_Adjoin_x: "=============",

		i_Encompass_x: "=============",
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "--===",

		i_Bisect_x_1: "==*",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*======",

		i_Adjoin_x: "",

		i_Encompass_x: "===========",
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "-====",

		i_Bisect_x_1: "=*",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*=====",

		i_Adjoin_x: "",

		i_Encompass_x: "==========",
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "*====",

		i_Bisect_x_1: "=",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*=====",

		i_Adjoin_x: "",

		i_Encompass_x: "==========",
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "",
		x_Bisect_i_2: "----*====",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "*=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "=",
		x_Bisect_i_2: "-----*===",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "-=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "=*",
		x_Bisect_i_2: "-----*===",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "--=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "==*",
		x_Bisect_i_2: "------*==",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "---=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "===*",
		x_Bisect_i_2: "-------*=",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "---=====*",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "===*",
		x_Bisect_i_2: "--------=",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         true,
		i_Interval_x:      "----=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "",
		x_Bisect_i_1: "====*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "=========",
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "----=====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------**",
		x_Bisect_i_1: "====*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "=========*",
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "-----====",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------*=",
		x_Bisect_i_1: "=====*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "==========",
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "------===",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------*==",
		x_Bisect_i_1: "======*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "===========",
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "-------==",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------*===",
		x_Bisect_i_1: "=======*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "============",
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "--------=",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------*====",
		x_Bisect_i_1: "========*",
		x_Bisect_i_2: "",

		i_Adjoin_x: "=============",

		i_Encompass_x: "=============",
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "--------*=====",
		x_Bisect_i_1: "=========",
		x_Bisect_i_2: "",

		i_Adjoin_x: "==============",

		i_Encompass_x: "==============",
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "---------=====",
		x_Bisect_i_1: "=========",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "==============",
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "---------*=====",
		x_Bisect_i_1: "=========",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "===============",
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "----------=====",
		x_Bisect_i_1: "=========",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "===============",
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
		i_Cover_x:         false,
		x_Cover_i:         false,
		i_Interval_x:      "",

		i_Bisect_x_1: "",
		i_Bisect_x_2: "-----------=====",
		x_Bisect_i_1: "=========",
		x_Bisect_i_2: "",

		i_Adjoin_x: "",

		i_Encompass_x: "================",
	},
}

func parseInterval[T constraints.Integer | constraints.Float](s string) Interval[T] {
	if s == "" {
		return Interval[T]{}
	}
	var begin int
	begin = strings.IndexAny(s, "*=")
	end := strings.LastIndexAny(s, "*=")
	l_lowerUnbounded := false
	l_upperUnbounded := false
	if strings.ContainsRune(s, '<') {
		l_lowerUnbounded = true
	}
	if strings.ContainsRune(s, '>') {
		l_upperUnbounded = true
	}
	return Interval[T]{
		lower:          T(begin),
		lowerIncluded:  s[begin] == '=',
		lowerUnbounded: l_lowerUnbounded,
		upper:          T(end),
		upperIncluded:  s[end] == '=',
		upperUnbounded: l_upperUnbounded,
	}
}

func TestIntInterval(t *testing.T) {
	for n, tc := range testIntervals {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[float64](tc.i_interval_string)
			x := parseInterval[float64](tc.x_interval_string)

			a, b := i.LtBeginOf(x), x.LtBeginOf(i)
			if a != tc.i_Before_x {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", i, x, tc.i_Before_x, a)
			}
			if b != tc.x_Before_i {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", x, i, tc.x_Before_i, b)
			}

			c, d := i.Contains(x), x.Contains(i)
			if c != tc.i_Cover_x {
				t.Errorf("want %s.Cover(%s) = %v but get %v", i, x, tc.i_Cover_x, c)
			}
			if d != tc.x_Cover_i {
				t.Errorf("want %s.Cover(%s) = %v but get %v", x, i, tc.x_Cover_i, d)
			}

			e, f := i.Intersect(x), x.Intersect(i)
			we := parseInterval[float64](tc.i_Interval_x)
			if !e.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", i, x, we, e)
			}
			if !f.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", x, i, we, f)
			}

			g, h := i.Bisect(x)
			wg, wh := parseInterval[float64](tc.i_Bisect_x_1), parseInterval[float64](tc.i_Bisect_x_2)
			if !g.Equal(wg) || !h.Equal(wh) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", i, x, wg, wh, g, h)
			}
			j, k := x.Bisect(i)
			wj, wk := parseInterval[float64](tc.x_Bisect_i_1), parseInterval[float64](tc.x_Bisect_i_2)
			if !j.Equal(wj) || !k.Equal(wk) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", x, i, wj, wk, k, k)
			}

			l, m := i.Adjoin(x), x.Adjoin(i)
			wl := parseInterval[float64](tc.i_Adjoin_x)
			if !l.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", i, x, wl, l)
			}
			if !m.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", x, i, wl, m)
			}

			o, p := i.Encompass(x), x.Encompass(i)
			wo := parseInterval[float64](tc.i_Encompass_x)
			if !o.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", i, x, wo, o)
			}
			if !p.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", x, i, wo, p)
			}
		})
	}
}

func TestFloatInterval(t *testing.T) {
	for n, tc := range testIntervals {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[float64](tc.i_interval_string)
			x := parseInterval[float64](tc.x_interval_string)

			a, b := i.LtBeginOf(x), x.LtBeginOf(i)
			if a != tc.i_Before_x {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", i, x, tc.i_Before_x, a)
			}
			if b != tc.x_Before_i {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", x, i, tc.x_Before_i, b)
			}

			c, d := i.Contains(x), x.Contains(i)
			if c != tc.i_Cover_x {
				t.Errorf("want %s.Cover(%s) = %v but get %v", i, x, tc.i_Cover_x, c)
			}
			if d != tc.x_Cover_i {
				t.Errorf("want %s.Cover(%s) = %v but get %v", x, i, tc.x_Cover_i, d)
			}

			e, f := i.Intersect(x), x.Intersect(i)
			we := parseInterval[float64](tc.i_Interval_x)
			if !e.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", i, x, we, e)
			}
			if !f.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", x, i, we, f)
			}

			g, h := i.Bisect(x)
			wg, wh := parseInterval[float64](tc.i_Bisect_x_1), parseInterval[float64](tc.i_Bisect_x_2)
			if !g.Equal(wg) || !h.Equal(wh) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", i, x, wg, wh, g, h)
			}
			j, k := x.Bisect(i)
			wj, wk := parseInterval[float64](tc.x_Bisect_i_1), parseInterval[float64](tc.x_Bisect_i_2)
			if !j.Equal(wj) || !k.Equal(wk) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", x, i, wj, wk, k, k)
			}

			l, m := i.Adjoin(x), x.Adjoin(i)
			wl := parseInterval[float64](tc.i_Adjoin_x)
			if !l.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", i, x, wl, l)
			}
			if !m.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", x, i, wl, m)
			}

			o, p := i.Encompass(x), x.Encompass(i)
			wo := parseInterval[float64](tc.i_Encompass_x)
			if !o.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", i, x, wo, o)
			}
			if !p.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", x, i, wo, p)
			}
		})
	}
}
