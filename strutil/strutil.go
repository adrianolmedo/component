package strutil

import "strings"

// EqualNoStrict verifies if two strings are equal.
// It's case insensitive.
// It ignores space like characters.
// It ignores accents.
func EqualNoStrict(str1 string, str2 string) bool {
	str1 = equalNoStrictTransformer(str1)
	str2 = equalNoStrictTransformer(str2)
	return str1 == str2
}

func equalNoStrictTransformer(str string) string {
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "á", "a")
	str = strings.ReplaceAll(str, "é", "e")
	str = strings.ReplaceAll(str, "í", "i")
	str = strings.ReplaceAll(str, "ó", "o")
	str = strings.ReplaceAll(str, "ú", "u")
	str = strings.ReplaceAll(str, "ä", "a")
	str = strings.ReplaceAll(str, "ë", "e")
	str = strings.ReplaceAll(str, "ï", "i")
	str = strings.ReplaceAll(str, "ö", "o")
	str = strings.ReplaceAll(str, "ü", "u")
	str = strings.ReplaceAll(str, "à", "a")
	str = strings.ReplaceAll(str, "è", "e")
	str = strings.ReplaceAll(str, "ì", "i")
	str = strings.ReplaceAll(str, "ò", "o")
	str = strings.ReplaceAll(str, "ù", "u")
	return str
}
