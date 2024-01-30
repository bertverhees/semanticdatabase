package primitives

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {
	for n, tc := range testsIsEqual {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			b1 := NewBoolean(tc.test.b1)
			b2 := NewBoolean(tc.test.b2)
			result := b1.IsEqual(b2)
			if tc.result != result.Value() {
				t.Errorf("\nwant %v.IsEqual(%v) = %v (result conform test)\n but is actually %v, counter: %v",
					b1, b2, result, tc.result, tc.test.counter)
				return
			}
		})
	}
}

func TestAndThen(t *testing.T) {
	for n, tc := range testsAndThen {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			b1 := NewBoolean(tc.test.b1)
			b2 := NewBoolean(tc.test.b2)
			result := b1.AndThen(b2)
			if tc.result != result.Value() {
				t.Errorf("\nwant %v.AndThen(%v) = %v (result conform test)\n but is actually %v, counter: %v",
					b1, b2, result, tc.result, tc.test.counter)
				return
			}
		})
	}
}

func TestOrElse(t *testing.T) {
	for n, tc := range testsOrElse {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			b1 := NewBoolean(tc.test.b1)
			b2 := NewBoolean(tc.test.b2)
			result := b1.OrElse(b2)
			if tc.result != result.Value() {
				t.Errorf("\nwant %v.OrElse(%v) = %v (result conform test)\n but is actually %v, counter: %v",
					b1, b2, result, tc.result, tc.test.counter)
				return
			}
		})
	}
}

type testGeneral struct {
	b1      bool
	b2      bool
	counter string
}

var testsBooleans = []testGeneral{
	{
		b1:      true,
		b2:      true,
		counter: "0",
	},
	{
		b1:      true,
		b2:      false,
		counter: "1",
	},
	{
		b1:      false,
		b2:      true,
		counter: "2",
	},
	{
		b1:      false,
		b2:      false,
		counter: "3",
	},
}

var testsIsEqual = []struct {
	test   testGeneral
	result bool
}{
	{
		test:   testsBooleans[0],
		result: true,
	},
	{
		test:   testsBooleans[1],
		result: false,
	},
	{
		test:   testsBooleans[2],
		result: false,
	},
	{
		test:   testsBooleans[3],
		result: true,
	},
}

var testsAndThen = []struct {
	test   testGeneral
	result bool
}{
	{
		test:   testsBooleans[0],
		result: true,
	},
	{
		test:   testsBooleans[1],
		result: false,
	},
	{
		test:   testsBooleans[2],
		result: false,
	},
	{
		test:   testsBooleans[3],
		result: false,
	},
}

var testsOrElse = []struct {
	test   testGeneral
	result bool
}{
	{
		test:   testsBooleans[0],
		result: true,
	},
	{
		test:   testsBooleans[1],
		result: true,
	},
	{
		test:   testsBooleans[2],
		result: true,
	},
	{
		test:   testsBooleans[3],
		result: false,
	},
}
