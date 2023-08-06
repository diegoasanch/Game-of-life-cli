package app

import (
	"fmt"
	"time"

	"github.com/diegoasanch/Game-of-life-cli/engine"
	"github.com/diegoasanch/Game-of-life-cli/renderer"
)

func Run() {
	// fps := 45
	// rate := 1 / float64(fps)

	start := time.Now()
	width, height := renderer.GetConsoleDimensions()
	iterationsCount := 1000

	table := engine.CreateTable(height-4, width/2-1)
	engine.FillTableRandom(&table)
	footer := ""
	times := make([]int64, iterationsCount)
	var iterationStart time.Time
	lastIterationTimeMicro := int64(0)
	lastIterationTimeSeconds := float64(0)
	lastIterationTimeMilli := float64(0)
	// remaining := float64(0)
	// lastIterationTimeNano := int64(0)

	for i := 0; i < iterationsCount; i++ {
		iterationStart = time.Now()

		renderer.PrintTable(&table, footer)
		table = engine.Iterate(&table)
		// table = engine.IterateConcurrent(&table)

		lastIterationTimeMicro = time.Since(iterationStart).Microseconds()
		lastIterationTimeSeconds = time.Since(iterationStart).Seconds()
		lastIterationTimeMilli = float64(lastIterationTimeMicro) / 1000.0
		// lastIterationTimeNano = time.Since(iterationStart).Nanoseconds()

		times[i] = lastIterationTimeMicro
		currentFPS := 1.0 / lastIterationTimeSeconds

		footer = fmt.Sprintf("Iteration %d    FPS %.0f     Last iteration time %.2f ms", i, currentFPS, lastIterationTimeMilli)

		// remaining = rate - lastIterationTimeSeconds

		// sleep for 1/fps of a second
		// if remaining > 0 {
		// 	time.Sleep(time.Duration(time.Duration((remaining)).Seconds()))

		// }
	}

	duration := time.Since(start)
	averageIterationTime := sumSlice(times) / int64(iterationsCount)
	averageFPS := 1.0 / (float64(averageIterationTime) / 1000000.0)

	fmt.Printf("\nRun time: %s\n- avg FPS %.1f \n- avg iteration time: %f ms\n- Dimensions - w: %d, h: %d\n", duration, averageFPS, float64(averageIterationTime)/1000.0, width/2, height)
}

func sumSlice(toSum []int64) int64 {
	var sum int64
	for _, value := range toSum {
		sum += value
	}
	return sum
}
