package helper

import "strconv"

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
