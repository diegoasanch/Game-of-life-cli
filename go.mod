module GameOfLife

go 1.18

replace GameOfLife/app => ./app

replace GameOfLife/engine => ./engine

replace GameOfLife/renderer => ./rederer

require github.com/buger/goterm v1.0.4

require golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
