package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ACCELERATION = 1

func main() {
	xtime, xdistance := readInput("input.txt")

	partOne(xtime, xdistance)
	partTwo(xtime, xdistance)
}

func partTwo(xtime, xdistance []string) {
	t := strings.Join(xtime, "")
	d := strings.Join(xdistance, "")

	time, _ := strconv.Atoi(t)
	highScore, _ := strconv.Atoi(d)
	waysToBeat := 0

	for hold := 0; hold < time; hold++ {
		timeToMove := time - hold
		velocity := hold * ACCELERATION

		distanceMoved := timeToMove * velocity

		if distanceMoved > highScore {
			waysToBeat += 1
		}
	}

	fmt.Println(waysToBeat)

}

func partOne(xtime, xdistance []string) {
	answer := 1
	for i, time := range xtime {
		waysToBeat := 0
		highScore, _ := strconv.Atoi(xdistance[i])

		t, _ := strconv.Atoi(time)
		for hold := 0; hold < t; hold++ {
			timeToMove := t - hold

			velocity := hold * ACCELERATION

			distanceMoved := timeToMove * velocity

			if distanceMoved > highScore {
				waysToBeat += 1
			}

		}

		answer *= waysToBeat
	}
}

func readInput(input string) ([]string, []string) {
	f, _ := os.Open(input)

	scanner := bufio.NewScanner(f)
	xtime := make([]string, 0)
	xdistance := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, "Time:") {
			text = strings.Trim(strings.Replace(text, "Time:", "", 1), " ")
			xsplit := strings.Split(text, " ")
			xnum := make([]string, 0)

			for _, v := range xsplit {
				if v != "" && v != " " {
					xnum = append(xnum, v)
				}
			}

			xtime = xnum
		}

		if strings.Contains(text, "Distance:") {
			text = strings.Trim(strings.Replace(text, "Distance:", "", 1), " ")
			xsplit := strings.Split(text, " ")
			xnum := make([]string, 0)

			for _, v := range xsplit {
				if v != "" && v != " " {
					xnum = append(xnum, v)
				}
			}

			xdistance = xnum
		}
	}

	return xtime, xdistance
}
