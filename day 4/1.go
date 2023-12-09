package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input := "test.txt"
	cards := readInput(input)

	xwins := parseInput(cards)

	sum := 0
	for _, v := range xwins {
		l := len(v) - 1

		// fmt.Printf("%s, %d :: ", v, l)

		if l != -1 {
			sum += int(math.Pow(2, float64(l)))
			// fmt.Printf("%d, %d", sum, int(math.Pow(2, float64(l))))
		}

		// fmt.Println()
	}

	fmt.Println(sum)

}

func readInput(input string) []string {

	f, _ := os.Open(input)

	scanner := bufio.NewScanner(f)

	xinp := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		xinp = append(xinp, line)
	}

	return xinp
}

func parseInput(input []string) [][]string {
	xwins := make([][]string, 0)
	for _, v := range input {
		xspl := strings.Split(v, ":")
		xspl[1] = strings.TrimSpace(xspl[1])

		xcombo := strings.Split(xspl[1], "|")

		xcombo[0] = strings.TrimSpace(xcombo[0])
		xcombo[1] = strings.TrimSpace(xcombo[1])

		xwin := strings.Split(xcombo[0], " ")
		xgoblin := strings.Split(xcombo[1], " ")

		xnewgoblin := make([]string, 0)
		for _, gn := range xgoblin {
			for _, wn := range xwin {
				if gn == wn && gn != " " && gn != "" {
					xnewgoblin = append(xnewgoblin, gn)
				}
			}
		}

		fmt.Println(xnewgoblin, len(xnewgoblin))

		xwins = append(xwins, xnewgoblin)
	}

	return xwins
}
