package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/loov/hrtime"
	"github.com/loov/plot"
)

func main() {
	ConsoleHistogram()
	DensityPlot()
	PercentilesPlot()
	TimingPlot()
	StackedPlot()
}

// ConsoleHistogram demonstrates how to measure and print the output to console.
func ConsoleHistogram() {
	fmt.Println("Console Histogram")

	bench := hrtime.NewBenchmark(4 << 10)
	for bench.Next() {
		time.Sleep(5000 * time.Nanosecond)
	}
	fmt.Println(bench.Histogram(10))
}

// TimingPlot demonstrates how to plot timing values based on the order.
func TimingPlot() {
	fmt.Println("Timing Plot (timing.svg)")

	bench := hrtime.NewBenchmark(4 << 10)
	for bench.Next() {
		time.Sleep(5000 * time.Nanosecond)
	}

	seconds := plot.DurationToSeconds(bench.Laps())

	p := plot.New()
	p.Margin = plot.R(5, 0, 0, 5)
	p.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewLine("", plot.Points(nil, seconds)),
		plot.NewTickLabels(),
	)

	svg := plot.NewSVG(800, 300)
	p.Draw(svg)
	ioutil.WriteFile("timing.svg", svg.Bytes(), 0755)
}

// DensityPlot demonstrates how to create a density plot from the values.
func DensityPlot() {
	fmt.Println("Density Plot (density.svg)")

	bench := hrtime.NewBenchmark(4 << 10)
	for bench.Next() {
		time.Sleep(5000 * time.Nanosecond)
	}

	seconds := plot.DurationToSeconds(bench.Laps())

	p := plot.New()
	p.Margin = plot.R(5, 0, 0, 5)
	p.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewDensity("", seconds),
		plot.NewTickLabels(),
	)

	svg := plot.NewSVG(800, 300)
	p.Draw(svg)
	ioutil.WriteFile("density.svg", svg.Bytes(), 0755)
}

// PercentilesPlot demonstrates how to create a percentiles plot from the values.
func PercentilesPlot() {
	fmt.Println("Percentiles Plot (percentiles.svg)")

	bench := hrtime.NewBenchmark(4 << 10)
	for bench.Next() {
		time.Sleep(5000 * time.Nanosecond)
	}

	seconds := plot.DurationToSeconds(bench.Laps())

	p := plot.New()
	p.Margin = plot.R(5, 0, 0, 5)
	p.X = plot.NewPercentilesAxis()
	p.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewPercentiles("", seconds),
		plot.NewTickLabels(),
	)

	svg := plot.NewSVG(800, 300)
	p.Draw(svg)
	ioutil.WriteFile("percentiles.svg", svg.Bytes(), 0755)
}

// StackedPlot demonstrates how to combine plots
func StackedPlot() {
	fmt.Println("Stacked Plot (stacked.svg)")

	bench := hrtime.NewBenchmark(4 << 10)
	for bench.Next() {
		time.Sleep(5000 * time.Nanosecond)
	}

	p := plot.New()
	stack := plot.NewVStack()
	stack.Margin = plot.R(5, 5, 5, 5)
	p.Add(stack)

	seconds := plot.DurationToSeconds(bench.Laps())

	lineplot := plot.NewAxisGroup()
	stack.Add(lineplot)
	lineplot.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewLine("", plot.Points(nil, seconds)),
		plot.NewTickLabels(),
	)

	densityplot := plot.NewAxisGroup()
	stack.Add(densityplot)
	densityplot.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewDensity("", seconds),
		plot.NewTickLabels(),
	)

	percentilesplot := plot.NewAxisGroup()
	percentilesplot.X = plot.NewPercentilesAxis()
	stack.Add(percentilesplot)
	percentilesplot.AddGroup(
		plot.NewGrid(),
		plot.NewGizmo(),
		plot.NewPercentiles("", seconds),
		plot.NewTickLabels(),
	)

	svg := plot.NewSVG(800, 600)
	p.Draw(svg)
	ioutil.WriteFile("stacked.svg", svg.Bytes(), 0755)
}
