package helper

import (
	"strings"
)

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
