package day_05

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	pwd := Password("cxdnnyjw")
	fmt.Printf("password: %q\n", pwd)
}

func Part2() {
	pwd := Password2("cxdnnyjw")
	fmt.Printf("password2: %q\n", pwd)
}

func Password(doorID string) string {
	var idx int = 0
	var pwd string
	for {
		h := hash(doorID, idx)
		if strings.HasPrefix(h, "00000") {
			pwd += string(h[5])
			if len(pwd) == 8 {
				return pwd
			}
		}
		idx++
	}
}

func Password2(doorID string) string {
	var idx int = 0
	pwd := []byte("________")

	found := map[int]bool{}
	isComplete := func() bool {
		for i := 0; i < 8; i++ {
			if _, ok := found[i]; !ok {
				return false
			}
		}
		return true
	}

	for {
		h := hash(doorID, idx)
		if strings.HasPrefix(h, "00000") {
			posS := h[5]
			val := h[6]
			pos, err := strconv.ParseInt(string(posS), 10, 64)
			if err != nil {
				//fmt.Printf("not a number %q\n", string(posS))
			} else if pos >= 0 && pos < 8 {
				if _, ok := found[int(pos)]; !ok {
					pwd[pos] = val
					found[int(pos)] = true
					fmt.Printf("idx %d: assign: %q to %d. pwd = %q\n", idx, string(val), pos, string(pwd))
					if isComplete() {
						return string(pwd)
					}
				}
			}
		}
		idx++
	}
}

func hash(doorID string, idx int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", doorID, idx))))
}
