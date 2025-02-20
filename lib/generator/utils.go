package generator

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"regexp"
)

var dirFilter b.FilterFunc = func(file b.BlogFile) bool {
	if file.GetKind() != b.Dir {
		return false
	}

	pattern := fmt.Sprintf("%s|%s", Public, Posts)
	re := regexp.MustCompile(pattern)

	if !re.MatchString(file.GetPath()) {
		return true
	} else {
		return false
	}
}

var postsFilter b.FilterFunc = func(file b.BlogFile) bool {
	if file.GetKind() != b.HTML {
		return false
	}

	pattern := fmt.Sprintf("(^|/)%s/%s(/|$)", Public, Posts)
	re := regexp.MustCompile(pattern)

	if re.MatchString(file.GetPath()) {
		return true
	} else {
		return false
	}
}
