package utils

import "strings"

func RemovePrefixSuffix(str, ps string) string {
	str = strings.TrimPrefix(str, ps)
	str = strings.TrimSuffix(str, ps)
	return str
}
