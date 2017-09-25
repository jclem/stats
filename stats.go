package main

import (
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strconv"
)

var (
	app        = kingpin.New("stats", "A CLI for common stats functions")
	p75        = app.Command("p75", "Return the p75 of the input")
	perc       = app.Command("perc", "Return the relative standing in the input")
	percentile = perc.Arg("percentile", "The percentile to return the relative standing of").Required().Float64()
)

func main() {
	app.Version("0.1.0")
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case p75.FullCommand():
		commandP75()
	case perc.FullCommand():
		commandPerc(*percentile)
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
