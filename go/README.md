go-life
=======

A terminal screen Conway's Game of Life implementation that I made when I was learning Go.

Critics and sugestions are welcome!

Have fun!

## Requirements

To build and run this project you will need:

* Go 1.5.1 (older version sure may work)

## Using

To build and run, clone this repository and navigate to go folder. Now run:

    go build -o go-life
    ./go-life

To just run:

    go run *.go

There are some optional arguments for go-life:

* `-width=X` set world width, default: 30
* `-height=X` set world height, default: 30
* `-fps=X` set screen update rate in frames per second, default: 5
* `-limit=X` set limit of generations to run, 0 to no limit, default: 0
