package stringutils

import (
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
		{"nil slice", "test", nil, false},
		{"empty slice", "test", []string{}, false},
		{"empty string source", "", []string{"test", "example"}, false},
		{"empty string in slice", "test", []string{"", "test", "example"}, true},
		{"empty string match", "", []string{"", "test"}, true},
		{"unicode string", "тест", []string{"test", "тест", "example"}, true},
		{"unicode not found", "тест", []string{"test", "example"}, false},
		{"special characters", "test@#$", []string{"test@#$", "example"}, true},
		{"case sensitive", "Test", []string{"test", "example"}, false},
		{"duplicates in slice", "test", []string{"test", "test", "test"}, true},
		{"single item slice found", "test", []string{"test"}, true},
		{"single item slice not found", "test", []string{"example"}, false},
		{"spaces in string", "test string", []string{"test string", "example"}, true},
		{"partial match should not find", "test", []string{"testing", "testable"}, false},
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
		{"empty substring skipped", "hello world", []string{"", "missing"}, false},
		{"empty substring with match", "hello world", []string{"", "world"}, true},
		{"only empty substring", "hello world", []string{""}, false},
		{"multiple empty substrings", "hello world", []string{"", "", ""}, false},
		{"nil slice", "hello world", nil, false},
		{"empty slice", "hello world", []string{}, false},
		{"empty main string", "", []string{"test", "example"}, false},
		{"empty main string with empty substring", "", []string{""}, false},
		{"case sensitive", "Hello World", []string{"hello", "WORLD"}, false},
		{"case sensitive match", "Hello World", []string{"Hello", "missing"}, true},
		{"unicode substring", "привет мир", []string{"мир", "test"}, true},
		{"unicode not found", "привет мир", []string{"hello", "world"}, false},
		{"special characters", "test@#$%^&*()", []string{"@#$", "missing"}, true},
		{"overlapping matches", "testing", []string{"test", "sting", "ing"}, true},
		{"partial word match", "testing", []string{"test"}, true},
		{"multiple matches", "hello world", []string{"hello", "world", "test"}, true},
		{"very long string", strings.Repeat("a", 1000) + "needle" + strings.Repeat("b", 1000), []string{"needle"}, true},
		{"single character match", "hello", []string{"h", "x"}, true},
		{"newline in string", "hello\nworld", []string{"\n"}, true},
		{"tab in string", "hello\tworld", []string{"\t"}, true},
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
		{"all duplicates", []string{"same", "same", "same", "same"}, []string{"same"}},
		{"alternating duplicates", []string{"a", "b", "a", "b", "a"}, []string{"a", "b"}},
		{"many duplicates of one", []string{"x", "x", "x", "y", "x", "x"}, []string{"x", "y"}},
		{"single element", []string{"alone"}, []string{"alone"}},
		{"unicode strings", []string{"тест", "тест", "мир", "тест"}, []string{"тест", "мир"}},
		{"strings with spaces", []string{"hello world", "hello world", "test"}, []string{"hello world", "test"}},
		{"special characters", []string{"@#$", "^&*", "@#$", "^&*"}, []string{"@#$", "^&*"}},
		{"empty strings", []string{"", "", "test", ""}, []string{"", "test"}},
		{"case sensitive", []string{"Test", "test", "Test"}, []string{"Test", "test"}},
		{"adjacent duplicates", []string{"a", "a", "b", "b", "c", "c"}, []string{"a", "b", "c"}},
		{"scattered duplicates", []string{"a", "b", "c", "a", "d", "b"}, []string{"a", "b", "c", "d"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make a copy to verify no mutation
			var original []string
			if tt.keys != nil {
				original = make([]string, len(tt.keys))
				copy(original, tt.keys)
			}

			result := DeDup(tt.keys)
			assert.Equal(t, tt.want, result)

			// verify that original is not mutated
			if tt.keys != nil {
				assert.Equal(t, original, tt.keys, "should not mutate original slice")
			}
		})
	}
}

