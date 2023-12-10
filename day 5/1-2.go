package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Element struct {
	Value  int
	Length int
}

type Seed = Element

type Soil = Element

type Fertilizer = Element

type Water = Element

type Light = Element

type Temperature = Element

type Humidity = Element

type Location = Element

// MakeNumberPositive returns the absolute value of a number using a bitwise trick.
func MakeNumberPositive(num int) int {
	mask := num >> 31          // Create a mask of the sign bit
	return (num ^ mask) - mask // Flip the number if it is negative, unchanged if positive
}

var mapp = make([]map[Element]Element, 0)

func main() {
	seeds, mm, _ := mapInput("input.txt")
	locations := make([]*Location, 0)

	mapp = mm

	xs := make([]*Element, 0)

	for i := 0; i < len(seeds); i += 2 {
		xs = append(xs, &Seed{Value: seeds[i].Value, Length: seeds[i+1].Value})
	}

	c := make(chan *Element, 0)
	var wg sync.WaitGroup

	for _, seed := range xs {
		lowest := FindLowestLocationsFromSeed(seed)
		if lowest != nil {
			locations = append(locations, lowest)
			continue
		}

		wg.Add(1)
		fmt.Println("started a new thread")
		go bruteforce(seed, c)
	}

	go func(c chan *Element) {
		wg.Wait()
		close(c)
	}(c)

	for loc := range c {
		locations = append(locations, loc)
		wg.Done()
	}

	for _, loc := range locations {
		fmt.Println(loc.Value, loc.Length)
	}

}

func bruteforce(element *Element, c chan *Element) {

	length := element.Length
	elem := &Element{Value: -1, Length: length}

	for i := 0; i < length; i++ {

		newelem := getLocationForElement(&Element{Value: element.Value + i, Length: length})

		if elem.Value == -1 {
			elem = newelem
			continue
		}

		if newelem.Value < elem.Value {
			elem = newelem
		}

	}

	fmt.Println("Found:: ", elem.Value)
	c <- elem

}

func FindLowestLocationsFromSeed(element *Element) *Element {
	left := getLocationForElement(element)

	right := getLocationForElement(&Element{Value: element.Value + element.Length - 1, Length: element.Length})

	if MakeNumberPositive(right.Value-left.Value)+1 == element.Length {
		fmt.Println("element", element, "already found lowest")
		return left
	}

	return nil
}

func mapInput(input string) ([]Seed, []map[Element]Element, map[Element]Element) {
	seeds := make([]Seed, 0)
	seedToSoil := make(map[Seed]Soil)
	soilToFertilizer := make(map[Soil]Fertilizer)
	fertilizerToWater := make(map[Fertilizer]Water)
	waterToLight := make(map[Water]Light)
	lightToTemperature := make(map[Light]Temperature)
	temperatureToHumidity := make(map[Temperature]Humidity)
	humidityToLocation := make(map[Humidity]Location)

	var myMap = []map[Element]Element{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation}

	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			split := strings.Split(line, ":")
			xsv := strings.Split(strings.TrimSpace(split[1]), " ")

			for _, v := range xsv {
				s, _ := strconv.Atoi(v)
				seeds = append(seeds, Seed{Value: s})
			}
		}

		if strings.Contains(line, "seed-to-soil map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}

				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}

				seedToSoil[Seed{Value: start, Length: length}] = Soil{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "soil-to-fertilizer map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}

				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}

				soilToFertilizer[Soil{Value: start, Length: length}] = Fertilizer{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "fertilizer-to-water map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}

				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}

				fertilizerToWater[Fertilizer{Value: start, Length: length}] = Water{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "water-to-light map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}
				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}
				waterToLight[Water{Value: start, Length: length}] = Light{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "light-to-temperature map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}
				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}
				lightToTemperature[Light{Value: start, Length: length}] = Temperature{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "temperature-to-humidity map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}

				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}
				temperatureToHumidity[Temperature{Value: start, Length: length}] = Humidity{Value: end, Length: length}

			}
		}

		if strings.Contains(line, "humidity-to-location map:") {
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}

				line := scanner.Text()
				split := strings.Split(line, " ")

				start, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println(err)
				}

				end, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println(err)
				}
				length, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("err", err)
				}
				humidityToLocation[Humidity{Value: start, Length: length}] = Location{Value: end, Length: length}

			}
		}

	}

	return seeds, myMap, humidityToLocation
}

func findNextElement(seedToSoil map[Element]Element, elem *Element) *Element {

	for k, v := range seedToSoil {
		start := k.Value
		end := k.Value + k.Length - 1

		if elem.Value >= start && elem.Value <= end {
			diff := elem.Value - start
			e := v.Value + diff

			return &Element{Value: e, Length: elem.Length}
		}
	}

	return &Element{Value: elem.Value, Length: elem.Length}
}

func getLocationForElement(e *Element) *Element {
	elem := e
	for _, m := range mapp {
		elem = findNextElement(m, elem)
	}

	return elem
}
