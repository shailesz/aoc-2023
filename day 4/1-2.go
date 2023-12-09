package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cardsWon(match map[int]int, cardNumber int) []int {
	cardswon := make([]int, 0)

	if match[cardNumber] > 0 {
		for j := 0; j < match[cardNumber]; j++ {
			cardswon = append(cardswon, cardNumber+j+1)
		}
	}

	return cardswon
}

func main() {
	input := "input.txt"
	cards := readInput(input)

	xwins := parseInput(cards)
	match := make(map[int]int)
	winnings := make(map[int][]int)

	queue := make([]int, 0)
	processed := make([]int, 0)

	for i, v := range xwins {
		match[i] = len(v)
	}

	for i := 0; i < len(match); i++ {
		cardswon := cardsWon(match, i)

		winnings[i] = cardswon

		queue = append(queue, i)
	}

	for len(queue) > 0 {
		queue = append(queue, winnings[queue[0]]...)

		processed = append(processed, queue[0])

		if len(queue) == 1 {
			queue = []int{}
		} else {
			queue = queue[1:]
		}
	}

	fmt.Println(len(processed))

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

		xwins = append(xwins, xnewgoblin)
	}

	return xwins
}
