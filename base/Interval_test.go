package base

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
	"testing"
)

/*
Reading the test-sets
The '========' stands for interval-range
Upper- en lower-Included is true
The '<' is for lowerUnbound, no lower-end
The '>' is for upperUnbound, no upper-end
The '*' is for Included is false (depending on the side of the string for upper en lower included)

The interval-string is divided in three parts by |,(or is empty)
First part indicates the lower-unbounded/included, the third part for the right side
It is allowed to have as many characters as is convenient on the first and the last part,
only the indicated characters are meaningful

It looks hard to understand and is maybe confusing, but read on,
So:
*|=================|  lowerIncluded is false, and lower  is 0
|=================|   lowerIncluded is true, and lower is 0
<|================| lowerUnbounded is true and lower is 0
<*|==============| lowerUnbounded is true, and lowerIncluded is false and lower is 0
*|---=================|  lowerIncluded is false, and lower  is 4
|---=================|   lowerIncluded is true, and lower is 4
|---<================| lowerUnbounded is true and lower is 4
<*|---==============| lowerUnbounded is true, and lowerIncluded is false and lower is 4

# On the upperside are mutatis mutandis the same rules

So, these two intervals below, although visible overlapping, do NOT overlap
"<*|====--------------|" 0,4
"*|----=======|*"       5,12
So it can be written like this, which is much more readable
"<*|====--------------| 0,4"
"* |----=======|*"       5,12 or
" *|----=======|*"

Remember, this is only for creating tests-sets, it has nothing to do with a notation language of meaning outside these test-sets
*/
func parseInterval[T constraints.Integer | constraints.Float](s string) (Interval[T], error) {
	if s == "" {
		return Interval[T]{}, nil
	}
	parts := strings.Split(s, "|")
	if len(parts) != 3 {
		return Interval[T]{}, errors.New(fmt.Sprintf("The interval string '%s' is not wellformed, it must have 2 '|' (pipes).", s))
	}
	leftside := parts[0]
	rightside := parts[2]
	interval := parts[1]
	if strings.ContainsAny(interval, "*<>") {
		return Interval[T]{}, errors.New(fmt.Sprintf("The interval string '%s' is not wellformed, it has not allowed characters in the middlepart", s))
	}
	begin := strings.Index(interval, "=")
	end := strings.LastIndex(interval, "=") + 1
	lowerunbounded, lowerincluded, upperunbounded, upperincluded := false, true, false, true
	if len(leftside) > 0 {
		lowerunbounded = strings.Contains(leftside, "<")
		lowerincluded = !strings.Contains(leftside, "*")
	}
	if len(rightside) > 0 {
		upperunbounded = strings.Contains(rightside, ">")
		upperincluded = !strings.Contains(rightside, "*")
	}
	return Interval[T]{
		lower:          T(begin),
		lowerIncluded:  lowerincluded,
		lowerUnbounded: lowerunbounded,
		upper:          T(end),
		upperIncluded:  upperincluded,
		upperUnbounded: upperunbounded,
	}, nil
}

func TestParseInterval(t *testing.T) {
	testParseInterval[int](t)
	testParseInterval[float64](t)
}

func TestIntervalLtBeginOf(t *testing.T) {
	testIntervalLtBeginOf[int](t)
	testIntervalLtBeginOf[float64](t)
}

func TestIntervalLeEndOf(t *testing.T) {
	testIntervalLeEndOf[int](t)
	testIntervalLeEndOf[float64](t)
}

func TestIntervalContains(t *testing.T) {
	testIntervalContains[int](t)
	testIntervalContains[float64](t)
}

func TestIntervalIntersect(t *testing.T) {
	testIntervalIntersect[int](t)
	testIntervalIntersect[float64](t)
}

func TestIntervalSubtract(t *testing.T) {
	testIntervalSubtract[int](t)
	testIntervalSubtract[float64](t)
}

func TestIntervalAdjoin(t *testing.T) {
	testIntervalAdjoin[int](t)
	testIntervalAdjoin[float64](t)
}

