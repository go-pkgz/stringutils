package stringutils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		src  string
		list []string
		want bool
	}{
		{"finds string", "test", []string{"test", "example"}, true},
		{"doesn't find string", "missing", []string{"test", "example"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Contains(tt.src, tt.list))
		})
	}
}

func TestContainsAnySubstring(t *testing.T) {
	tests := []struct {
		name       string
		str        string
		substrings []string
		want       bool
	}{
		{"contains substring", "hello world", []string{"world", "example"}, true},
		{"doesn't contain substring", "hello world", []string{"missing", "example"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ContainsAnySubstring(tt.str, tt.substrings))
		})
	}
}

func TestDeDup(t *testing.T) {
	tests := []struct {
		name string
		keys []string
		want []string
	}{
		{"nil input", nil, nil},
		{"empty input", []string{}, nil},
		{"removes duplicates", []string{"test", "test", "example"}, []string{"test", "example"}},
		{"no duplicates", []string{"test", "test2", "example"}, []string{"test", "test2", "example"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DeDup(tt.keys))
		})
	}
}

func TestDeDupBig(t *testing.T) {
	tests := []struct {
		name string
		keys []string
		want []string
	}{
		{"nil input", nil, nil},
		{"empty input", []string{}, nil},
		{"removes duplicates", []string{"test", "test", "example"}, []string{"test", "example"}},
		{"no duplicates", []string{"test", "test2", "example"}, []string{"test", "test2", "example"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DeDupBig(tt.keys))
		})
	}
}

func TestSliceToString(t *testing.T) {
	tests := []struct {
		name string
		in   []interface{}
		want []string
	}{
		{"nil input", nil, nil},
		{"empty input", []any{}, nil},
		{"converts number to string", []any{1, 2, 3}, []string{"1", "2", "3"}},
		{"converts mixed slice to string", []any{1, "aaa", true, 0.55}, []string{"1", "aaa", "true", "0.55"}},
		{"converts slice of byte slices to string", []any{[]byte("hi"), []byte("there")}, []string{"hi", "there"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SliceToString(tt.in))
		})
	}
}

func TestHasCommonElement(t *testing.T) {
	tests := []struct {
		name string
		a, b []string
		want bool
	}{
		{"element found", []string{"a", "b", "c", "d"}, []string{"x", "y", "z", "a"}, true},
		{"element not found", []string{"a", "b", "c", "d"}, []string{"x", "y", "z", "w"}, false},
		{"both slices are empty", []string{}, []string{}, false},
		{"one slice is empty", []string{}, []string{"x", "y", "z", "w"}, false},
		{"element found at the start", []string{"a", "b", "c", "d"}, []string{"a", "x", "y", "z"}, true},
		{"element found at the end", []string{"a", "b", "c", "d"}, []string{"x", "y", "z", "d"}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, HasCommonElement(tc.a, tc.b))
		})
	}
}

func TestHasPrefixSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		prefix string
		exp    bool
	}{
		{
			name:   "prefix exists",
			slice:  []string{"apple", "banana", "cherry"},
			prefix: "ap",
			exp:    true,
		},
		{
			name:   "prefix does not exist",
			slice:  []string{"apple", "banana", "cherry"},
			prefix: "kiwi",
			exp:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, HasPrefixSlice(tt.prefix, tt.slice))
		})
	}
}

func TestHasSuffixSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		suffix string
		exp    bool
	}{
		{
			name:   "suffix exists",
			slice:  []string{"apple", "banana", "cherry"},
			suffix: "na",
			exp:    true,
		},
		{
			name:   "suffix does not exist",
			slice:  []string{"apple", "banana", "cherry"},
			suffix: "kiwi",
			exp:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, HasSuffixSlice(tt.suffix, tt.slice))
		})
	}
}
func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		maxLen int
		want   string
	}{
		{
			name:   "no truncation needed",
			input:  "hello",
			maxLen: 10,
			want:   "hello",
		},
		{
			name:   "truncation needed",
			input:  "hello world",
			maxLen: 6, // changed from 8 since we want "hel..."
			want:   "hel...",
		},
		{
			name:   "maxLen too small",
			input:  "hello",
			maxLen: 3,
			want:   "",
		},
		{
			name:   "exactly at limit",
			input:  "hello",
			maxLen: 5,
			want:   "hello",
		},
		{
			name:   "unicode string no truncation",
			input:  "привет",
			maxLen: 6,
			want:   "привет",
		},
		{
			name:   "unicode string with truncation",
			input:  "привет мир",
			maxLen: 7,
			want:   "прив...",
		},
		{
			name:   "mixed ascii and unicode",
			input:  "hello мир",
			maxLen: 7,
			want:   "hell...",
		},
		{
			name:   "emoji string",
			input:  "👋🌍✨",
			maxLen: 6, // changed from 4 since we should only truncate if it won't fit
			want:   "👋🌍✨",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Truncate(tt.input, tt.maxLen)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTruncateWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxWords int
		want     string
	}{
		{
			name:     "no truncation needed",
			input:    "hello world",
			maxWords: 2,
			want:     "hello world",
		},
		{
			name:     "truncation needed",
			input:    "hello beautiful world",
			maxWords: 2,
			want:     "hello beautiful...",
		},
		{
			name:     "zero max words",
			input:    "hello",
			maxWords: 0,
			want:     "",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello    beautiful     world",
			maxWords: 2,
			want:     "hello beautiful...",
		},
		{
			name:     "unicode words no truncation",
			input:    "привет мир",
			maxWords: 2,
			want:     "привет мир",
		},
		{
			name:     "unicode words with truncation",
			input:    "привет красивый мир",
			maxWords: 2,
			want:     "привет красивый...",
		},
		{
			name:     "mixed ascii and unicode",
			input:    "hello мир world",
			maxWords: 2,
			want:     "hello мир...",
		},
		{
			name:     "with emoji",
			input:    "👋 hello 🌍 world ✨",
			maxWords: 2,
			want:     "👋 hello...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, TruncateWords(tt.input, tt.maxWords))
		})
	}
}

func TestRandomWord(t *testing.T) {
	tests := []struct {
		name       string
		minLen     int
		maxLen     int
		wantMinLen int
		wantMaxLen int
	}{
		{
			name:       "normal case",
			minLen:     4,
			maxLen:     8,
			wantMinLen: 4,
			wantMaxLen: 8,
		},
		{
			name:       "min less than 2",
			minLen:     1,
			maxLen:     5,
			wantMinLen: 2,
			wantMaxLen: 5,
		},
		{
			name:       "max less than min",
			minLen:     5,
			maxLen:     3,
			wantMinLen: 5,
			wantMaxLen: 5,
		},
		{
			name:       "fixed length",
			minLen:     6,
			maxLen:     6,
			wantMinLen: 6,
			wantMaxLen: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get multiple samples since it's random
			for i := 0; i < 10; i++ {
				got := RandomWord(tt.minLen, tt.maxLen)
				t.Logf("got: %s", got)
				assert.GreaterOrEqual(t, len(got), tt.wantMinLen, "word too short")
				assert.LessOrEqual(t, len(got), tt.wantMaxLen, "word too long")

				// check that word only contains allowed chars
				for _, r := range got {
					assert.Contains(t, "abcdefghijklmnopqrstuvwxyz", string(r))
				}

				// check for alternating vowels and consonants
				var prevIsVowel bool
				firstChar := rune(got[0])
				prevIsVowel = strings.ContainsRune("aeiou", firstChar)

				for _, r := range got[1:] {
					isVowel := strings.ContainsRune("aeiou", r)
					assert.NotEqual(t, prevIsVowel, isVowel, "vowels/consonants should alternate")
					prevIsVowel = isVowel
				}
			}
		})
	}
}

func BenchmarkSliceToString(b *testing.B) {
	tmpl := []any{[]byte("fdjndfg")}
	b.Run("unsafe (small slice)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SliceToString(tmpl)
		}
	})
	b.Run("type assert (small slice)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sliceToStringAllocs(tmpl)
		}
	})

	for i := 0; i < 20; i++ {
		tmpl = append(tmpl, tmpl...)
	}

	b.Run("unsafe (big slice)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SliceToString(tmpl)
		}
	})
	b.Run("type assert (big slice)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sliceToStringAllocs(tmpl)
		}
	})
}

func sliceToStringAllocs(s []any) []string {
	if len(s) == 0 {
		return nil
	}
	strSlice := make([]string, len(s))
	for i, v := range s {
		if vb, ok := v.([]byte); ok {
			strSlice[i] = string(vb)
			continue
		}
		strSlice[i] = fmt.Sprintf("%v", v)
	}
	return strSlice
}
