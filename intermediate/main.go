package main

import "bytes"

func messageFromBinaryCode(msg string) string {
	return ""
}

func pow(o, p int) int {
	tmp := 1
	for i := 0; i < p; i++ {
		tmp *= o
	}
	return tmp
}

func solution(code string) string {
	strs := make([]string, 0, int(len(code)/8))
	tmp := make([]int, 0, int(len(code)/8))
	var b bytes.Buffer
	for i := range code {
		b.WriteByte(code[i])
		if (i+1)%8 == 0 {
			strs = append(strs, b.String())
			b.Reset()
		}
	}
	for _, s := range strs {
		byt := 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '1' {
				byt += pow(2, 7-i)
			}
		}
		tmp = append(tmp, byt)
	}
	res := make([]byte, 0, len(tmp))
	for _, current := range tmp {
		res = append(res, byte(current))
	}
	return string(res)
}

func main() {}
