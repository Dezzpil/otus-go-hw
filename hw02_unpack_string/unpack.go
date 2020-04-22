package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packed string) (string, error) {
	var escaping bool
	var rPrev rune
	var unpacked strings.Builder
	for len(packed) > 0 {
		r, size := utf8.DecodeRuneInString(packed)
		packed = packed[size:]
		switch {
		case r == '\\' && !escaping:
			escaping = true
			continue
		case unicode.IsDigit(r) && !escaping:
			if rPrev == 0 {
				return "", ErrInvalidString
			}
			repeat, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}
			for i := 2; i <= repeat; i++ {
				unpacked.WriteRune(rPrev)
			}
			rPrev = 0
		default:
			rPrev = r
			unpacked.WriteRune(rPrev)
		}
		escaping = false
	}

	return unpacked.String(), nil
}
