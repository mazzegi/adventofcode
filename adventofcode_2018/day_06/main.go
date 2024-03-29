package main

import (
	"bytes"

	"github.com/mazzegi/adventofcode/adventofcode_2018/coord"
)

func main() {
	//in := inputTest
	in := input
	g, err := coord.ParseGrid(bytes.NewBufferString(in))
	if err != nil {
		panic(err)
	}
	//g.CalcDistances()
	g.CalcSafest(10000)
}

var inputTest = `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

var input = `
260, 78
42, 40
87, 276
219, 124
166, 137
341, 138
82, 121
114, 174
218, 289
61, 358
328, 164
279, 50
218, 107
273, 320
192, 349
354, 103
214, 175
128, 196
237, 67
333, 150
98, 260
166, 217
92, 212
55, 165
205, 138
321, 199
285, 148
217, 130
357, 319
160, 67
63, 75
345, 123
316, 220
41, 253
240, 245
201, 124
336, 166
95, 301
55, 181
219, 315
209, 237
317, 254
314, 300
242, 295
295, 293
285, 263
330, 204
112, 106
348, 49
81, 185
`
