=> a = 12

cpy a b     b = 12

dec b       b = 11

# loop entry :bloop
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


tgl c           (0) b == 0 => inc c

cpy -16 c
bloop: cpy 1 c     jnz 1 c         (2) b == 1 => cpy 1 c

cpy 78 c

cpy 70 d    jnz 70 d        (4) b == 2 => cpy 70 d

inc a
dec d        inc d           (6) b == 3 => dec d
jnz d -2



dec c       inc c           (8) b == 4 => dec c
jnz c -5