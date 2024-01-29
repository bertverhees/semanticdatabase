package identifiers

import (
	"fmt"
	"semanticdatabase/foundation/primitives"
	"testing"
)

func TestIsEqual(t *testing.T) {
	for n, tc := range testsIsEqual {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			b1 := primitives.NewBoolean(tc.test.b1)
			b2 := primitives.NewBoolean(tc.test.b2)
			result := b1.IsEqual(b2)
			if tc.result != result.Value() {
				t.Errorf("\nwant %v.IsEqual(%v) = %v (result conform test)\n but is actually %v, counter: %v",
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
