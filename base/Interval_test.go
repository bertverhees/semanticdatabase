package base

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
	"testing"
)

var testsIntervalAdjoin = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string
	//l
	// i_interval_string.Adjoin(x_interval_string)
	i_Adjoin_x string
}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_Adjoin_x:        "",
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_Adjoin_x:        "",
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Adjoin_x:        "",
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Adjoin_x:        "",
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Adjoin_x:        "",
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Adjoin_x:        "==============",
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Adjoin_x:        "=============",
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Adjoin_x:        "",
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Adjoin_x:        "",
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Adjoin_x:        "",
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "=============",
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "==============",
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Adjoin_x:        "",
	},
}

var testsIntervalBisect = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string

	//g,h
	// i_interval_string.Bisect(x_interval_string)
	i_Bisect_x_1, i_Bisect_x_2 string
	//j,k
	// x_interval_string.Bisect(i_interval_string)
	x_Bisect_i_1, x_Bisect_i_2 string
}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "-------=========",
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "-------=========",
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "------=========",
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "-----*=========",
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "-----=========",
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Bisect_x_1:      "=====",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*=========",
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Bisect_x_1:      "====*",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*========",
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Bisect_x_1:      "==*",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*======",
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Bisect_x_1:      "=*",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*=====",
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Bisect_x_1:      "=",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*=====",
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "",
		x_Bisect_i_2:      "----*====",
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "=",
		x_Bisect_i_2:      "-----*===",
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "=*",
		x_Bisect_i_2:      "-----*===",
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "==*",
		x_Bisect_i_2:      "------*==",
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "===*",
		x_Bisect_i_2:      "-------*=",
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "===*",
		x_Bisect_i_2:      "--------=",
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "",
		x_Bisect_i_1:      "====*",
		x_Bisect_i_2:      "",
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------**",
		x_Bisect_i_1:      "====*",
		x_Bisect_i_2:      "",
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------*=",
		x_Bisect_i_1:      "=====*",
		x_Bisect_i_2:      "",
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------*==",
		x_Bisect_i_1:      "======*",
		x_Bisect_i_2:      "",
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------*===",
		x_Bisect_i_1:      "=======*",
		x_Bisect_i_2:      "",
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------*====",
		x_Bisect_i_1:      "========*",
		x_Bisect_i_2:      "",
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "--------*=====",
		x_Bisect_i_1:      "=========",
		x_Bisect_i_2:      "",
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "---------=====",
		x_Bisect_i_1:      "=========",
		x_Bisect_i_2:      "",
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "---------*=====",
		x_Bisect_i_1:      "=========",
		x_Bisect_i_2:      "",
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "----------=====",
		x_Bisect_i_1:      "=========",
		x_Bisect_i_2:      "",
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Bisect_x_1:      "",
		i_Bisect_x_2:      "-----------=====",
		x_Bisect_i_1:      "=========",
		x_Bisect_i_2:      "",
	},
}

var testsIntervalContains = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string

	//c
	// i_interval_string.Cover(x_interval_string)
	i_Cover_x bool
	//d
	// x_interval_string.Cover(i_interval_string)
	x_Cover_i bool
}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
}

var testsIntervalLtBeginOf = []struct {
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
}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Before_x:        false,
		x_Before_i:        true,
	},
}

var testsIntervalIntersect = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string

	//e
	// i_interval_string.Intersect(x_interval_string)
	i_intersect_x string
	// x_interval_string.Intersect(i_interval_string)
	// == i_intersect_x
	// f string

}{
	{ // 0
		i_interval_string: "=====",
		x_interval_string: "-------=========",
		i_intersect_x:     "",
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_intersect_x:     "-------=========",
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_intersect_x:     "",
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_intersect_x:     "",
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_intersect_x:     "",
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_intersect_x:     "",
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_intersect_x:     "----=",
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_intersect_x:     "--===",
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_intersect_x:     "-====",
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_intersect_x:     "*====",
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_intersect_x:     "=====",
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_intersect_x:     "*=====",
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_intersect_x:     "-=====",
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_intersect_x:     "--=====",
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_intersect_x:     "---=====",
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_intersect_x:     "---=====*",
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_intersect_x:     "----=====",
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_intersect_x:     "----=====",
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_intersect_x:     "-----====",
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_intersect_x:     "------===",
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_intersect_x:     "-------==",
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_intersect_x:     "--------=",
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_intersect_x:     "",
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_intersect_x:     "",
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_intersect_x:     "",
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_intersect_x:     "",
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_intersect_x:     "",
	},
}

