package app

import (
	"fmt"
	"time"

	"github.com/diegoasanch/Game-of-life-cli/engine"
	"github.com/diegoasanch/Game-of-life-cli/renderer"
)

func Run() {
	start := time.Now()
	width, height := renderer.GetConsoleDimensions()
	iterationsCount := 1000

	table := engine.CreateTable(height-4, width/2 -1)
	engine.FillTableRandom(&table)
	footer := ""
	times := make([]int64, iterationsCount)
	var iterationStart time.Time
	lastIterationTime := int64(0)

	for i := 0; i<iterationsCount; i++{
		iterationStart = time.Now()

		footer = fmt.Sprintf("Iteration %d    Last iteration time %d", i, lastIterationTime)
		renderer.PrintTable(&table, footer)
		table = engine.Iterate(&table)
		// table = engine.IterateConcurrent(&table)

		// time.Sleep(time.Millisecond * 30)
		lastIterationTime = time.Since(iterationStart).Microseconds()
		times[i] = lastIterationTime
	}

	duration := time.Since(start)
	averageIterationTime := sumSlice(times) / int64(iterationsCount)

	fmt.Printf("\nRun time: %s\n- avg iteration time: %f ms\n- Dimensions - w: %d, h: %d\n", duration, float64(averageIterationTime) / 1000.0, width, height)
}

func sumSlice(toSum []int64) int64 {
	var sum int64
	for _, value := range(toSum) {
		sum += value
	}
	return sum
}