func TestIntervalEncompass(t *testing.T) {
	testIntervalEncompass[int](t)
	testIntervalEncompass[float64](t)
}

func testParseInterval[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsParseInterval {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.s)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
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

func testIntervalLtBeginOf[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalLtBeginOf {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.test.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.test.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
				return
			}

			a, b := i.LtBeginOf(x), x.LtBeginOf(i)
			if a != tc.i_Before_x {
				t.Errorf("want %s.LtBeginOf(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", i, x, tc.i_Before_x, a, tc.test.counter, tc.test.i_interval_string, tc.test.x_interval_string)
			}
			if b != tc.x_Before_i {
				t.Errorf("want %s.LtBeginOf(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", x, i, tc.x_Before_i, b, tc.test.counter, tc.test.x_interval_string, tc.test.i_interval_string)
			}
		})
	}
}

func testIntervalLeEndOf[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalLeEndOf {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.test.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.test.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}

			a, b := i.LeEndOf(x), x.LeEndOf(i)
			if a != tc.i_LeEnd_x {
				t.Errorf("want %s.LeEndOf(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", i, x, tc.i_LeEnd_x, a, tc.test.counter, tc.test.i_interval_string, tc.test.x_interval_string)
			}
			if b != tc.x_LeEnd_i {
				t.Errorf("want %s.LeEndOf(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", x, i, tc.x_LeEnd_i, b, tc.test.counter, tc.test.x_interval_string, tc.test.i_interval_string)
			}
		})
	}
}

func testIntervalContains[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalContains {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.test.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.test.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}

			c, d := i.Contains(x), x.Contains(i)
			if c != tc.i_Cover_x {
				t.Errorf("want %s.Contains(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", i, x, tc.i_Cover_x, c, tc.test.counter, tc.test.i_interval_string, tc.test.x_interval_string)
			}
			if d != tc.x_Cover_i {
				t.Errorf("want %s.Contains(%s) = %v (in test) but is %v, counter: %v\n%s\n%s", x, i, tc.x_Cover_i, d, tc.test.counter, tc.test.x_interval_string, tc.test.i_interval_string)
			}

		})
	}
}

func testIntervalIntersect[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalIntersect {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}

			e, f := i.Intersect(x), x.Intersect(i)
			we, er := parseInterval[T](tc.i_intersect_x)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
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
			i, er := parseInterval[T](tc.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}

			l, m := i.Adjoin(x), x.Adjoin(i)
			wl, er := parseInterval[T](tc.i_Adjoin_x)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
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
			i, er := parseInterval[T](tc.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			g, h := i.Subtract(x)
			wg, er := parseInterval[T](tc.i_Subtract_x_before)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			wh, er := parseInterval[T](tc.i_Subtract_x_after)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			if !g.Equal(wg) || !h.Equal(wh) {
				t.Errorf("want %s.Subtract(%s) = %s, %s but get %s, %s", i, x, wg, wh, g, h)
			}
			j, k := x.Subtract(i)
			wj, er := parseInterval[T](tc.x_Subtract_i_before)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			wk, er := parseInterval[T](tc.x_Subtract_i_before)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			if !j.Equal(wj) || !k.Equal(wk) {
				t.Errorf("want %s.Subtract(%s) = %s, %s but get %s, %s", x, i, wj, wk, k, k)
			}
		})
	}
}

func testIntervalEncompass[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalEncompass {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i, er := parseInterval[T](tc.i_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			x, er := parseInterval[T](tc.x_interval_string)
			if er != nil {
				t.Errorf(er.Error())
				return
			}

			o, p := i.Encompass(x), x.Encompass(i)
			wo, er := parseInterval[T](tc.i_Encompass_x)
			if er != nil {
				t.Errorf(er.Error())
				return
			}
			if !o.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", i, x, wo, o)
			}
			if !p.Equal(wo) {
				t.Errorf("want %s.Encompass(%s) = %s but get %s", x, i, wo, p)
			}
		})
	}
}
