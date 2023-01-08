package valid_parentheses

var pMap = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid_Myans(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var st []rune
	for _, r := range s {
		p, ok := pMap[r]
		if ok {
			if len(st) != 0 {
				if st[len(st)-1] != p {
					return false
				}
				st = st[:len(st)-1]
				continue
			}
		}
		st = append(st, r)
	}
	return len(st) == 0
}
