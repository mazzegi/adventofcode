=> a = (12/7)

b = a

for b > 0{
    b--

    d = a
    a = 0

    for d > 0{
        c = b
        for c > 0{
            a++
            c--
        }
        d--
        // c = 0; a = a+c
    }
    //b = 2; a = a*b
    // c = 0; a = a + d*b => c = 0; a = a*b

    b--
    c = b
    d = c

    for d > 0{
        d--
        c++
    }
    // d = 0; c = 2*b; a = a*b

    ??

} //c = -16
//d = 0; c = 2*b; a = a*b


c = 78

for c > 0{
    d = 70

    for d > 0{
        a++
        d--
    }
    //d = 0; a = a + 70
    c--
}
//d = 0; c = 0; a = a*b+70*78


//a = a*b + 70*78
//a = a*a + 70*78




