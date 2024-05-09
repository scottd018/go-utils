package common

import "unicode/utf8"

func ReverseString(str string) string {
	size := len(str)

	buf := make([]byte, size)

	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(str[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}

	return string(buf)
}
