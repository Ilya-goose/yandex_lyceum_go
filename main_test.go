package main

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("Hello, World!"), 13, nil},            // Valid UTF-8	{[]byte("こんにちは"), 5, nil},                     // Valid UTF-8 (Japanese)
		{[]byte("Привет"), 6, nil},                    // Valid UTF-8 (Russian)
		{[]byte("😊"), 1, nil},                         // Valid UTF-8 (Emoji)
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8}, // Invalid UTF-8	{[]byte{0x80}, 0, ErrInvalidUTF8},                 // Invalid UTF-8 (overlong encoding)
		{[]byte{0xc3, 0x28}, 0, ErrInvalidUTF8},       // Invalid UTF-8 (invalid continuation byte)
	}

	for _, test := range tests {
		length, err := GetUTFLength(test.input)
		if length != test.expected || !errors.Is(err, test.err) {
			t.Errorf("GetUTFLength(%q) = (%d, %v), expected (%d, %v)", test.input, length, err, test.expected, test.err)
		}
	}
}

// еще новенький комит
