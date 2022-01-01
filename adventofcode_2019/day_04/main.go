package main

import "fmt"

func main() {
	start := []int{3, 6, 7, 4, 7, 9}
	end := []int{8, 9, 3, 6, 9, 8}

	atEnd := func() bool {
		for i := 0; i < len(end); i++ {
			if start[i] != end[i] {
				return false
			}
		}
		return true
	}

	inc := func() {
		for i := len(start) - 1; i >= 0; i-- {
			if start[i] < 9 {
				start[i]++
				return
			}
			start[i] = 0
		}
	}

	var cnt int
	for !atEnd() {
		if IsValidPwdV2(start) {
			cnt++
		}
		inc()
	}
	fmt.Printf("there are %d valid passwords\n", cnt)
	//495

	pwd := []int{1, 1, 2, 2, 3, 3}
	fmt.Printf("%v => %t\n", pwd, IsValidPwdV2(pwd))
	pwd = []int{1, 2, 3, 4, 4, 4}
	fmt.Printf("%v => %t\n", pwd, IsValidPwdV2(pwd))
	pwd = []int{1, 1, 1, 1, 2, 2}
	fmt.Printf("%v => %t\n", pwd, IsValidPwdV2(pwd))
}

func IsValidPwd(pwd []int) bool {
	var hasDouble bool
	for i := 0; i < len(pwd); i++ {
		if i > 0 && pwd[i] < pwd[i-1] {
			return false
		}
		if i > 0 && pwd[i] == pwd[i-1] {
			hasDouble = true
		}
	}
	if !hasDouble {
		return false
	}
	return true
}

func IsValidPwdV2(pwd []int) bool {
	var hasDouble bool
	rep := pwd[0]
	repCnt := 1
	for i := 1; i < len(pwd); i++ {
		if i > 0 && pwd[i] < pwd[i-1] {
			return false
		}
		if pwd[i] == rep {
			repCnt++
		} else {
			if repCnt == 2 {
				hasDouble = true
			}
			rep = pwd[i]
			repCnt = 1
		}
	}
	if repCnt == 2 {
		hasDouble = true
	}
	if !hasDouble {
		return false
	}
	return true
}
