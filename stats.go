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
		simpleStatCommand(stats.Max)
	case mean.FullCommand():
		simpleStatCommand(stats.Mean)
	case median.FullCommand():
		simpleStatCommand(stats.Median)
	case min.FullCommand():
		simpleStatCommand(stats.Min)
	case mode.FullCommand():
		commandMode()
	case p75.FullCommand():
		commandPerc(75)
	case perc.FullCommand():
		commandPerc(*percentile)
	}
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

func simpleStatCommand(fn func(stats.Float64Data) (float64, error)) {
	floats, err := readNumbers()
	if err != nil {
		panic(err)
	}

	result, err := fn(floats)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
