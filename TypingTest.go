package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

var version = "0.0.1"

// Test represents a typing test
type Test struct {
	started bool

	totalWords int
	rightWords int
	wrongWords int

	totalKeys int
	rightKeys int
	wrongKeys int
}

// NewTest makes and returns a pointer to a new Test instance
func NewTest(wordCount int) *Test {
	return &Test{totalWords: wordCount}
}

// Init initializes a test
func (t *Test) Init() *Test {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	t.CreateUIElements()
	t.SetUIHandlers()

	return t
}

var timeGauge ui.Gauge

// CreateUIElements sets up the termui layout
func (t *Test) CreateUIElements() {
	p := ui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 80
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = fmt.Sprintf("GoType v%s", version)
	p.BorderFg = ui.ColorCyan

	timeGauge := ui.NewGauge()
	timeGauge.Percent = 50
	timeGauge.Height = 3
	timeGauge.BorderLabel = "Time Remaining"
	timeGauge.Label = "{{percent}}"
	timeGauge.LabelAlign = ui.AlignCenter

	timeGauge.BarColor = ui.ColorBlue
	timeGauge.BorderFg = ui.ColorWhite
	timeGauge.BorderLabelFg = ui.ColorCyan

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(8, 2, p)),
		ui.NewRow(ui.NewCol(8, 2, timeGauge)))

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)
}

// SetUIHandlers sets the termui event handlers needed
func (t *Test) SetUIHandlers() {
	ui.Handle("/sys/kbd/q", func(e ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		k := e.Data.(ui.EvtKbd)

		t.totalKeys++
		if k.KeyStr == "<space>" {
			fmt.Println("-> S P A C E")
		} else {
			fmt.Printf("-> %v (%v)\n", e.Path, k.KeyStr)
		}
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		timeGauge.Label = fmt.Sprintf("Time Remaining (%ds)", 60-t.Count)
		timeGauge.Percent = int(t.Count / 60)
		ui.Render(&timeGauge)

		if t.Count == 10 {
			delete(ui.DefaultEvtStream.Handlers, "/timer/1s")
		}
	})
}

func (t *Test) Show() {
	ui.Loop()
}
