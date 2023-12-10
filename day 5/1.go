package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {
	seeds, mm := mapInput("input.txt")
	locations := make([]Location, 0)

	for _, seed := range seeds {

		var curr = seed
		for _, m := range mm {
			curr = findElement(m, curr)
		}

		locations = append(locations, curr)
	}

	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Value < locations[j].Value
	})

	fmt.Println(len(seeds))
}

func mapInput(input string) ([]Seed, []map[Element]Element) {
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

	return seeds, myMap
}

func findElement(seedToSoil map[Element]Element, elem Element) Element {

	for k, v := range seedToSoil {
		start := k.Value
		end := k.Value + k.Length - 1

		if elem.Value >= start && elem.Value <= end {
			diff := elem.Value - start
			e := v.Value + diff

			return Element{Value: e}
		}
	}

	return Element{Value: elem.Value}
}
