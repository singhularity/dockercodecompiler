package sanitizers

import "github.com/dockercodecompiler/compiler/sandbox/staging/sanitizers/javasanitizer"

func Sanitize(language string, code string) string {
	if language == "java" {
		code = javasanitizer.Sanitize(code)
	}

	return code
}
