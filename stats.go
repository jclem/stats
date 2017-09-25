package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("stats", "A CLI for common stats functions")

	max = app.Command("max", "Calculate the max of the input")

	mean = app.Command("mean", "Calculate the mean of the input")

	median = app.Command("median", "Calculate the median of the input")

	min = app.Command("min", "Calculate the min of the input")

	mode = app.Command("mode", "Calculate the mode of the input")

	p75 = app.Command("p75", "Return the p75 of the input")

	perc       = app.Command("perc", "Return the relative standing in the input")
	percentile = perc.Arg("percentile",
		"The percentile to return the relative standing of").Required().Float64()
)

func main() {
	app.Version("0.1.0")
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case max.FullCommand():
		commandMax()
	case mean.FullCommand():
		commandMean()
	case median.FullCommand():
		commandMedian()
	case min.FullCommand():
		commandMin()
	case mode.FullCommand():
		commandMode()
	case p75.FullCommand():
		commandP75()
	case perc.FullCommand():
		commandPerc(*percentile)
	}
}

func commandMax() {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	max, err := stats.Max(floats)
	if err != nil {
		panic(err)
	}

	fmt.Println(max)
}

func commandMean() {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	mean, err := stats.Mean(floats)
	if err != nil {
		panic(err)
	}

	fmt.Println(mean)
}

func commandMedian() {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	mean, err := stats.Median(floats)
	if err != nil {
		panic(err)
	}

	fmt.Println(mean)
}

func commandMin() {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	min, err := stats.Min(floats)
	if err != nil {
		panic(err)
	}

	fmt.Println(min)
}

func commandMode() {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	mode, err := stats.Mode(floats)
	if err != nil {
		panic(err)
	}

	for _, int := range mode {
		fmt.Println(int)
	}
}

func commandPerc(percentile float64) {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	standing, err := stats.Percentile(floats, percentile)
	if err != nil {
		panic(err)
	}

	fmt.Println(standing)
}

func commandP75() {
	commandPerc(75)
}

func readNumbers() ([]float64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	values := make([]float64, 0)

	for scanner.Scan() {
		text := scanner.Text()
		float, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return nil, err
		}
		values = append(values, float)
	}

	return values, nil
}
