package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	client "github.com/briancain/waypoint-station/waypoint-client"
)

const (
	// TODO(briancain): Abstract this out to its own func so we can build
	// release versions off of git SHAs rather than it being manual
	Version = "0.1.0"
)

func main() {
	_, err := client.New(false)
	if err != nil {
		panic(err)
	}

	// Run application

	// Default color
	tview.Styles.PrimitiveBackgroundColor = tcell.NewRGBColor(40, 44, 48)

	box := tview.NewBox().SetBorder(true).SetTitle("Hello, Waypoint!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
