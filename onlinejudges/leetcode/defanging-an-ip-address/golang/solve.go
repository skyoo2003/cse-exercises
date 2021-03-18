// nolint
package defanginganipaddress

func defangIPaddr(address string) string {
	buf := make([]rune, 0)
	for _, ch := range address {
		if ch == '.' {
			buf = append(buf, '[', ch, ']')
		} else {
			buf = append(buf, ch)
		}
	}
	return string(buf)
}
