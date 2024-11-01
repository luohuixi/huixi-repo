package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var target int

func sum(n []int, t int) int {
	r := make(map[int]int)
	for k, v := range n {
		if _, ok := r[t-v]; ok {
			fmt.Println(r[t-v], k)
			return 0
		} else {
			r[v] = k
		}
	}
	return 0
}
func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	u := strings.Split(input, " ")
	var numbers []int
	for _, uu := range u {
		num, _ := strconv.Atoi(uu)
		numbers = append(numbers, num)
	}
	fmt.Scan(&target)
	sum(numbers, target)
}
