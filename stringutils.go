// Package stringutils provides utilities for working with strings.
package stringutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// Contains string in slice
func Contains(src string, inSlice []string) bool {
	for _, a := range inSlice {
		if a == src {
			return true
		}
	}
	return false
}

// ContainsAnySubstring checks if string contains any of provided substring
func ContainsAnySubstring(s string, subStrings []string) bool {
	for _, mx := range subStrings {
		if mx == "" {
			continue // skip empty substrings
		}
		if strings.Contains(s, mx) {
			return true
		}
	}
	return false
}

// DeDup remove duplicates from slice.
// This function is stable - it preserves the order of first occurrences.
func DeDup(keys []string) []string {
	if len(keys) == 0 {
		return nil
	}
	result := make([]string, 0, len(keys))
	visited := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		if _, found := visited[k]; !found {
			visited[k] = struct{}{}
			result = append(result, k)
		}
	}
	return result
}

// DeDupBig remove duplicates from slice.
// Deprecated: Use DeDup instead. This function now just calls DeDup for backwards compatibility.
func DeDupBig(keys []string) []string {
	return DeDup(keys)
}

// SliceToString converts slice of any to slice of string
func SliceToString(s []any) []string {
	if len(s) == 0 {
		return nil
	}
	strSlice := make([]string, len(s))
	for i, v := range s {
		if vb, ok := v.([]byte); ok {
			strSlice[i] = string(vb) // safe conversion
			continue
		}
		strSlice[i] = fmt.Sprintf("%v", v)
	}
	return strSlice
}

// HasCommonElement checks if any element of the second slice is in the first slice
func HasCommonElement(a, b []string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	// build set from smaller slice for better performance
	if len(a) > len(b) {
		a, b = b, a
	}
	set := make(map[string]struct{}, len(a))
	for _, x := range a {
		set[x] = struct{}{}
	}
	for _, y := range b {
		if _, ok := set[y]; ok {
			return true
		}
	}
	return false
}

// HasPrefixSlice checks if any string in the slice starts with the given prefix
func HasPrefixSlice(prefix string, slice []string) bool {
	for _, v := range slice {
		if strings.HasPrefix(v, prefix) {
			return true
		}
	}
	return false
}

// HasSuffixSlice checks if any string in the slice ends with the given suffix
func HasSuffixSlice(suffix string, slice []string) bool {
	for _, v := range slice {
		if strings.HasSuffix(v, suffix) {
			return true
		}
	}
	return false
}

// Truncate cuts string to the given length (in runes) and adds ellipsis if it was truncated
// if maxLen is less than 4 (3 chars for ellipsis + 1 rune from string), returns empty string
func Truncate(s string, maxLen int) string {
	if maxLen < 4 {
		return ""
	}

	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}

	return string(runes[:maxLen-3]) + "..."
}

// TruncateWords cuts string to the given number of words and adds ellipsis if it was truncated
// returns empty string if maxWords is 0
func TruncateWords(s string, maxWords int) string {
	if maxWords == 0 {
		return ""
	}

	words := strings.Fields(s)
	if len(words) <= maxWords {
		return s
	}

	return strings.Join(words[:maxWords], " ") + "..."
}

// RandomWord generates pronounceable random word with length between minLen and maxLen
func RandomWord(minLen, maxLen int) string {
	if minLen < 2 {
		minLen = 2
	}
	if maxLen < minLen {
		maxLen = minLen
	}

	vowels := []rune("aeiou")
	consonants := []rune("bcdfghjklmnpqrstvwxyz")

	// make a random length between min and max
	n, err := rand.Int(rand.Reader, big.NewInt(int64(maxLen-minLen+1)))
	length := minLen
	if err == nil {
		length += int(n.Int64())
	}

	var result strings.Builder
	// decide to start with vowel or consonant
	n, _ = rand.Int(rand.Reader, big.NewInt(2))
	startWithVowel := n.Int64() == 0

	for i := 0; i < length; i++ {
		isVowel := (i%2 == 0) == startWithVowel
		if isVowel {
			n, _ = rand.Int(rand.Reader, big.NewInt(int64(len(vowels))))
			result.WriteRune(vowels[n.Int64()])
		} else {
			n, _ = rand.Int(rand.Reader, big.NewInt(int64(len(consonants))))
			result.WriteRune(consonants[n.Int64()])
		}
	}

	return result.String()
}
