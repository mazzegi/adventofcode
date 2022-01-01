a = 1

set b 93            b = 93
set c b             c = 93
jnz a 2 
jnz 1 5             nop
mul b 100           b = 9300
sub b -100000       b = 109300
set c b             c = 109300
sub c -17000        c = 126300

    set f 1             f = 1
    set d 2             d = 2
        set e 2             e = 2
            set g d             g = 2
            mul g e             g = 4
            sub g b             g = 109304
            jnz g 2             
            set f 0
            sub e -1            e = 3
            set g e             g = 3
            sub g b             g = -109297
            jnz g -8            if g
        sub d -1
        set g d
        sub g b
        jnz g -13
    jnz f 2
    sub h -1
    set g b         g = 109300
    sub g c         g = -17000
    jnz g 2
    jnz 1 3    exit!
    sub b -17 // loops 1000 times
jnz 1 -23