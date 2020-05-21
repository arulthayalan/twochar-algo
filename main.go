package main

import (
	"fmt"
	"strings"
)

func main() {

	var l int
	var input string
	fmt.Scanf("%d", &l)
	fmt.Scanf("%s", &input)

	fmt.Printf("%d\n", alternative(input))

}

func alternative(s string) int {
	var ret string
	//get unique characters
	uch := unichar(s)

	//all possible combination of chars
	var p []string
	for i := 0; i < len(uch)-1; i++ {
		for j := i + 1; j < len(uch); j++ {
			t := make([]rune, 2)
			t = append(t, uch[i], uch[j])
			p = append(p, string(t))
		}
	}

	//Iterate through all combination
	//remove other chars except combination
	//check contains alternative
	for _, cb := range p {
		alt := remove(s, cb)
		if containsalertnative(alt) && len(alt) > len(ret) {
			ret = alt
		}
	}
	return len(ret)
}

func remove(in string, ex string) string {
	for _, ch := range in {
		if strings.IndexRune(ex, ch) < 0 {
			in = strings.ReplaceAll(in, string(ch), "")
		}
	}
	return in
}

func containsalertnative(in string) bool {
	var ret bool = true

	for idx := 0; idx < len(in); idx++ {
		if (idx-1) >= 0 && in[idx] == in[idx-1] {
			ret = false
			break
		}
		if (idx - 2) <= -1 {
			continue
		}
		if in[idx] != in[idx-2] {
			ret = false
			break
		}
	}
	return ret
}

func unichar(in string) []rune {
	var temp = make(map[rune]bool)

	for _, ch := range in {
		if _, ok := temp[ch]; !ok {
			temp[ch] = true
		}
	}
	ret := make([]rune, len(temp))
	i := 0
	for k := range temp {
		ret[i] = k
		i++
	}
	return ret
}
