package stringutils

import (
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
		{"converts number to string", []any{1, 2, 3}, []string{"1", "2", "3"}},
		{"converts mixed slice to string", []any{1, "aaa", true, 0.55}, []string{"1", "aaa", "true", "0.55"}},
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
