package main

import (
    "fmt"
    "flag"
    "time"
    "os"
    "os/exec"
)

// The Main
func main() {
    // command-line flags parsing
    var width, height, fps, limit int
    flag.IntVar(&width, "width", 30, "grid width") // grid width
    flag.IntVar(&height, "height", 30, "grid height") // grid height
    flag.IntVar(&fps, "fps", 5, "frames per second") // frames displayed per second
    flag.IntVar(&limit, "limit", 0, "limit of generations") // limit of generations to run
    flag.Parse()

    // create a new life habitat
    life := NewLife(width, height, limit)
    // generate random initial life
    life.Generate()

    // loop over generations
    for cont := true; cont; cont = life.NextGeneration() {
        // clear screen (won't work in Windows)
        clear := exec.Command("clear")
        clear.Stdout = os.Stdout
        clear.Run()

        // print current generation
        fmt.Println(life.String())
        // sleeps according to desired fps
        time.Sleep(time.Second / time.Duration(fps))
    }
}
