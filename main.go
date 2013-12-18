// main
package main

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	gtk.Init(nil)
	panel, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(err)
	}

	panel.Connect("destroy", func() {
		gtk.MainQuit()
	})

	panel.SetDecorated(false)
	panel.SetKeepAbove(true)
	panel.SetDefaultGeometry(1920, 24)
	panel.SetHasResizeGrip(false)
	panel.Move(0, 1080)
	panel.Add(giveClock())

	panel.ShowAll()
	gtk.Main()
}

func giveClock() *gtk.Widget {
	// TODO: This should really update every second...
	now := time.Now().String()
	clock, err := gtk.LabelNew(now)
	if err != nil {
		panic(err)
	}
	return &clock.Widget
}
