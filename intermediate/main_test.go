package main

import (
	"fmt"
	"testing"
)

func Test_isIPV4(t *testing.T) {
	for i, test := range []struct {
		input   string
		decoded string
	}{
		{
			input:   "010010000110010101101100011011000110111100100001",
			decoded: "Hello!",
		},
		{
			input:   "01001101011000010111100100100000011101000110100001100101001000000100011001101111011100100110001101100101001000000110001001100101001000000111011101101001011101000110100000100000011110010110111101110101",
			decoded: "May the Force be with you",
		},
		{
			input:   "010011100110010101110110011001010111001000100000011100100110010101100111011100100110010101110100001000000110000101101110011110010111010001101000011010010110111001100111001000000111010001101000011000010111010000100000011011010110000101100100011001010010000001111001011011110111010100100000011100110110110101101001011011000110010100101110",
			decoded: "Never regret anything that made you smile.",
		},
		{
			input:   "010001000110100101100101001000000111011101101001011101000110100000100000011011010110010101101101011011110111001001101001011001010111001100101100001000000110111001101111011101000010000001100100011100100110010101100001011011010111001100101110",
			decoded: "Die with memories, not dreams.",
		},
	} {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			if solution(test.input) != test.decoded {
				t.Fail()
			}
		})
	}
}
