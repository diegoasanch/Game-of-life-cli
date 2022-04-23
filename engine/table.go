package engine

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type Table = [][]bool

func CreateTable(rows, columns int) Table {
	table := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		table[i] = make([]bool, columns)
	}
	return table
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FillTableRandom(table *Table) {
	for _, row := range(*table) {
		for j := range(row) {
			row[j] = math.Round(rand.Float64()) == 1
		}
	}
}

func Iterate(table *Table) Table {
	newTable := CreateTable(len(*table), len((*table)[0]))

	for row := range(*table) {
		for column := range((*table)[row]) {
			newTable[row][column] = determineCell(row, column, table)
		}
	}

	return newTable
}

func IterateConcurrent(table *Table) Table {
	rows := len(*table)
	columns := len((*table)[0])
	newTable := CreateTable(rows, columns)

	wg := new(sync.WaitGroup)
	wg.Add(4)

	go iteratePortion(table, &newTable, 0, rows/2, 0, columns/2, wg)
	go iteratePortion(table, &newTable, rows/2, rows, 0, columns/2, wg)
	go iteratePortion(table, &newTable, 0, rows/2, columns/2, columns, wg)
	go iteratePortion(table, &newTable, rows/2, rows, columns/2, columns, wg)

	wg.Wait()

	return newTable
}

func iteratePortion(ogTable, newTable *Table, fromRow, toRow, fromCol, toCol int, wg *sync.WaitGroup) {
	defer wg.Done()
	for row := fromRow; row < toRow; row++ {
		for column := fromCol; column < toCol; column++ {
			(*newTable)[row][column] = determineCell(row, column, ogTable)
		}
	}
}

func determineCell(row, column int, table *Table) bool {
	isAlive := isCellAlive(row, column, table)
	willLive := false
	liveNeighbours := getLiveNeighboursCount(row, column, table)

	if isAlive {
		switch {
			case liveNeighbours < 2: // underpopulation
				willLive = false
			case liveNeighbours > 3: // overpopulation
				willLive = false
			default:
				willLive = true
		}
	} else {
		willLive = liveNeighbours == 3
	}
	return willLive
}

func isCellAlive(row, column int, table *Table) bool {
	return (*table)[row][column]
}

func getLiveNeighboursCount(row, column int, table *Table) int {
	rows, columns := len(*table), len((*table)[0])
	liveNeighbours := 0

	isTop := row == 0
	isBottom := row == (rows - 1)
	isBorderLeft := column == 0
	isBorderRight := column == (columns - 1)

	if !isTop {
		// Top left
		if !isBorderLeft && isCellAlive(row-1, column-1, table) {
			liveNeighbours++
		}
		// Top
		if isCellAlive(row-1, column, table) {
			liveNeighbours++
		}
		// Top right
		if !isBorderRight && isCellAlive(row-1, column+1, table) {
			liveNeighbours++
		}
	}
	// Center left
	if !isBorderLeft && isCellAlive(row, column-1, table) {
		liveNeighbours++
	}
	// Center right
	if !isBorderRight && isCellAlive(row, column+1, table) {
		liveNeighbours++
	}
	if !isBottom {
		// Bottom left
		if !isBorderLeft && isCellAlive(row+1, column-1, table) {
			liveNeighbours++
		}
		// Bottom
		if isCellAlive(row+1, column, table) {
			liveNeighbours++
		}
		// Bottom right
		if !isBorderRight && isCellAlive(row+1, column+1, table) {
			liveNeighbours++
		}
	}
	return liveNeighbours
}
