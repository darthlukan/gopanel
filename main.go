// main
package main

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	//"time"
)

func main() {
	fmt.Println("Hello World!")
	gtk.Init(nil)

	builder, err := gtk.BuilderNew()

	if err != nil {
		panic(err)
	}

	builder.AddFromFile("panel.glade")
	panel, err := builder.GetObject("Panel")

	if err != nil {
		panic(err)
	}

	if w, ok := panel.(*gtk.Window); ok {
		w.Connect("destroy", func() { gtk.MainQuit() })
		button, err := builder.GetObject("SysActions")

		if err != nil {
			panic(err)
		}

		if b, good := button.(*gtk.Button); good {
			b.Connect("clicked", func() { gtk.MainQuit() })
		}

		w.ShowAll()
	} else {
		panic(ok)
	}
	gtk.Main()
}
