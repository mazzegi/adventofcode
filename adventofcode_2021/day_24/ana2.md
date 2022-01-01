=> z0
****
inp w           w = in
mul x 0         x = 0
add x z         x = z0
mod x 26        x = z0 % 26
*div z C0       z = z0 / C0
*add x C1       x = z0 % 26 + C1
eql x w         z0 % 26 + C1 == in ? x = 1 : x = 0
eql x 0         z0 % 26 + C1 == in ? x = 0 : x = 1
mul y 0         y = 0
add y 25        y = 25
mul y x         z0 % 26 + C1 == in ? y = 0 : y = 25
add y 1         z0 % 26 + C1 == in ? y = 1 : y = 26
mul z y         z0 % 26 + C1 == in ? z = z0 / C0 : z = z0 / C0 * 26
mul y 0         y = 0
add y w         y = in
*add y C2       y = in + C2
mul y x         z0 % 26 + C1 == in ? y = 0 : y = in + C2
add z y         z0 % 26 + C1 == in ? z = z0 / C0: z = z0 / C0 * 26 + in + C2