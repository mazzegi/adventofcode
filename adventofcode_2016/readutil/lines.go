package readutil

import "strings"

func ReadLines(s string) []string {
	var lines []string
	sl := strings.Split(s, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}
	return lines
}

func ReadString(s string) string {
	return strings.Join(ReadLines(s), "")
}
