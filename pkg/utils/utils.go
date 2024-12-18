package utils

import "strings"

// InEnums ...
func InEnums(str string, enums []string) bool {
	for _, enum := range enums {
		if enum == str {
			return true
		}
	}
	return false
}

func SplitLink(link string) string {
	return strings.Split(link, "/")[5]
}
