package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	in := input
	for n := 1; true; n++ {
		sn := fmt.Sprintf("%s%d", in, n)
		h := fmt.Sprintf("%x", md5.Sum([]byte(sn)))
		if strings.HasPrefix(h, "000000") {
			fmt.Printf("%d => %s\n", n, h)
			break
		}
	}
}

var inputTest0 = `abcdef`
var inputTest1 = `pqrstuv`
var input = `yzbqklnj`
