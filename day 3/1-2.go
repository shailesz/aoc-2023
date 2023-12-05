package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Positions []Position

type Position struct {
	i int
	j int
}

func (p Position) containsNumber(matrix [][]string) bool {
	return isNumber(matrix[p.i][p.j])
}

func (p Position) findNumberNeighbors(matrix [][]string) []Position {
	// Relative positions of the 8 neighbors
	rowOffsets := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	colOffsets := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])

	// Function to check if a given position is valid (inside the matrix)
	isValid := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	np := make([]Position, 0)

	for i := 0; i < 8; i++ {
		newRow := p.i + rowOffsets[i]
		newCol := p.j + colOffsets[i]
		if isValid(newRow, newCol) {

			n := matrix[newRow][newCol]

			if isNumber(n) {
				p := Position{newRow, newCol}

				np = append(np, p)
			}

		}
	}

	return np
}

func (p Position) findPartNumberHead(matrix [][]string) Position {
	if p.containsNumber(matrix) {
		for j := p.j - 1; j >= 0; j-- {
			if j == 0 {
				pp := Position{p.i, j}
				if !pp.containsNumber(matrix) {
					return Position{p.i, j + 1}
				}

				return Position{p.i, j}
			}

			if !isNumber(matrix[p.i][j]) {
				return Position{p.i, j + 1}
			}
		}
	}

	return Position{-1, -1}
}

func (xp Positions) removeDuplicates() Positions {
	seen := make(map[Position]bool)
	np := make(Positions, 0)

	for _, p := range xp {
		if _, exists := seen[p]; !exists {
			seen[p] = true
			np = append(np, p)
		}
	}

	return np
}

func (xp Positions) findAdjacentProduct(matrix [][]string) int {
	a, _ := strconv.Atoi(matrix[xp[0].i][xp[0].j])
	b, _ := strconv.Atoi(matrix[xp[1].i][xp[1].j])

	return a * b
}

func (p Position) findNumber(matrix [][]string) int {
	if p.containsNumber(matrix) {
		number := matrix[p.i][p.j]

		for j := p.j + 1; j < len(matrix[p.i]); j++ {
			if isNumber(matrix[p.i][j]) {
				number += matrix[p.i][j]
			} else {
				break
			}
		}

		i, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		return i
	}

	return -1
}

func (p Position) findPartNumberTail(matrix [][]string) Position {
	if p.containsNumber(matrix) {
		for j := p.j + 1; j < len(matrix[p.i]); j++ {
			if !isNumber(matrix[p.i][j]) {
				return Position{p.i, j - 1}
			}
		}
	}

	return Position{-1, -1}
}

func (p Position) isGear(matrix [][]string) bool {
	if matrix[p.i][p.j] == "*" {
		return true
	}

	return false
}

func main() {

	file := "input.txt"

	rep := readInput(file)

	op := make([]Position, 0)
	for i, r := range rep {
		for j, v := range r {
			if !isNumber(v) && !isPeriod(v) {
				op = append(op, Position{i, j})
			}
		}
	}

	gp := make([]Position, 0)
	for _, v := range op {
		if v.isGear(rep) {
			gp = append(gp, v)
		}
	}

	xnh := make([]Positions, 0)
	for _, v := range gp {
		nn := v.findNumberNeighbors(rep)
		nh := make(Positions, 0)

		for _, pos := range nn {
			nh = append(nh, pos.findPartNumberHead(rep))
		}
		nh = nh.removeDuplicates()

		if len(nh) == 2 {
			xnh = append(xnh, nh)
		}
	}

	xn := make([][]int, 0)
	for _, xp := range xnh {
		xxn := make([]int, 0)

		for _, nh := range xp {
			xxn = append(xxn, nh.findNumber(rep))
		}

		xn = append(xn, xxn)
	}

	xap := make([]int, 0)
	for _, xxp := range xn {
		ap := 1

		for _, v := range xxp {
			ap *= v
		}

		xap = append(xap, ap)
	}

	sum := 0
	for _, v := range xap {
		sum += v
	}

	fmt.Println(sum)
}

func findSumForOperators(op []Position, matrix [][]string) int {
	np := make([]Position, 0)
	for _, v := range op {
		np = append(np, v.findNumberNeighbors(matrix)...)
	}

	nh := make(Positions, 0)
	for _, v := range np {
		nh = append(nh, v.findPartNumberHead(matrix))
	}
	nh = nh.removeDuplicates()

	xpn := make([]int, 0)
	for _, v := range nh {
		xpn = append(xpn, v.findNumber(matrix))
	}

	sum := 0
	for _, pn := range xpn {
		sum += pn
	}

	return sum
}

func readInput(input string) [][]string {
	rep := make([][]string, 0)

	f, _ := os.Open(input)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		r := make([]string, 0)

		for _, v := range line {
			r = append(r, string(v))
		}

		rep = append(rep, r)
	}

	return rep
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func isPeriod(s string) bool {
	return s == "."
}
