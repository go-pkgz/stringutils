# stringutils [![Build Status](https://github.com/go-pkgz/stringutils/workflows/build/badge.svg)](https://github.com/go-pkgz/stringutils/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/go-pkgz/stringutils)](https://goreportcard.com/report/github.com/go-pkgz/stringutils) [![Coverage Status](https://coveralls.io/repos/github/go-pkgz/stringutils/badge.svg?branch=master)](https://coveralls.io/github/go-pkgz/stringutils?branch=master)

Package `stringutils` provides useful string operations.

## Details

String manipulation:
- **Contains**: checks if slice contains a string.
- **ContainsAnySubstring**: checks if string contains any of provided substring.
- **DeDup**: removes duplicates from slice of strings, optimized for performance, good for short slices only.
- **DeDupBig**: removes duplicates from slice. Should be used instead of `DeDup` for large slices.
- **SliceToString**: converts slice of `any` to a slice of strings.
- **HasCommonElement**: checks if any element of the second slice is in the first slice.
- **HasPrefixSlice**: checks if any string in the slice starts with the given prefix.
- **HasSuffixSlice**: checks if any string in the slice ends with the given suffix.
- **Truncate**: cuts string to the given length and adds ellipsis if it was truncated.
- **TruncateWords**: cuts string to the given number of words and adds ellipsis if it was truncated.
- **RandomWord**: generates pronounceable random word with given min/max length.

## Install and update

`go get -u github.com/go-pkgz/stringutils`

## Usage examples