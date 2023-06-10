package rss

func StripHTMLTags(s string) string {
	var result string
	var withinTag bool

	for _, c := range s {
		if c == '<' {
			withinTag = true
		} else if c == '>' {
			withinTag = false
		} else if !withinTag {
			result += string(c)
		}
	}

	return result
}
