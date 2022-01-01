package polymer

import "unicode"

func destroys(r1, r2 rune) bool {
	if unicode.ToLower(r1) != unicode.ToLower(r2) {
		return false
	}
	return unicode.IsLower(r1) != unicode.IsLower(r2)
}

func React(p string) string {
	changed := true
	for changed {
		p, changed = reactOnceZ(p)
	}
	return p
}

func reactOnceZ(p string) (string, bool) {
	if len(p) < 2 {
		return p, false
	}
	for i := 0; i < len(p)-1; i++ {
		if destroys(rune(p[i]), rune(p[i+1])) {
			p = p[:i] + p[i+2:]
			return p, true
		}
	}
	return p, false
}

func reactOnce(p string) (string, bool) {
	if len(p) < 2 {
		return p, false
	}
	var np string
	changed := false
	i := 0
	for i = 0; i < len(p)-1; i++ {
		if destroys(rune(p[i]), rune(p[i+1])) {
			changed = true
			i += 1
			continue
		} else {
			np += string(p[i])
		}
	}
	if i < len(p) {
		np += string(p[i])
	}
	return np, changed
}
