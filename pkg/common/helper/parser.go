package helper

import (
	"strconv"
	"strings"
)

func ParseUint16(s string) (u uint16, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint16(tmp), nil
}

func ParseUint(s string) (u uint, err error) {
	var tmp uint64
	tmp, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(tmp), nil
}

func ParseInt(s string) (i int, err error) {
	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(tmp), nil
}

func TrimSlashPrefix(s string) string {
	if len(s) == 0 {
		return s
	}

	if (s[0] == '\\') || (s[0] == '/') {
		return s[1:]
	}

	return s
}

func ParseCSV(s string) (values []string) {
	parts := strings.Split(s, ",")

	values = make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) > 0 {
			values = append(values, part)
		}
	}

	return values
}

func ToUpperCase(lc []string) (uc []string) {
	uc = make([]string, len(lc))
	for i, lcs := range lc {
		uc[i] = strings.ToUpper(lcs)
	}
	return uc
}
