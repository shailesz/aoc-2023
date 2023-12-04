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

		sets := strings.Split(str[1], ";")

		shi := make(map[string]int)

		for _, v := range sets {
			balls := strings.Split(v, ",")

			for _, ball := range balls {
				bc := strings.Split(strings.Trim(ball, " "), " ")

				count := bc[0]
				color := bc[1]

				c, _ := strconv.Atoi(count)

				switch color {
				case Green:
					if shi[Green] < c {
						shi[Green] = c
					}
				case Blue:
					if shi[Blue] < c {
						shi[Blue] = c
					}

				case Red:
					if shi[Red] < c {
						shi[Red] = c
					}

				}

			}

		}

		tmp := 1
		for _, v := range shi {
			tmp *= v
		}

		sum += tmp

	}

	fmt.Println(sum)
}
