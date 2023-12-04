package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Green = "green"
	Blue  = "blue"
	Red   = "red"
)

func main() {
	sum()
}

func sum() {
	file := "input.txt"

	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		str := strings.Split(line, ":")
		sn := strings.ReplaceAll(str[0], "Game ", "")
		num, _ := strconv.Atoi(sn)

		sets := strings.Split(str[1], ";")

		valid := true
	loop:
		for _, v := range sets {
			balls := strings.Split(v, ",")

			for _, ball := range balls {
				bc := strings.Split(strings.Trim(ball, " "), " ")

				count := bc[0]
				color := bc[1]

				c, _ := strconv.Atoi(count)

				switch color {
				case Green:

					if c > 13 {
						valid = false
						break loop
					}
				case Blue:

					if c > 14 {
						valid = false
						break loop
					}
				case Red:

					if c > 12 {
						valid = false
						break loop
					}
				}

			}

		}

		if valid {
			sum += num
		}

	}

	fmt.Println(sum)
}
