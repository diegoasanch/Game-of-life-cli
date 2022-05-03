package renderer

import (
	"fmt"
	"strings"

	"github.com/diegoasanch/Game-of-life-cli/engine"

	tm "github.com/buger/goterm"
)

func init() {
	tm.Clear()
}
func PrintTable(table *engine.Table, footer string) {
	tm.MoveCursor(1, 1)
	printString := ""

	columns := len((*table)[0])
	printString += fmt.Sprintf("\n.%s.\n", strings.Repeat("-", columns * 2))

	for _, row := range(*table) {
		printString += "|"
		for _, col := range(row) {
			if col {
				printString += "\u25a0 "
			} else {
				printString += "  "
			}
		}
		printString += "|\n"
	}
	printString += fmt.Sprintf(".%s.\n%s", strings.Repeat("-", columns * 2), footer)
	tm.Print(printString)
	tm.Flush()
}
