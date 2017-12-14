//So this worked for day 1, but instead of comparing the last int to the first here, I just did it manually. 
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func max(x []int) int {
	maxInteger := 0
	for _, i := range x{
		if i > maxInteger{
			maxInteger = i
		}
	}
	return maxInteger
}

func min(x []int) int {
        minInteger := 999999
        for _, i := range x{
                if i < minInteger{
                        minInteger = i
                }
        }
        return minInteger
}



func main() {
	dat, _ := os.Open("/home/mgibson/go/src/github.com/gibsonMatt/AdventOfCode2017/Day2/day2_input.txt")

	defer dat.Close()
	fileScanner := bufio.NewScanner(dat)

	sum := 0
	for fileScanner.Scan(){
		words := strings.Fields(fileScanner.Text())
		var integers = []int{}
		for _, word := range words{
			i, _ := strconv.Atoi(word)
			integers = append(integers, i)
		}

		sum = sum + (max(integers)-min(integers))
	}
	fmt.Println(sum)
}
