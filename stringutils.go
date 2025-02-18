package stringutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"unsafe"
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
		if strings.Contains(s, mx) {
			return true
		}
	}
	return false
}

// DeDup remove duplicates from slice. optimized for performance, good for short slices only!
func DeDup(keys []string) []string {
	if len(keys) == 0 {
		return nil
	}
	l := len(keys) - 1
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if keys[i] == keys[j] {
				keys[j] = keys[l]
				keys = keys[0:l]
				l--
				j--
			}
		}
	}
	return keys
}

// DeDupBig remove duplicates from slice. Should be used instead of DeDup for large slices
func DeDupBig(keys []string) (result []string) {
	if len(keys) == 0 {
		return nil
	}
	result = make([]string, 0, len(keys))
	visited := map[string]bool{}
	for _, k := range keys {
		if _, found := visited[k]; !found {
			visited[k] = found
			result = append(result, k)
		}
	}
	return result
}

// SliceToString converts slice of any to slice of string
func SliceToString(s []any) []string {
	if len(s) == 0 {
		return nil
	}
	strSlice := make([]string, len(s))
	for i, v := range s {
		if vb, ok := v.([]byte); ok {
			strSlice[i] = bytesToString(vb)
			continue
		}
		strSlice[i] = fmt.Sprintf("%v", v)
	}
	return strSlice
}

// nolint
func bytesToString(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}

// HasCommonElement checks if any element of the second slice is in the first slice
func HasCommonElement(a, b []string) bool {
	for _, second := range b {
		for _, first := range a {
			if first == second {
				return true
			}
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
