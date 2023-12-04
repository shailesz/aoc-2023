package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var dm = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var dmv = make([]string, 0)

func parseForward(line string, c chan<- string) {
	fn := ""

	for i := 0; i < len(line); i++ {
		// check if a slice contains a string and print true/false
		ms := string([]rune(line)[i])
		if strings.Contains(strings.Join(dmv, ""), ms) {
			fn = ms
			break
		}

		if i+3 < len(line) {
			three := line[i : i+3]
			v, ok := dm[three]

			if ok {
				fn = v
				break
			}
		}

		if i+4 < len(line) {
			four := line[i : i+4]
			v, ok := dm[four]

			if ok {
				fn = v
				break
			}
		}

		if i+5 < len(line) {
			five := line[i : i+5]
			v, ok := dm[five]

			if ok {
				fn = v
				break
			}
		}

	}

	c <- fn
}

func parseBackward(line string, c chan<- string) {
	ln := ""

	for i := len(line) - 1; i >= 0; i-- {
		ms := string([]rune(line)[i])
		if strings.Contains(strings.Join(dmv, ""), ms) {
			ln = ms
			break
		}

		if i-3 >= 0 {
			three := line[i-2 : i+1]

			v, ok := dm[three]

			if ok {
				ln = v
				break
			}
		}

		if i-4 >= 0 {
			four := line[i-3 : i+1]
			v, ok := dm[four]

			if ok {
				ln = v
				break
			}
		}

		if i-5 >= 0 {
			five := line[i-4 : i+1]
			v, ok := dm[five]

			if ok {
				ln = v
				break
			}
		}
	}

	c <- ln
}

func parseLine(line string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fn := ""
	ln := ""

	fc := make(chan string)
	bc := make(chan string)

	go parseForward(line, fc)
	go parseBackward(line, bc)

	for {
		select {
		case fn = <-fc:
		case ln = <-bc:
		}

		if fn != "" && ln != "" {
			break
		}
	}

	c <- fn + ln
}

func SumFromFile(filePath string) int {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	c := make(chan string)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go parseLine(line, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		i, _ := strconv.Atoi(v)

		sum += i
	}

	return sum
}

func main() {
	filePath := "1.txt"
	for _, v := range dm {
		dmv = append(dmv, v)
	}

	fmt.Println(SumFromFile(filePath))
}
