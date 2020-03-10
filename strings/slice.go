package strings

func Slice(s string) string {
	if len(s) <= 10 {
		return s
	}
	return "..." + s[len(s)-10:]
}