var testsIntervalEncompass = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string
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
		i_Encompass_x:     "================",
	},
	{ // 0a
		i_interval_string: "=====>",
		x_interval_string: "-------=========",
		i_Encompass_x:     "================",
	},
	{ // 1
		i_interval_string: "=====",
		x_interval_string: "------=========",
		i_Encompass_x:     "===============",
	},
	{ // 2
		i_interval_string: "=====",
		x_interval_string: "-----*=========",
		i_Encompass_x:     "===============",
	},
	{ // 3
		i_interval_string: "=====",
		x_interval_string: "-----=========",
		i_Encompass_x:     "==============",
	},
	{ // 4
		i_interval_string: "=====",
		x_interval_string: "----*=========",
		i_Encompass_x:     "==============",
	},
	{ // 5
		i_interval_string: "=====",
		x_interval_string: "----=========",
		i_Encompass_x:     "=============",
	},
	{ // 6
		i_interval_string: "=====",
		x_interval_string: "--=========",
		i_Encompass_x:     "===========",
	},
	{ // 7
		i_interval_string: "=====",
		x_interval_string: "-=========",
		i_Encompass_x:     "==========",
	},
	{ // 8
		i_interval_string: "=====",
		x_interval_string: "*=========",
		i_Encompass_x:     "==========",
	},
	{ // 9
		i_interval_string: "=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 10
		i_interval_string: "*=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 11
		i_interval_string: "-=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 12
		i_interval_string: "--=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 13
		i_interval_string: "---=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 14
		i_interval_string: "---=====*",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 15
		i_interval_string: "----=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=========",
	},
	{ // 16
		i_interval_string: "----=====*",
		x_interval_string: "=========",
		i_Encompass_x:     "=========*",
	},
	{ // 17
		i_interval_string: "-----=====",
		x_interval_string: "=========",
		i_Encompass_x:     "==========",
	},
	{ // 18
		i_interval_string: "------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "===========",
	},
	{ // 19
		i_interval_string: "-------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "============",
	},
	{ // 20
		i_interval_string: "--------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "=============",
	},
	{ // 21
		i_interval_string: "--------*=====",
		x_interval_string: "=========",
		i_Encompass_x:     "==============",
	},
	{ // 22
		i_interval_string: "---------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "==============",
	},
	{ // 23
		i_interval_string: "---------*=====",
		x_interval_string: "=========",
		i_Encompass_x:     "===============",
	},
	{ // 24
		i_interval_string: "----------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "===============",
	},
	{ // 25
		i_interval_string: "-----------=====",
		x_interval_string: "=========",
		i_Encompass_x:     "================",
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
	l_lowerIncluded := s[begin] == '='
	l_upperIncluded := s[end] == '='
	if strings.ContainsRune(s, '<') {
		l_lowerUnbounded = true
		//begin = 0
		//l_lowerIncluded = false
	}
	if strings.ContainsRune(s, '>') {
		l_upperUnbounded = true
		//end = 0
		//l_upperIncluded = false
	}
	return Interval[T]{
		lower:          T(begin),
		lowerIncluded:  l_lowerIncluded,
		lowerUnbounded: l_lowerUnbounded,
		upper:          T(end),
		upperIncluded:  l_upperIncluded,
		upperUnbounded: l_upperUnbounded,
	}
}

func TestIntervalBisect(t *testing.T) {
	testIntervalBisect[int](t)
	testIntervalBisect[float64](t)
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

func testIntervalBisect[T constraints.Integer | constraints.Float](t *testing.T) {
	for n, tc := range testsIntervalBisect {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			i := parseInterval[T](tc.i_interval_string)
			x := parseInterval[T](tc.x_interval_string)

			g, h := i.Bisect(x)
			wg, wh := parseInterval[T](tc.i_Bisect_x_1), parseInterval[T](tc.i_Bisect_x_2)
			if !g.Equal(wg) || !h.Equal(wh) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", i, x, wg, wh, g, h)
			}
			j, k := x.Bisect(i)
			wj, wk := parseInterval[T](tc.x_Bisect_i_1), parseInterval[T](tc.x_Bisect_i_2)
			if !j.Equal(wj) || !k.Equal(wk) {
				t.Errorf("want %s.Bisect(%s) = %s, %s but get %s, %s", x, i, wj, wk, k, k)
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
