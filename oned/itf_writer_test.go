package oned

import (
	"testing"

	"github.com/kaxap/gozxing"
)

func TestITFEncoder_encode(t *testing.T) {
	enc := itfEncoder{}

	failtests := []string{
		"012",
		"0123456789012345678901234567890123456789012345678901234567890123456789012345678901",
		"abcdef",
	}
	for _, test := range failtests {
		_, e := enc.encode(test)
		if e == nil {
			t.Fatalf("encode(%v) must be error", test)
		}
	}
}

func TestITFWriter(t *testing.T) {
	writer := NewITFWriter()

	tests := []struct {
		content string
		wants   string
	}{
		{
			"0123456789",
			"000001010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"1110100000",
		},
		{
			"01234567890123456789012345678901234567890123456789012345678901234567890123456789",
			"000001010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"100010111011101000100011100010101110100010111000101110101110111010001000111010001011100010" +
				"1110100000",
		},
	}
	for _, test := range tests {
		testEncode(t, writer, gozxing.BarcodeFormat_ITF, test.content, test.wants)
	}
}
