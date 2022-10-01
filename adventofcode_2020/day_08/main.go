/*
--- Day 8: Handheld Halting ---
Your flight to the major airline hub reaches cruising altitude without incident. While you consider checking the in-flight menu for one of those drinks that come with a little umbrella, you are interrupted by the kid sitting next to you.

Their handheld game console won't turn on! They ask if you can take a look.

You narrow the problem down to a strange infinite loop in the boot code (your puzzle input) of the device. You should be able to fix it, but first you need to be able to run the code in isolation.

The boot code is represented as a text file with one instruction per line of text. Each instruction consists of an operation (acc, jmp, or nop) and an argument (a signed number like +4 or -20).

acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
For example, consider the following program:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
These instructions are visited in this order:

nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
First, the nop +0 does nothing. Then, the accumulator is increased from 0 to 1 (acc +1) and jmp +4 sets the next instruction to the other acc +1 near the bottom. After it increases the accumulator from 1 to 2, jmp -4 executes, setting the next instruction to the only acc +3. It sets the accumulator to 5, and jmp -3 causes the program to continue back at the first acc +1.

This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction a second time, you know it will never terminate.

Immediately before the program would run an instruction a second time, the value in the accumulator is 5.

Run your copy of the boot code. Immediately before any instruction is executed a second time, what value is in the accumulator?
*/

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mazzegi/adventofcode/adventofcode_2020/instruction"
)

func main() {
	buf := bytes.NewBufferString(input)
	//buf := bytes.NewBufferString(input)
	is, err := instruction.ParseMany(buf)
	if err != nil {
		panic(err)
	}
	//	m := instruction.NewMachine(is)
	// n, err := m.ExecOneLoop()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("acc: %d in (%d)\n", m.Acc(), n)

	//m.Reset()
	// n, err := m.Exec()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("acc: %d in (%d)\n", m.Acc(), n)

	m := instruction.NewMachine(is)
	_, err = m.Exec()
	if err == nil {
		fmt.Println("machine terminates properly")
		os.Exit(0)
	}
	for i := 0; i < len(is); i++ {
		if is[i].Op == instruction.Acc {
			continue
		}
		if is[i].Op == instruction.Jmp {
			is[i].Op = instruction.Nop
			m := instruction.NewMachine(is)
			if _, err := m.Exec(); err == nil {
				fmt.Printf("works after changing %d from jmp->nop: acc = %d\n", i, m.Acc())
				break
			}
			is[i].Op = instruction.Jmp
		} else if is[i].Op == instruction.Nop {
			is[i].Op = instruction.Jmp
			m := instruction.NewMachine(is)
			if _, err := m.Exec(); err == nil {
				fmt.Printf("works after changing %d from nop->jmp: acc = %d\n", i, m.Acc())
				break
			}
			is[i].Op = instruction.Nop
		}
	}
}

var inputTest = `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

var inputTest1 = `
jmp +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

var inputTest2 = `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
nop -4
acc +6
`

