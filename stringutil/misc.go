package stringutil

func Reverse(s string) string {
	var rs []rune
	for i := len(s) - 1; i >= 0; i-- {
		rs = append(rs, rune(s[i]))
	}
	return string(rs)
}
