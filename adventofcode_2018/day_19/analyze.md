ip: 5

 0: addi 5 16 5                 goto 17

 1: seti 1 7 3                  r3 = 1

 2: seti 1 4 1                  r1 = 1

 3: mulr 3 1 4                  r4 = r1 * r3
 4: eqrr 4 2 4                  r4 = ( r4 == r2 )
 5: addr 4 5 5                  r5 = r4 + 5 => goto r4 + 6 => goto 6 if r2 != r4 /  goto 7 if r2 == r4

 6: addi 5 1 5                  r5 = 1 + 6 => goto 1+7 => goto 8
 7: addr 3 0 0                  r0 = r0 + r3 => !!!

 8: addi 1 1 1                  r1 = r1 + 1
 9: gtrr 1 2 4                  r4 = (r1 > r2)
10: addr 5 4 5                  r5 = r4 + 10 => goto r4 + 11 => goto 11 if r1 <= r2 / goto 12 if r1 > r2

11: seti 2 1 5                  r5 = 2 => goto 3

12: addi 3 1 3                  r3 = r3 + 1
13: gtrr 3 2 4                  r4 = (r3 > r2)
14: addr 4 5 5                  r5 = r4 + 14 => goto r4 + 15 => goto 15 if r3 <= r2 / goto 16 if r3 > r2

15: seti 1 4 5                  r5 = 1 => goto 2
16: mulr 5 5 5                  r5 = 16 * 16 = 256 => goto 257 => EXIT!

17: addi 2 2 2                  r2 = r2 + 2
18: mulr 2 2 2                  r2 = r2 * r2
19: mulr 5 2 2                  r2 = r2 * 19
20: muli 2 11 2                 r2 = r2 * 11
21: addi 4 1 4                  r4 = r4 + 1
22: mulr 4 5 4                  r4 = r4 * 22
23: addi 4 19 4                 r4 = r4 + 19
24: addr 2 4 2                  r2 = r2 + r4

25: addr 5 0 5                  r5 = 25 + r0 => goto r0 + 26
26: seti 0 9 5                  r5 = 0 => goto 1

27: setr 5 7 4                  r4 = 27
28: mulr 4 5 4                  r4 = r4 * 28
29: addr 5 4 4                  r4 = r4 + 29
30: mulr 5 4 4                  r4 = r4 * 30
31: muli 4 14 4                 r4 = r4 * 14
32: mulr 4 5 4                  r4 = r4 * 32
33: addr 2 4 2                  r2 = r2 + r4
34: seti 0 9 0                  set r0 = 0
35: seti 0 6 5                  r5 = 0 => goto 1