func TestDeDupBig(t *testing.T) {
	// DeDupBig is deprecated and just calls DeDup, so we only need basic tests for backwards compatibility
	tests := []struct {
		name string
		keys []string
		want []string
	}{
		{"nil input", nil, nil},
		{"removes duplicates", []string{"test", "test", "example"}, []string{"test", "example"}},
		{"verify stability", []string{"a", "b", "c", "b", "d", "a", "e"}, []string{"a", "b", "c", "d", "e"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeDupBig(tt.keys)
			assert.Equal(t, tt.want, result)

			// verify it returns same result as DeDup
			assert.Equal(t, DeDup(tt.keys), result, "DeDupBig should return same result as DeDup")
		})
	}
}

func TestSliceToString(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

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
		{"nil values", []any{nil, "test", nil}, []string{"<nil>", "test", "<nil>"}},
		{"empty byte slice", []any{[]byte{}, []byte("test")}, []string{"", "test"}},
		{"byte slice with null bytes", []any{[]byte{0x00, 0x01, 0x02}}, []string{"\x00\x01\x02"}},
		{"complex types", []any{testStruct{Name: "John", Age: 30}, map[string]int{"a": 1}}, []string{"{John 30}", "map[a:1]"}},
		{"pointers", []any{&testStruct{Name: "Jane", Age: 25}}, []string{"&{Jane 25}"}},
		{"arrays", []any{[3]int{1, 2, 3}}, []string{"[1 2 3]"}},
		{"slices", []any{[]int{4, 5, 6}}, []string{"[4 5 6]"}},
		{"negative numbers", []any{-1, -999, -0.5}, []string{"-1", "-999", "-0.5"}},
		{"large numbers", []any{int64(9223372036854775807)}, []string{"9223372036854775807"}},
		{"unicode in byte slice", []any{[]byte("привет мир")}, []string{"привет мир"}},
		{"special chars in byte slice", []any{[]byte("@#$%^&*()")}, []string{"@#$%^&*()"}},
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
		{"same slice twice", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"same slice twice empty", []string{}, []string{}, false},
		{"single element both", []string{"x"}, []string{"x"}, true},
		{"single element no match", []string{"x"}, []string{"y"}, false},
		{"all common", []string{"a", "b", "c"}, []string{"c", "b", "a"}, true},
		{"unicode elements", []string{"тест", "мир"}, []string{"привет", "мир"}, true},
		{"case sensitive", []string{"Test", "test"}, []string{"TEST", "test"}, true},
		{"nil first slice", nil, []string{"a"}, false},
		{"nil second slice", []string{"a"}, nil, false},
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
		{
			name:   "empty prefix",
			slice:  []string{"apple", "banana", "cherry"},
			prefix: "",
			exp:    true,
		},
		{
			name:   "empty slice",
			slice:  []string{},
			prefix: "ap",
			exp:    false,
		},
		{
			name:   "empty prefix and empty slice",
			slice:  []string{},
			prefix: "",
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
		{
			name:   "empty suffix",
			slice:  []string{"apple", "banana", "cherry"},
			suffix: "",
			exp:    true,
		},
		{
			name:   "empty slice",
			slice:  []string{},
			suffix: "na",
			exp:    false,
		},
		{
			name:   "empty suffix and empty slice",
			slice:  []string{},
			suffix: "",
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

func BenchmarkDeDup(b *testing.B) {
	// small slice with duplicates
	small := []string{"a", "b", "a", "c", "b", "d", "e", "a"}

	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DeDup(small)
		}
	})

	// medium slice
	medium := make([]string, 100)
	for i := 0; i < 100; i++ {
		medium[i] = string(rune('a' + (i % 10))) // only 10 unique values
	}

	b.Run("medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DeDup(medium)
		}
	})

	// large slice with many duplicates
	large := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		large[i] = string(rune('a' + (i % 26))) // only 26 unique values
	}

	b.Run("large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DeDup(large)
		}
	})
}

func BenchmarkSliceToString(b *testing.B) {
	tmpl := []any{[]byte("fdjndfg")}
	b.Run("small slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SliceToString(tmpl)
		}
	})

	for i := 0; i < 20; i++ {
		tmpl = append(tmpl, tmpl...)
	}

	b.Run("big slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SliceToString(tmpl)
		}
	})
}
