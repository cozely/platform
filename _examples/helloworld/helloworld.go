package main

import (
	"fmt"
	"time"

	"github.com/cozely/platform/window"
)

func main() {
	var err error

	w, err := window.New(
		window.Title("Hello, World"),
		window.Size(800, 600),
		window.VSync(true),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(window.InfoString())
	w.Present()

	fmt.Print("Window opened...")
	time.Sleep(4 * time.Second)

	w.Close()
	fmt.Println(" and closed.")
}
