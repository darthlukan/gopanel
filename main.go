// main
package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"log"
)

type Panel struct {
	bgColor, height, width, display int
	extType                         []string
}

func (p *Panel) SetHeight(h int) {
	p.height = h
	return
}

func (p *Panel) GetHeight() int {
	height := p.height
	return height
}

func (p *Panel) SetWidth(w int) {
	p.width = w
	return
}

func (p *Panel) GetWidth() int {
	width := p.width
	return width
}

func (p *Panel) SetBgColor(c int) {
	p.bgColor = c
	return
}

func (p *Panel) GetBgColor() int {
	color := p.bgColor
	return color
}

func (p *Panel) SetDisplay(d int) {
	p.display = d
	return
}

func (p *Panel) GetDisplay() int {
	display := p.display
	return display
}

func NewPanel(X *xgbutil.XUtil) {
	panel := new(Panel)

	window, err := xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}
	panel.SetBgColor(0x191919)
	panel.SetHeight(26)

	screen := X.Screen()
	width := int(screen.WidthInPixels)

	if width > 1920 { // Must be multi-head, don't span.
		width = 1920
	}
	panel.SetWidth(width)
	panel.SetDisplay(0) // Default to the primary display
	panel.extType = []string{"_NET_WM_WINDOW_TYPE_DOCK"}

	fmt.Sprintf("panel == %v\n", panel)

	window.Create(X.RootWin(), 0, 0, panel.GetWidth(), panel.GetHeight(),
		xproto.CwBackPixel|xproto.CwEventMask, uint32(panel.GetBgColor()),
		xproto.EventMaskButtonRelease)

	// TODO: Val will be either an error or nothing, check for that.
	val := ewmh.WmWindowTypeSet(X, window.Id, panel.extType)
	fmt.Sprintf("val == %v\n", val)

	window.Map()
}

func main() {
	fmt.Println("Hello World!")

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	mousebind.Initialize(X)

	NewPanel(X)

	xevent.Main(X)

}
