p = 0

set i 31        i = 31
set a 1         a = 1
mul p 17        p = 0
jgz p p         nop

mul a 2         a = 2   a = 4 ---- 2^30 => 1073741824
add i -1        i = 30  i = 29
jgz i -2        

add a -1        a = 1073741824 - 1
set i 127       i = 127
set p 735       p = 735

mul p 8505      p = 6251175
mod p a         p = 6251175
mul p 129749    p = 811083705075
add p 12345     p = 811083717420
mod p a         p = 408640300
set b p         b = 408640300
mod b 10000     b = 300
snd b           send 300
add i -1        i = 126
jgz i -9        9 times

jgz a 3         
rcv b
jgz b -1
set f 0

set i 126       i = 126
rcv a           <- a
rcv b           <- b

set p a         p = a
mul p -1        p = -a
add p b         
jgz p 4

snd a
set a b
jgz 1 3

snd b
set f 1
add i -1
jgz i -11 until i = 0

snd a  a ->
jgz f -16

jgz a -19