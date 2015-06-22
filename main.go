package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Amida(array []int, amida []int) []int {
	for i, v := range amida {
		if 1 == v {
			t := array[i+1]
			array[i+1] = array[i]
			array[i] = t
		}
	}
	return array
}

func readInput() (n [][]int) {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == nil && len(line) == 0 {
			err = io.EOF
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		a := strings.Split(line, ",")
		an := make([]int, len(a))
		for i, v := range a {
			an[i], err = strconv.Atoi(v)
		}
		n = append(n, an)
	}
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}

	return n
}

func main() {
	var actual []int
	input := readInput()
	array := make([]int, len(input[0])+1)

	for i, _ := range array {
		array[i] = i + 1
	}
	for _, v := range input {
		actual = Amida(array, v)
	}
	for i, s := range actual {
		fmt.Print(s)
		if i < len(actual)-1 {
			fmt.Print(",")
		} else {
			fmt.Println()
		}

	}
}
