package base

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
	"testing"
)

/*
*
Reading the test-sets
The '========' stands for interval-range
Upper- en lower-Included is true
The '<' is for lowerUnbound, no lower-end
The '>' is for upperUnbound, no upper-end
The '*' is for Included is false (depending on the site for upper en lower included)
The * and < and > do not count!!!!!!

It looks hard to understand and is maybe confusing, but read on,
So:
*=================  lowerIncluded is false, and lower  is 0
=================   lowerIncluded is true, and lower is 0
<================ lowerUnbounded is true and lower is 0
<*============== lowerUnbounded is true, and lowerIncluded is false and lower is 0
---*=================  lowerIncluded is false, and lower  is 4
---=================   lowerIncluded is true, and lower is 4
---<================ lowerUnbounded is true and lower is 4
---<*============== lowerUnbounded is true, and lowerIncluded is false and lower is 4

# On the uppersite are mutatis mutandis the same rules

So, these intervals, although visible giving another picture, do NOT overlap
<*====-------------- 0,4
----*=======*       5,12

Remember, this is only for creating tests-sets, it has nothing to do with a notation language of meaning outside these test-sets
*/
func parseInterval[T constraints.Integer | constraints.Float](s string) Interval[T] {
	if s == "" {
		return Interval[T]{}
	}
	lLowerunbounded := false
	lUpperunbounded := false
	lLowerincluded := false
	lUpperincluded := false
	var begin int
	var end int
	if strings.Index(s, "<*=") >= 0 {
		lLowerunbounded = true
		begin = strings.Index(s, "=") - 2
		end = strings.LastIndex(s, "=") - 1
	} else if strings.Index(s, "<=") >= 0 {
		lLowerunbounded = true
		lLowerincluded = true
		begin = strings.Index(s, "=") - 1
		end = strings.LastIndex(s, "=")
	} else if strings.Index(s, "*=") >= 0 {
		begin = strings.Index(s, "=") - 1
		end = strings.LastIndex(s, "=")
	} else if strings.Index(s, "=") >= 0 {
		lLowerincluded = true
		begin = strings.Index(s, "=")
		end = strings.LastIndex(s, "=") + 1
	}
	if strings.LastIndex(s, "=*>") == strings.LastIndex(s, "=") {
		lUpperunbounded = true
	} else if strings.LastIndex(s, "=>") == strings.LastIndex(s, "=") {
		lUpperunbounded = true
		lUpperincluded = true
	} else if strings.LastIndex(s, "=*") == strings.LastIndex(s, "=") {
		lUpperincluded = false
	} else {
		lUpperincluded = true
	}

	return Interval[T]{
		lower:          T(begin),
		lowerIncluded:  lLowerincluded,
		lowerUnbounded: lLowerunbounded,
		upper:          T(end),
		upperIncluded:  lUpperincluded,
		upperUnbounded: lUpperunbounded,
	}
}

func TestParseInterval(t *testing.T) {
	testParseInterval[int](t)
	testParseInterval[float64](t)
}

func TestIntervalSubtract(t *testing.T) {
	testIntervalSubtract[int](t)
	testIntervalSubtract[float64](t)
}

func TestIntervalIntersect(t *testing.T) {
	testIntervalIntersect[int](t)
	testIntervalIntersect[float64](t)
}

func TestIntervalLtBeginOf(t *testing.T) {
	testIntervalLtBeginOf[int](t)
	testIntervalLtBeginOf[float64](t)
}

func TestIntervalContains(t *testing.T) {
	testIntervalContains[int](t)
	testIntervalContains[float64](t)
}

func TestIntervalAdjoin(t *testing.T) {
	testIntervalAdjoin[int](t)
	testIntervalAdjoin[float64](t)
}

func testParseInterval[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsParseInterval {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.s)
			if i.lower != T(tc.begin) ||
				i.upper != T(tc.end) ||
				i.lowerIncluded != tc.lowerIncluded ||
				i.lowerUnbounded != tc.lowerUnbounded ||
				i.upperIncluded != tc.upperIncluded ||
				i.upperUnbounded != tc.upperUnbounded {
				t.Errorf("String s: %s want tc = %v but get %v", tc.s, tc, i)
			}
		})
	}
}

func TestIntervalEncompass(t *testing.T) {
	testIntervalEncompass[int](t)
	testIntervalEncompass[float64](t)
}

func testIntervalLtBeginOf[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalLtBeginOf {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			a, b := i.LtBeginOf(x), x.LtBeginOf(i)
			if a != tc.i_Before_x {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", i, x, tc.i_Before_x, a)
			}
			if b != tc.x_Before_i {
				t.Errorf("want %s.LtBeginOf(%s) = %v but get %v", x, i, tc.x_Before_i, b)
			}
		})
	}
}

func testIntervalContains[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalContains {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			c, d := i.Contains(x), x.Contains(i)
			if c != tc.i_Cover_x {
				t.Errorf("want %s.Cover(%s) = %v but get %v", i, x, tc.i_Cover_x, c)
			}
			if d != tc.x_Cover_i {
				t.Errorf("want %s.Cover(%s) = %v but get %v", x, i, tc.x_Cover_i, d)
			}
		})
	}
}

func testIntervalIntersect[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalIntersect {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			e, f := i.Intersect(x), x.Intersect(i)
			we := parseInterval[T](tc.i_intersect_x)
			if !e.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", i, x, we, e)
			}
			if !f.Equal(we) {
				t.Errorf("want %s.Intersect(%s) = %s but get %s", x, i, we, f)
			}
		})
	}
}

func testIntervalAdjoin[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalAdjoin {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			l, m := i.Adjoin(x), x.Adjoin(i)
			wl := parseInterval[T](tc.i_Adjoin_x)
			if !l.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", i, x, wl, l)
			}
			if !m.Equal(wl) {
				t.Errorf("want %s.Adjoin(%s) = %s but get %s", x, i, wl, m)
			}
		})
	}
}

func testIntervalSubtract[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalSubtract {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			g, h := i.Subtract(x)
			wg, wh := parseInterval[T](tc.i_Subtract_x_before), parseInterval[T](tc.i_Subtract_x_after)
			if !g.Equal(wg) || !h.Equal(wh) {
				t.Errorf("want %s.Subtract(%s) = %s, %s but get %s, %s", i, x, wg, wh, g, h)
			}
			j, k := x.Subtract(i)
			wj, wk := parseInterval[T](tc.x_Subtract_i_before), parseInterval[T](tc.x_Subtract_i_after)
			if !j.Equal(wj) || !k.Equal(wk) {
				t.Errorf("want %s.Subtract(%s) = %s, %s but get %s, %s", x, i, wj, wk, k, k)
			}
		})
	}
}

func testIntervalEncompass[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalEncompass {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			o, p := i.Encompass(x), x.Encompass(i)
			wo := parseInterval[T](tc.i_Encompass_x)
			if !o.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", i, x, wo, o)
			}
			if !p.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", x, i, wo, p)
			}
		})
	}
}
