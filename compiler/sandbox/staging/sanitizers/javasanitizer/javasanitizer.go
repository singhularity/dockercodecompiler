package javasanitizer

import (
	"regexp"
	"strings"
)

func Sanitize(code string) string {
	return escapeQuotes(replacePublicClass(removePackage(code)))
}

func replacePublicClass(code string) string {
	return strings.Replace(code, "public class", "class", -1)
}

func removePackage(code string) string {
	re := regexp.MustCompile("^package\\s+([a-zA_Z_][\\.\\w]*);")
	return re.ReplaceAllString(code, "")
}

func escapeQuotes(code string) string {
	return strings.Replace(strings.Replace(code, "’", "\\'", -1), "‘", "\\'", -1)
}