var input = `
acc -17
nop +318
jmp +1
acc -10
jmp +394
acc +43
acc +47
nop +570
jmp +176
acc -9
jmp +322
jmp +73
acc +4
acc -4
jmp +460
jmp +228
acc +25
acc +39
jmp +50
acc -12
acc -14
nop +275
jmp +489
acc -11
jmp +338
acc +21
acc +10
jmp +1
acc +20
jmp +445
acc +7
jmp +419
acc -8
acc +32
jmp +181
acc +19
acc +5
acc +46
jmp +417
acc +28
acc +23
acc +16
jmp +225
jmp +317
jmp +309
jmp +69
acc -6
acc -6
jmp +127
acc +49
nop -38
jmp +467
acc +0
acc -12
acc -14
jmp -46
acc +14
acc +2
acc +2
jmp +311
acc +33
jmp +364
nop +234
acc +24
acc +37
acc +18
jmp +22
jmp +303
jmp +414
jmp +318
acc +22
acc +27
jmp +1
jmp +150
acc +34
acc +15
nop +200
acc +6
jmp +320
nop +534
acc +29
jmp +147
nop -20
jmp +255
jmp +10
acc -15
acc +3
jmp +338
nop +362
acc -4
jmp +1
jmp +286
acc -1
jmp +497
acc -4
acc +23
acc +4
jmp +400
acc +35
acc +50
jmp +133
acc -17
jmp -90
jmp +7
acc -17
jmp +472
acc +20
jmp +280
jmp +133
jmp -15
jmp +16
acc -19
acc -2
jmp -64
acc -17
jmp +1
jmp +385
acc -5
acc +34
jmp +382
acc +24
acc -17
acc +0
acc +15
jmp +466
jmp +300
acc +16
jmp +302
nop +479
acc +16
jmp +71
acc +23
jmp +1
acc +8
jmp +154
jmp +410
acc -8
jmp +402
acc +48
acc +42
acc +22
acc +35
jmp +50
jmp -7
acc -13
acc +37
acc +24
jmp +243
jmp +410
acc -3
acc +45
jmp +416
acc +2
acc +25
jmp -109
jmp -41
jmp +318
acc -8
acc -12
jmp +169
nop +393
acc +7
acc -12
acc +35
jmp +381
acc +41
nop -98
acc +15
acc -19
jmp +218
acc +24
acc +47
jmp +65
acc +29
jmp -129
acc +23
acc -13
nop +60
jmp -26
nop -4
acc -5
acc +13
nop -12
jmp -13
jmp -53
acc +21
jmp +276
nop -27
jmp +165
acc +42
nop +43
jmp +1
acc +26
acc +22
acc -3
jmp +405
acc +29
nop -118
acc +21
nop -190
jmp +217
acc -1
nop +223
acc -8
acc +45
jmp +49
acc +8
acc +22
jmp +209
acc +44
jmp +66
acc +7
acc -7
acc +48
jmp +318
nop +398
acc +2
jmp +16
nop +207
nop +358
acc +45
acc +48
jmp +267
nop +248
acc +26
jmp +307
acc +27
jmp -197
jmp -68
acc +34
acc +25
acc -13
jmp +133
jmp -77
acc -13
acc +10
jmp -193
jmp -62
acc +4
acc -14
jmp +261
jmp +151
jmp +208
acc -10
jmp +40
acc +31
jmp -216
acc +23
acc +34
jmp +364
nop +205
acc -3
acc +14
jmp +59
nop +359
acc -4
jmp +1
jmp -248
acc +47
acc +35
jmp +184
acc +16
nop -92
acc -12
jmp +354
acc +27
jmp -152
acc -14
acc -16
acc +43
jmp +147
acc +45
acc +24
acc +6
nop -46
jmp +21
acc +26
jmp +1
jmp +293
acc -8
acc +12
acc -19
acc -9
jmp +94
jmp +299
acc +10
acc -2
jmp +75
acc -7
acc +3
acc +47
jmp +171
acc +16
acc +44
acc -3
jmp +14
acc +30
acc +34
jmp -178
acc +35
nop -238
acc +39
jmp +1
jmp -133
acc +34
acc -6
jmp -276
acc +1
jmp -207
acc +10
jmp -43
jmp -302
acc -1
nop -29
jmp +1
acc +17
jmp -281
acc +17
jmp -109
jmp +1
acc +13
nop -9
jmp +245
acc +5
nop -15
acc +3
acc +7
jmp +65
acc -11
jmp -313
acc +47
jmp +29
jmp -289
acc +18
acc -17
nop +73
acc -12
jmp +80
acc +32
acc -4
acc +3
jmp -126
acc +16
jmp -275
nop -188
acc -3
acc +14
jmp -155
acc +33
acc -19
nop -166
acc +20
jmp +30
nop -169
acc +49
nop +168
jmp -24
nop -345
acc +34
jmp -40
jmp -56
jmp +29
jmp +191
acc +24
jmp +219
acc +34
acc +27
acc +11
jmp -260
jmp -339
acc +15
nop +16
jmp +1
jmp +138
jmp +1
jmp +1
jmp +14
acc -11
acc +45
jmp -19
acc +0
jmp +27
acc +0
nop +128
jmp -65
nop -23
jmp -318
jmp -325
jmp +1
jmp -229
jmp -270
jmp -137
acc +34
acc +7
jmp +1
jmp -346
acc +18
jmp +37
acc +40
acc -16
nop -146
acc +35
jmp -12
acc +1
acc +27
acc +44
acc +8
jmp -276
acc +16
acc +42
nop -342
acc +13
jmp -165
acc -11
acc -17
acc -10
jmp -26
acc +10
acc +43
jmp -276
acc +5
acc +34
acc +17
acc -9
jmp +99
acc +29
jmp -370
acc -11
jmp -412
acc +47
acc +21
acc -12
jmp -136
jmp -124
acc +12
acc +0
acc +25
acc +27
jmp -290
acc +5
acc +49
acc +32
nop +29
jmp -202
nop -296
acc -12
acc +9
acc +21
jmp +23
jmp -345
acc +26
nop -123
jmp -373
nop +118
jmp +43
acc -15
jmp -386
jmp +1
nop -370
acc +47
nop -141
jmp -426
acc +42
acc +12
acc +4
nop -103
jmp -122
acc +23
acc -4
acc +11
jmp -314
jmp -73
nop -1
jmp -411
acc +13
acc +9
nop -372
jmp -293
acc +46
acc +3
acc -1
jmp +86
acc +36
jmp +100
acc +27
acc +49
nop -4
acc +47
jmp -445
acc +31
acc +47
acc -11
acc +14
jmp -181
nop -438
acc +31
jmp -428
nop -115
nop -244
jmp -464
jmp -29
nop -240
jmp -241
acc -12
jmp -329
nop +78
acc +6
jmp +1
acc +49
jmp -322
jmp -133
acc +20
nop -83
acc +35
acc +29
jmp -41
acc +15
jmp -46
jmp -29
acc +45
acc -14
acc +21
jmp -366
nop +84
acc -6
acc +25
acc -17
jmp -326
acc -5
nop -159
acc +5
jmp -171
acc +42
jmp -28
acc +42
acc -11
acc +45
acc +19
jmp -305
acc +38
acc -13
acc -16
jmp -134
acc +45
jmp -256
acc -15
acc -18
acc +28
jmp -114
acc -11
acc +47
nop -420
jmp -90
nop -330
jmp +13
acc -15
acc +9
jmp -159
acc -12
acc +0
acc +0
jmp -538
acc +31
acc +24
acc +32
acc -16
jmp -95
jmp -466
acc +19
acc +2
jmp -172
acc -12
jmp -207
acc +39
acc +18
acc +5
jmp -211
nop -507
jmp +1
jmp -197
nop -227
acc +28
jmp -494
acc +22
acc +2
acc -14
jmp -377
acc +8
acc +29
jmp -573
acc -17
acc +14
acc +29
acc +11
jmp -351
acc +9
nop -540
acc +30
nop -344
jmp -564
acc -4
nop -465
jmp -293
acc -18
acc +5
acc +29
jmp -302
acc -17
acc +14
acc +2
acc -11
jmp -527
jmp -563
acc +14
acc +10
jmp -505
acc +43
jmp -188
nop -448
acc +44
acc +3
acc +16
jmp +1
`
