package base

/*
Reading the test-sets
The '========' stands for interval-range
Upper- en lower-Included is true
The '<' is for lowerUnbound, no lower-end
The '>' is for upperUnbound, no upper-end
The '*' is for Included is false (depending on the site for upper en lower included)

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

# On the uppersite are mutatis mutandis the same rules

So, these two intervals below, although visible overlapping, do NOT overlap
<*|====--------------| 0,4
*|----=======|*       5,12
So it can be written like this, which is much more readable
<*|====--------------| 0,4
* |----=======|*       5,12

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
		i_interval_string: "  |=====|",
		x_interval_string: "  |-------=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 0a
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-------=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 1
		i_interval_string: "  |=====|",
		x_interval_string: "  |------=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 2
		i_interval_string: "  |=====|",
		x_interval_string: " *|-----=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 3
		i_interval_string: "  |=====|",
		x_interval_string: "  |-----=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 4
		i_interval_string: "  |=====|",
		x_interval_string: " *|----=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 5
		i_interval_string: "  |=====|",
		x_interval_string: "  |----=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 6
		i_interval_string: "  |=====|",
		x_interval_string: "  |--=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 7
		i_interval_string: "  |=====|",
		x_interval_string: "  |-=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 8
		i_interval_string: "  |=====|",
		x_interval_string: " *|=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 9
		i_interval_string: "  |=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 10
		i_interval_string: " *|=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 11
		i_interval_string: "  |-=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 12
		i_interval_string: "  |--=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 13
		i_interval_string: "  |---=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 14
		i_interval_string: "  |---=====|*",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 15
		i_interval_string: "  |----=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         true,
	},
	{ // 16
		i_interval_string: "  |----=====|*",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 17
		i_interval_string: "  |-----=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 18
		i_interval_string: "  |------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 19
		i_interval_string: "  |-------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 20
		i_interval_string: "  |--------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 21
		i_interval_string: " *|--------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 22
		i_interval_string: "  |---------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 23
		i_interval_string: " *|---------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 24
		i_interval_string: "  |----------=====|",
		x_interval_string: "  |=========|",
		i_Cover_x:         false,
		x_Cover_i:         false,
	},
	{ // 25
		i_interval_string: "  |-----------=====|",
		x_interval_string: "  |=========|",
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

	counter string
}{
	{ // 0
		i_interval_string: "  |=======|",
		x_interval_string: "  |-------========|",
		i_Before_x:        true,
		x_Before_i:        false,
		counter:           "0",
	},
	{ // 0a
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-----=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "0a",
	},
	{ // 0b
		i_interval_string: "  |=====|>",
		x_interval_string: "  |-----=========|>",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "0b",
	},
	{ // 0c
		i_interval_string: "  |=====|",
		x_interval_string: "  |-----=========|>",
		i_Before_x:        true,
		x_Before_i:        false,
		counter:           "0c",
	},
	{ // 0d
		i_interval_string: " <|=====|",
		x_interval_string: "  |------=========|",
		i_Before_x:        true,
		x_Before_i:        false,
		counter:           "0d",
	},
	{ // 0e
		i_interval_string: " <|=====|",
		x_interval_string: "<-----=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "0e",
	},
	{ // 0f
		i_interval_string: "  |=====|",
		x_interval_string: "<-----=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "0f",
	},
	{ // 0g
		i_interval_string: " <|=====|>",
		x_interval_string: "<-----=========|>",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "0g",
	},
	{ // 1
		i_interval_string: "  |=======|",
		x_interval_string: "  |------=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1",
	},
	{ // 1a
		i_interval_string: "  |=======|>",
		x_interval_string: "  |------==========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1a",
	},
	{ // 1b
		i_interval_string: "  |========|>",
		x_interval_string: "  |-------=========|>",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1b",
	},
	{ // 1c
		i_interval_string: "  |========|",
		x_interval_string: "  |------=========|>",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1c",
	},
	{ // 1d
		i_interval_string: " <|========|",
		x_interval_string: "  |-------=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1d",
	},
	{ // 1e
		i_interval_string: " <|=======|",
		x_interval_string: "<-----=========|",
		i_Before_x:        false,
		x_Before_i:        false,
		counter:           "1e",
	},
	{ // 1
		i_interval_string: "  |=====|>",
		x_interval_string: "  |------=========|>",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 2
		i_interval_string: "  |=====|",
		x_interval_string: " *|-----=========|",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 3
		i_interval_string: "  |=====|",
		x_interval_string: "  |-----=========|",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 4
		i_interval_string: "  |=====|",
		x_interval_string: " *|----=========|",
		i_Before_x:        true,
		x_Before_i:        false,
	},
	{ // 5
		i_interval_string: "  |=====|",
		x_interval_string: "  |----=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 6
		i_interval_string: "  |=====|",
		x_interval_string: "  |--=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 7
		i_interval_string: "  |=====|",
		x_interval_string: "  |-=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 8
		i_interval_string: "  |=====|",
		x_interval_string: " *|=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 9
		i_interval_string: "  |=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 10
		i_interval_string: " *|=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 11
		i_interval_string: "  |-=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 12
		i_interval_string: "  |--=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 13
		i_interval_string: "  |---=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 14
		i_interval_string: "  |---=====|*",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 15
		i_interval_string: "  |----=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 16
		i_interval_string: "  |----=====|*",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 17
		i_interval_string: "  |-----=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 18
		i_interval_string: "  |------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 19
		i_interval_string: "  |-------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 20
		i_interval_string: "  |--------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        false,
	},
	{ // 21
		i_interval_string: " *|--------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 22
		i_interval_string: "  |---------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 23
		i_interval_string: " *|---------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 24
		i_interval_string: "  |----------=====|",
		x_interval_string: "  |=========|",
		i_Before_x:        false,
		x_Before_i:        true,
	},
	{ // 25
		i_interval_string: "  |-----------=====|",
		x_interval_string: "  |=========|",
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
