package levenshtein

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Option allows to compute operations on the strings provided to Distance, before the distance computation.
// It will be call for both strings.
type Option func(s string) string

// IgnoreDiacritics allows to ignore diacritics (accents) when computing strings distance.
var IgnoreDiacritics Option = func(s string) string {
	t := transform.Chain(
		norm.NFD,
		transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }),
		norm.NFC,
	)
	result, _, _ := transform.String(t, s)
	return result
}

// IgnoreCase allows to disable case sensitivity when computing strings distance.
var IgnoreCase Option = func(s string) string {
	return strings.ToLower(s)
}

// Distance computes the distance between two strings.
// 0 means the provided strings are identical.
// The higher the result is, the more the strings differ.
//
// Not that without the IgnoreDiacritics option, the distance between an accentuated character and a different
// character will be 2, as an accentuated character is encoded on 2 octets.
func Distance(s1, s2 string, opts ...Option) int {
	for _, o := range opts {
		s1 = o(s1)
		s2 = o(s2)
	}

	// Create matrix.
	d := make([][]int, len(s1)+1)
	for i := range d {
		d[i] = make([]int, len(s2)+1)
	}
	// Set first row.
	for i := range d {
		d[i][0] = i
	}
	// Set first column.
	for j := range d[0] {
		d[0][j] = j
	}
	// Magic.
	for j := 1; j <= len(s2); j++ {
		for i := 1; i <= len(s1); i++ {
			if s1[i-1] == s2[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s1)][len(s2)]
}
