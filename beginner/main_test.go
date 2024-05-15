package main

import "testing"

func Test_isIPV4(t *testing.T) {
	for _, test := range []struct {
		description string
		addr        string
		expectIPV4  bool
	}{
		{
			description: "valid",
			addr:        "172.16.254.1",
			expectIPV4:  true,
		},
		{
			description: "invalid",
			addr:        "172.316.254.1",
		},
		{
			description: "missing byte",
			addr:        ".254.255.0",
		},
		{
			description: "invalid length",
			addr:        "1",
		},
		{
			description: "alphanumeric",
			addr:        "255.255.255.255abcdekjhf",
		},
		{
			description: "leading zero",
			addr:        "233.01.61.131",
		},
		{
			description: "this network",
			addr:        "0.0.0.0",
			expectIPV4:  true,
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			if solution(test.addr) != test.expectIPV4 {
				t.Fail()
			}
		})
	}
}
