package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	s := "Paste text below on this site: https://codebeautify.org/base64-to-image-converter\n"
	s += "data:image/png;base64,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"
	conv := convert(s)
	if err := ioutil.WriteFile("task5.bin", []byte(conv), 0666); err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("task5.bin")
	if err != nil {
		panic(err)
	}
	dec := decode(string(data))
	fmt.Println(dec)
}

func convert(s string) string {
	res := make([]byte, len(s))
	for i := range s {
		res[i] = (s[i] - byte(i%3)) & 0xFF
	}
	return string(res)
}

func decode(s string) string {
	res := make([]byte, len(s))
	for i := range s {
		switch i % 3 {
		case 0:
			res[i] = s[i]
		case 1:
			res[i] = s[i] + 1
		case 2:
			res[i] = s[i] + 2
		}
	}
	return string(res)
}
