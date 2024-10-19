package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("hello"), 5, nil},
		{[]byte("Привет"), 6, nil},
		{[]byte("你好"), 2, nil},
		{[]byte("😀"), 1, nil},
		{[]byte(""), 0, nil},
		{[]byte{0xff, 0xf0, 0x28, 0x8c, 0xc3, 0x28}, 0, ErrInvalidUTF8},
		{[]byte{0xe0, 0xa0, 0x28}, 0, ErrInvalidUTF8},
	}

	for _, test := range tests {
		result, err := GetUTFLength(test.input)
		if result != test.expected || !errors.Is(err, test.err) {
			t.Errorf("GetUTFLength(%q) = (%d, %v); expected (%d, %v)", test.input, result, err, test.expected, test.err)
		}
	}
}

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
