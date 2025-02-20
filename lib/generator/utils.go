package generator

import (
	b "homestead/lib/blogfsys"
	"regexp"
)

func filterInclude(pattern string, files []b.BlogFile) (filtered []b.BlogFile) {
	re := regexp.MustCompile(pattern)

	for _, file := range files {
		if re.MatchString(file.GetPath()) {
			filtered = append(filtered, file)
		}
	}

	return filtered
}

func filterExclude(pattern string, files []b.BlogFile) (filtered []b.BlogFile) {
	re := regexp.MustCompile(pattern)

	for _, file := range files {
		if !re.MatchString(file.GetPath()) {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
