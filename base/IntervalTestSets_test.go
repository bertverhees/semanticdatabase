package base

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
<*|====--------------| 0,4
*|----=======|*       5,12
So it can be written like this, which is much more readable
So it can be written like this, which is much more readable
"<*|====--------------| 0,4"
"* |----=======|*"       5,12 or
" *|----=======|*"
Remember, this is only for creating tests-sets, it has nothing to do with a notation language of meaning outside these test-sets
*/

var testsParseInterval = []struct {
	s              string
	begin          int
	lowerIncluded  bool
	lowerUnbounded bool
	end            int
	upperIncluded  bool
	upperUnbounded bool
}{
	{
		s:              "  |=====|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              " *|=====|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "  |=====|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              " *|=====|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              " <|=====|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              " |=====|>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              " <|=====|>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "  |=====|*>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              "<*|=====|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "<*|=====|*>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	//-------------
	{
		s:              "  |--=====|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              " *|--=====|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	{
		s:              " *|--=====|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              " <|--=====|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====|>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              " <|--=====|>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====|*>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	{
		s:              "<*|--=====|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "<*|--=====|*>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	//-------------
	{
		s:              "  |--=====--|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              " *|--=====--|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====--|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	{
		s:              " *|--=====--|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	{
		s:              " <|--=====--|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====--|>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              " <|--=====--|>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "  |--=====--|*>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	{
		s:              "<*|--=====--|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          2,
		end:            7,
	},
	{
		s:              "<*|--=====--|*>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          2,
		end:            7,
	},
	//-------------
	{
		s:              "  |=====--|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              " *|=====--|",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "  |=====--|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              " *|=====--|*",
		lowerUnbounded: false,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              " <|=====--|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "  |=====--|>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              " <|=====--|>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "  |=====--|*>",
		lowerUnbounded: false,
		upperUnbounded: true,
		lowerIncluded:  true,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
	{
		s:              "<*|=====--|",
		lowerUnbounded: true,
		upperUnbounded: false,
		lowerIncluded:  false,
		upperIncluded:  true,
		begin:          0,
		end:            5,
	},
	{
		s:              "<*|=====--|*>",
		lowerUnbounded: true,
		upperUnbounded: true,
		lowerIncluded:  false,
		upperIncluded:  false,
		begin:          0,
		end:            5,
	},
}

var testsIntervalSubtract = []struct {
	//i
	i_interval_string string
	//x
	x_interval_string string

	//g,h
	// i_interval_string.Subtract(x_interval_string)
	i_Subtract_x_before, i_Subtract_x_after string
	//j,k
	// x_interval_string.Subtract(i_interval_string)
	x_Subtract_i_before, x_Subtract_i_after string
}{
	{ // 0
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |-------=========|",
		i_Subtract_x_before: "  |=====|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  "  |-------=========|",
	},
	{ // 0a
		i_interval_string:   "  |=====|>",
		x_interval_string:   "  |-------=========|",
		i_Subtract_x_before: "  |=======|",
		i_Subtract_x_after:  "  |----------------|>",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  "",
	},
	{ // 1
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |------=========|",
		i_Subtract_x_before: "  |=====|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  "  |------=========|",
	},
	{ // 2
		i_interval_string:   "  |=====|",
		x_interval_string:   " *|-----=========|",
		i_Subtract_x_before: "  |=====|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|-----=========|",
	},
	{ // 3
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |-----=========|",
		i_Subtract_x_before: "  |=====|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  "  |-----=========|",
	},
	{ // 4
		i_interval_string:   "  |=====|",
		x_interval_string:   " *|----=========|",
		i_Subtract_x_before: "  |=====|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----=========|",
	},
	{ // 5
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |----=========|",
		i_Subtract_x_before: "  |====|*",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----========|",
	},
	{ // 6
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |--=========|",
		i_Subtract_x_before: "  |==|*",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----======|",
	},
	{ // 7
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |-=========|",
		i_Subtract_x_before: "  |=|*",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----=====|",
	},
	{ // 8
		i_interval_string:   "  |=====|",
		x_interval_string:   " *|=========|",
		i_Subtract_x_before: "  |=|",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----=====|",
	},
	{ // 9
		i_interval_string:   "  |=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "",
		x_Subtract_i_after:  " *|----====|",
	},
	{ // 10
		i_interval_string:   " *|=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |=|",
		x_Subtract_i_after:  " *|-----===|",
	},
	{ // 11
		i_interval_string:   "  |-=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |=|*",
		x_Subtract_i_after:  " *|-----===|",
	},
	{ // 12
		i_interval_string:   "  |--=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |==|*",
		x_Subtract_i_after:  "  |------*==|",
	},
	{ // 13
		i_interval_string:   "  |---=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |===|*",
		x_Subtract_i_after:  "  |-------*=|",
	},
	{ // 14
		i_interval_string:   "  |---=====|*",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |===|*",
		x_Subtract_i_after:  "  |--------=|",
	},
	{ // 15
		i_interval_string:   "  |----=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "",
		x_Subtract_i_before: "  |====|*",
		x_Subtract_i_after:  "",
	},
	{ // 16
		i_interval_string:   "  |----=====|*",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "**|--------",
		x_Subtract_i_before: "  |====|*",
		x_Subtract_i_after:  "",
	},
	{ // 17
		i_interval_string:   "  |-----=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|--------=|",
		x_Subtract_i_before: "  |=====|*",
		x_Subtract_i_after:  "",
	},
	{ // 18
		i_interval_string:   "  |------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|--------==|",
		x_Subtract_i_before: "  |======|*",
		x_Subtract_i_after:  "",
	},
	{ // 19
		i_interval_string:   "  |-------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|--------===|",
		x_Subtract_i_before: "  |=======|*",
		x_Subtract_i_after:  "",
	},
	{ // 20
		i_interval_string:   "  |--------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|--------====|",
		x_Subtract_i_before: "  |========|*",
		x_Subtract_i_after:  "",
	},
	{ // 21
		i_interval_string:   " *|--------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|--------=====|",
		x_Subtract_i_before: "  |=========|",
		x_Subtract_i_after:  "",
	},
	{ // 22
		i_interval_string:   "  |---------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "  |---------=====|",
		x_Subtract_i_before: "  |=========|",
		x_Subtract_i_after:  "",
	},
	{ // 23
		i_interval_string:   " *|---------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  " *|---------=====|",
		x_Subtract_i_before: "  |=========|",
		x_Subtract_i_after:  "",
	},
	{ // 24
		i_interval_string:   "  |----------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "  |----------=====|",
		x_Subtract_i_before: "  |=========|",
		x_Subtract_i_after:  "",
	},
	{ // 25
		i_interval_string:   "  |-----------=====|",
		x_interval_string:   "  |=========|",
		i_Subtract_x_before: "",
		i_Subtract_x_after:  "  |-----------=====|",
		x_Subtract_i_before: "  |=========|",
		x_Subtract_i_after:  "",
	},
}

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
		i_interval_string: "  |=====|",
		x_interval_string: "  |-------=========|",
		i_Adjoin_x:        "",
	},
	{ // 0a
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-------=========|",
		i_Adjoin_x:        "",
	},
	{ // 1
		i_interval_string: "  |=====|",
		x_interval_string: "  |------=========|",
		i_Adjoin_x:        "",
	},
	{ // 2
		i_interval_string: "  |=====|",
		x_interval_string: " *|-----=========|",
		i_Adjoin_x:        "",
	},
	{ // 3
		i_interval_string: "  |=====|",
		x_interval_string: "  |-----=========|",
		i_Adjoin_x:        "",
	},
	{ // 4
		i_interval_string: "  |=====|",
		x_interval_string: " *|----=========|",
		i_Adjoin_x:        "  |==============|",
	},
	{ // 5
		i_interval_string: "  |=====|",
		x_interval_string: "  |----=========|",
		i_Adjoin_x:        "  |=============|",
	},
	{ // 6
		i_interval_string: "  |=====|",
		x_interval_string: "  |--=========|",
		i_Adjoin_x:        "",
	},
	{ // 7
		i_interval_string: "  |=====|",
		x_interval_string: "  |-=========|",
		i_Adjoin_x:        "",
	},
	{ // 8
		i_interval_string: "  |=====|",
		x_interval_string: " *|=========|",
		i_Adjoin_x:        "",
	},
	{ // 9
		i_interval_string: "  |=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 10
		i_interval_string: " *|=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 11
		i_interval_string: "  |-=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 12
		i_interval_string: "  |--=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 13
		i_interval_string: "  |---=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 14
		i_interval_string: "  |---=====|*",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 15
		i_interval_string: "  |----=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 16
		i_interval_string: "  |----=====|*",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 17
		i_interval_string: "  |-----=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 18
		i_interval_string: "  |------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 19
		i_interval_string: "  |-------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 20
		i_interval_string: "  |--------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "  |=============|",
	},
	{ // 21
		i_interval_string: " *|--------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "  |==============|",
	},
	{ // 22
		i_interval_string: "  |---------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 23
		i_interval_string: " *|---------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 24
		i_interval_string: "  |----------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
	{ // 25
		i_interval_string: "  |-----------=====|",
		x_interval_string: "  |=========|",
		i_Adjoin_x:        "",
	},
}

var testsIntervalContains = []struct {
	test      testGeneral
	i_Cover_x bool
	x_Cover_i bool
}{
	{ // 0
		test:      testsGeneralSets[0],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[1],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[2],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[3],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[4],
		i_Cover_x: true,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[5],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[6],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[7],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[8],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[9],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[10],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[11],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[12],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[13],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[14],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[15],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[16],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[17],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[18],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[19],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[20],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[21],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[22],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[23],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[24],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[25],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[26],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[27],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[28],
		i_Cover_x: true,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[29],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[30],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[31],
		i_Cover_x: true,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[32],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[33],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[34],
		i_Cover_x: true,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[35],
		i_Cover_x: false,
		x_Cover_i: false,
	},
	{ // 0
		test:      testsGeneralSets[36],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[37],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[38],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[39],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[40],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[41],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[42],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[43],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[44],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[45],
		i_Cover_x: true,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[46],
		i_Cover_x: false,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[47],
		i_Cover_x: true,
		x_Cover_i: true,
	},
	{ // 0
		test:      testsGeneralSets[48],
		i_Cover_x: true,
		x_Cover_i: true,
	},
}

var testsIntervalLeEndOf = []struct {
	test      testGeneral
	i_LeEnd_x bool
	x_LeEnd_i bool
}{
	{ // 0
		test:      testsGeneralSets[0],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[1],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[2],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[3],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[4],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[5],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[6],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[7],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[8],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[9],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[10],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[11],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[12],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[13],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[14],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[15],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[16],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[17],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[18],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[19],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[20],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[21],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[22],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[23],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[24],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[25],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[26],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[27],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[28],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[29],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[30],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[31],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[32],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[33],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[34],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[35],
		i_LeEnd_x: false,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[36],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[37],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[38],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[39],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[40],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[41],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[42],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[43],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[44],
		i_LeEnd_x: true,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[45],
		i_LeEnd_x: true,
		x_LeEnd_i: true,
	},
	{ // 0
		test:      testsGeneralSets[46],
		i_LeEnd_x: false,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[47],
		i_LeEnd_x: false,
		x_LeEnd_i: false,
	},
	{ // 0
		test:      testsGeneralSets[48],
		i_LeEnd_x: false,
		x_LeEnd_i: false,
	},
}

var testsIntervalLtBeginOf = []struct {
	test       testGeneral
	i_Before_x bool
	x_Before_i bool
}{
	{ // 0
		test:       testsGeneralSets[0],
		i_Before_x: true,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[1],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[2],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[3],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[4],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[5],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[6],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[7],
		i_Before_x: false,
		x_Before_i: true,
	},
	{ // 0
		test:       testsGeneralSets[8],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[9],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[10],
		i_Before_x: false,
		x_Before_i: true,
	},
	{ // 0
		test:       testsGeneralSets[11],
		i_Before_x: true,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[12],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[13],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[14],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[15],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[16],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[17],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[18],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[19],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[20],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[21],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[22],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[23],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[24],
		i_Before_x: true,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[25],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[26],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[27],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[28],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[29],
		i_Before_x: false,
		x_Before_i: true,
	},
	{ // 0
		test:       testsGeneralSets[30],
		i_Before_x: true,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[31],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[32],
		i_Before_x: false,
		x_Before_i: true,
	},
	{ // 0
		test:       testsGeneralSets[33],
		i_Before_x: true,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[34],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[35],
		i_Before_x: false,
		x_Before_i: true,
	},
	{ // 0
		test:       testsGeneralSets[36],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[37],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[38],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[39],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[40],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[41],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[42],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[43],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[44],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[45],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[46],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[47],
		i_Before_x: false,
		x_Before_i: false,
	},
	{ // 0
		test:       testsGeneralSets[48],
		i_Before_x: false,
		x_Before_i: false,
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
	{ // 0	0
		i_interval_string: "  |=====|",
		x_interval_string: "  |-------=========|",
		i_intersect_x:     "",
	},
	{ // 0a	1
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-------=========|",
		i_intersect_x:     "  |-------=========|",
	},
	{ // 0b	2
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-------=========|>",
		i_intersect_x:     "  |-------=========|>",
	},
	{ // 0c	3
		i_interval_string: " <|=====|",
		x_interval_string: "  |-------=========|",
		i_intersect_x:     "",
	},
	{ // 0d	4
		i_interval_string: " <|=====|",
		x_interval_string: "  |-<-----=========|",
		i_intersect_x:     " <|=====|",
	},
	{ // 0e	5
		i_interval_string: " <|=====|>",
		x_interval_string: "  |-<-----=========|>",
		i_intersect_x:     "",
	},
	{ // 1	6
		i_interval_string: "  |=====|",
		x_interval_string: "  |----=========|",
		i_intersect_x:     "  |----=|",
	},
	{ // 1a	7
		i_interval_string: "  |=====|>",
		x_interval_string: "  |----=========|",
		i_intersect_x:     "  |----=========|",
	},
	{ // 1b	8
		i_interval_string: "  |=====|>",
		x_interval_string: "  |----=========|>",
		i_intersect_x:     "  |----=|>",
	},
	{ // 1c	9
		i_interval_string: " <|=====|",
		x_interval_string: "  |----=========|",
		i_intersect_x:     "  |----==|",
	},
	{ // 1d	10
		i_interval_string: " <|=====|",
		x_interval_string: "  |----<=========|",
		i_intersect_x:     "  |----<=|",
	},
	{ // 1e	11
		i_interval_string: " <|=====|>",
		x_interval_string: "  |----<=========|>",
		i_intersect_x:     "  |----<=|>",
	},
	{ // 6
		i_interval_string: "  |=====|",
		x_interval_string: "  |--=========|",
		i_intersect_x:     "  |--===|",
	},
	{ // 7
		i_interval_string: "  |=====|",
		x_interval_string: "  |-=========|",
		i_intersect_x:     "  |-====|",
	},
	{ // 8
		i_interval_string: "  |=====|",
		x_interval_string: " *|=========|",
		i_intersect_x:     " *|====|",
	},
	{ // 9
		i_interval_string: "  |=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |=====|",
	},
	{ // 10
		i_interval_string: " *|=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     " *|=====|",
	},
	{ // 11
		i_interval_string: "  |-=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |-=====|",
	},
	{ // 12
		i_interval_string: "  |--=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |--=====|",
	},
	{ // 13
		i_interval_string: "  |---=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |---=====|",
	},
	{ // 14
		i_interval_string: "  |---=====|*",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |---=====|*",
	},
	{ // 15
		i_interval_string: "  |----=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |----=====|",
	},
	{ // 16
		i_interval_string: "  |----=====|*",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |----=====|",
	},
	{ // 17
		i_interval_string: "  |-----=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |-----====|",
	},
	{ // 18
		i_interval_string: "  |------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |------===|",
	},
	{ // 19
		i_interval_string: "  |-------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |-------==|",
	},
	{ // 20
		i_interval_string: "  |--------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "  |--------=|",
	},
	{ // 21
		i_interval_string: " *|--------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "",
	},
	{ // 22
		i_interval_string: "  |---------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "",
	},
	{ // 23
		i_interval_string: " *|---------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "",
	},
	{ // 24
		i_interval_string: "  |----------=====|",
		x_interval_string: "  |=========|",
		i_intersect_x:     "",
	},
	{ // 25
		i_interval_string: "  |-----------=====|",
		x_interval_string: "  |=========|",
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
		i_interval_string: "  |=====|",
		x_interval_string: "  |-------=========|",
		i_Encompass_x:     "  |================|",
	},
	{ // 0a
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-------=========|",
		i_Encompass_x:     "  |================|",
	},
	{ // 1
		i_interval_string: "  |=====|",
		x_interval_string: "  |------=========|",
		i_Encompass_x:     "  |===============|",
	},
	{ // 2
		i_interval_string: "  |=====|",
		x_interval_string: " *|-----=========|",
		i_Encompass_x:     "  |===============|",
	},
	{ // 3
		i_interval_string: "  |=====|",
		x_interval_string: "  |-----=========|",
		i_Encompass_x:     "  |==============|",
	},
	{ // 4
		i_interval_string: "  |=====|",
		x_interval_string: " *|----=========|",
		i_Encompass_x:     "  |==============|",
	},
	{ // 5
		i_interval_string: "  |=====|",
		x_interval_string: "  |----=========|",
		i_Encompass_x:     "  |=============|",
	},
	{ // 6
		i_interval_string: "  |=====|",
		x_interval_string: "  |--=========|",
		i_Encompass_x:     "  |===========|",
	},
	{ // 7
		i_interval_string: "  |=====|",
		x_interval_string: "  |-=========|",
		i_Encompass_x:     "  |==========|",
	},
	{ // 8
		i_interval_string: "  |=====|",
		x_interval_string: " *|=========|",
		i_Encompass_x:     "  |==========|",
	},
	{ // 9
		i_interval_string: "  |=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 10
		i_interval_string: " *|=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 11
		i_interval_string: "  |-=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 12
		i_interval_string: "  |--=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 13
		i_interval_string: "  |---=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 14
		i_interval_string: "  |---=====|*",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 15
		i_interval_string: "  |----=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|",
	},
	{ // 16
		i_interval_string: "  |----=====|*",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=========|*",
	},
	{ // 17
		i_interval_string: "  |-----=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |==========|",
	},
	{ // 18
		i_interval_string: "  |------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |===========|",
	},
	{ // 19
		i_interval_string: "  |-------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |============|",
	},
	{ // 20
		i_interval_string: "  |--------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |=============|",
	},
	{ // 21
		i_interval_string: " *|--------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |==============|",
	},
	{ // 22
		i_interval_string: "  |---------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |==============|",
	},
	{ // 23
		i_interval_string: " *|---------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |===============|",
	},
	{ // 24
		i_interval_string: "  |----------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |===============|",
	},
	{ // 25
		i_interval_string: "  |-----------=====|",
		x_interval_string: "  |=========|",
		i_Encompass_x:     "  |================|",
	},
}
