=> a = 12

cpy a b     b = 12
dec b       b = 11
cpy a d     d = 12
cpy 0 a     a = 0

    cpy b c     c = 11

        inc a       a = 1
        dec c       c = 10
        jnz c -2      // loops until c == 0 

    dec d
    jnz d -5        // loops until d == 0

dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 78 c
jnz 70 d
inc a
inc d
jnz d -2
inc c
jnz c -5