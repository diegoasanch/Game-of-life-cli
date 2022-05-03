module github.com/diegoasanch/Game-of-life-cli

go 1.18

replace github.com/diegoasanch/Game-of-life-cli/app => ./app

replace github.com/diegoasanch/Game-of-life-cli/engine => ./engine

replace github.com/diegoasanch/Game-of-life-cli/renderer => ./rederer

require github.com/buger/goterm v1.0.4

require golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
