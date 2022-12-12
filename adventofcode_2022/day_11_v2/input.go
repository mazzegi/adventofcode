package day_11_v2

var input = []InputMonkey{
	{
		Items:          []int{72, 97},
		Operation:      MultBy(13),
		TestDivBy:      19,
		ThrowToIfTrue:  5,
		ThrowToIfFalse: 6,
	},
	{
		Items:          []int{55, 70, 90, 74, 95},
		Operation:      Square,
		TestDivBy:      7,
		ThrowToIfTrue:  5,
		ThrowToIfFalse: 0,
	},
	{
		Items:          []int{74, 97, 66, 57},
		Operation:      Add(6),
		TestDivBy:      17,
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 0,
	},
	{
		Items:          []int{86, 54, 53},
		Operation:      Add(2),
		TestDivBy:      13,
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 2,
	},
	{
		Items:          []int{50, 65, 78, 50, 62, 99},
		Operation:      Add(3),
		TestDivBy:      11,
		ThrowToIfTrue:  3,
		ThrowToIfFalse: 7,
	},
	{
		Items:          []int{90},
		Operation:      Add(4),
		TestDivBy:      2,
		ThrowToIfTrue:  4,
		ThrowToIfFalse: 6,
	},
	{
		Items:          []int{88, 92, 63, 94, 96, 82, 53, 53},
		Operation:      Add(8),
		TestDivBy:      5,
		ThrowToIfTrue:  4,
		ThrowToIfFalse: 7,
	},
	{
		Items:          []int{70, 60, 71, 69, 77, 70, 98},
		Operation:      MultBy(7),
		TestDivBy:      3,
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
	},
}
