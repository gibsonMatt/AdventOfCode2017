//So this worked for day 1, but instead of comparing the last int to the first here, I just did it manually. 
package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
)

func main() {
	dat, _ := ioutil.ReadFile("/home/mgibson/go/src/github.com/gibsonMatt/AdventOfCode2017/Day1/day1_input.txt")

	sum := 0
	p := byte('!')
	i := 0
	firstByte := byte('A')
	for _, c := range dat {
		if p == '!'{
			p = c
			firstByte = c
			continue
		}


		s, _ := strconv.Atoi(string(c))
		fmt.Println(s)
		if c == p{
			sum  = sum + s
		}
		p = c
		i++
	}
	fmt.Println(string(firstByte))
	fmt.Println(sum)
}

