package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

func main() {
	test := NewTest(100)

	defer func() {
		ui.Close()
		fmt.Printf("test.started: %v\n", test.started)
	}()

	test.Init()
	test.Show()
}
