package main

import (
	"github.com/rivo/tview"

	client "github.com/briancain/waypoint-station/waypoint-client"
)

func main() {
	_, err := client.New(false)
	if err != nil {
		panic(err)
	}

	box := tview.NewBox().SetBorder(true).SetTitle("Hello, Waypoint!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
