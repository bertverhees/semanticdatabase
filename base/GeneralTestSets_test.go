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

type testGeneral struct {
	i_interval_string string
	x_interval_string string
	counter           string
}

var testsGeneralSets = []testGeneral{
	{ //0
		i_interval_string: "  |=====---------------|  ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "0",
	},
	{ //1
		i_interval_string: "  |======--------------|  ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "1",
	},
	{ //2
		i_interval_string: "  |---=====------------|  ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "2",
	},
	{ //3
		i_interval_string: "  |-------=====--------|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "3",
	},
	{ //4
		i_interval_string: "  |------========------|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "4",
	},
	{ //5
		i_interval_string: "  |-----------=====----|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "5",
	},
	{ //6
		i_interval_string: "  |--------------======|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "6",
	},
	{ //7
		i_interval_string: "  |---------------=====|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "7",
	},
	{ //8
		i_interval_string: " *|======--------------|  ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "8",
	},
	{ //9
		i_interval_string: " *|------========------|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "9",
	},
	{ //10
		i_interval_string: " *|--------------======|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "10",
	},
	{ //11
		i_interval_string: "  |======--------------|* ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "11",
	},
	{ //12
		i_interval_string: "  |------========------|* ",
		x_interval_string: "  |------========------|  ",
		counter:           "12",
	},
	{ //13
		i_interval_string: "  |--------------======|* ",
		x_interval_string: "  |------========------|  ",
		counter:           "13",
	},
	{ //14
		i_interval_string: " <|-------=====--------|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "14",
	},
	{ //15
		i_interval_string: " <|------========------|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "15",
	},
	{ //16
		i_interval_string: " <|-----------=====----|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "16",
	},
	{ //17
		i_interval_string: " <|--------------======|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "17",
	},
	{ //18
		i_interval_string: " <|---------------=====|  ",
		x_interval_string: "  |------========------|  ",
		counter:           "18",
	},
	{ //19
		i_interval_string: "  |=====---------------|> ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "19",
	},
	{ //20
		i_interval_string: "  |======--------------|> ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "20",
	},
	{ //21
		i_interval_string: "  |---=====------------|> ",
		x_interval_string: "  |------=======-------|  ",
		counter:           "21",
	},
	{ //22
		i_interval_string: "  |-------=====--------|> ",
		x_interval_string: "  |------========------|  ",
		counter:           "22",
	},
	{ //23
		i_interval_string: "  |------========------|> ",
		x_interval_string: "  |------========------|  ",
		counter:           "23",
	},
	{ //24
		i_interval_string: "  |======--------------|  ",
		x_interval_string: " *|------=======-------|  ",
		counter:           "24",
	},
	{ //25
		i_interval_string: "  |------========------|  ",
		x_interval_string: " *|------========------|  ",
		counter:           "25",
	},
	{ //26
		i_interval_string: "  |--------------======|  ",
		x_interval_string: " *|------========------|  ",
		counter:           "26",
	},
	{ //27
		i_interval_string: "  |======--------------|  ",
		x_interval_string: "  |------=======-------|*  ",
		counter:           "27",
	},
	{ //28
		i_interval_string: "  |------========------|  ",
		x_interval_string: "  |------========------|* ",
		counter:           "28",
	},
	{ //29
		i_interval_string: "  |--------------======|  ",
		x_interval_string: "  |------========------|*  ",
		counter:           "29",
	},
	{ //30
		i_interval_string: " *|======--------------|  ",
		x_interval_string: " *|------=======-------|  ",
		counter:           "30",
	},
	{ //31
		i_interval_string: " *|------========------|  ",
		x_interval_string: " *|------========------|  ",
		counter:           "31",
	},
	{ //32
		i_interval_string: " *|--------------======|  ",
		x_interval_string: " *|------========------|  ",
		counter:           "32",
	},
	{ //33
		i_interval_string: "  |======--------------|* ",
		x_interval_string: "  |------=======-------|* ",
		counter:           "33",
	},
	{ //34
		i_interval_string: "  |------========------|* ",
		x_interval_string: "  |------========------|* ",
		counter:           "34",
	},
	{ //35
		i_interval_string: "  |--------------======|* ",
		x_interval_string: "  |------========------|* ",
		counter:           "35",
	},
	{ //36
		i_interval_string: "  |=====---------------|  ",
		x_interval_string: " <|------=======-------|  ",
		counter:           "36",
	},
	{ //37
		i_interval_string: "  |======--------------|  ",
		x_interval_string: " <|------=======-------|  ",
		counter:           "37",
	},
	{ //38
		i_interval_string: "  |---=====------------|  ",
		x_interval_string: " <|------=======-------|  ",
		counter:           "38",
	},
	{ //39
		i_interval_string: "  |-------=====--------|  ",
		x_interval_string: " <|------========------|  ",
		counter:           "39",
	},
	{ //40
		i_interval_string: "  |------========------|  ",
		x_interval_string: "  |------========------|> ",
		counter:           "40",
	},
	{ //41
		i_interval_string: "  |-----------=====----|  ",
		x_interval_string: "  |------========------|> ",
		counter:           "41",
	},
	{ //42
		i_interval_string: "  |--------------======|  ",
		x_interval_string: "  |------========------|> ",
		counter:           "42",
	},
	{ //43
		i_interval_string: "  |---------------=====|  ",
		x_interval_string: "  |------========------|> ",
		counter:           "43",
	},
	{ //44
		i_interval_string: " <|-------=====--------|  ",
		x_interval_string: " <|------========------|  ",
		counter:           "44",
	},
	{ //45
		i_interval_string: " <|------========------|  ",
		x_interval_string: " <|------========------|  ",
		counter:           "45",
	},
	{ //46
		i_interval_string: "  |-------=====--------|> ",
		x_interval_string: "  |------========------|> ",
		counter:           "46",
	},
	{ //47
		i_interval_string: " <|------========------|> ",
		x_interval_string: " <|------========------|> ",
		counter:           "47",
	},
	{ //48
		i_interval_string: " <|--------====--------|> ",
		x_interval_string: " <|------========------|> ",
		counter:           "48",
	},
}
