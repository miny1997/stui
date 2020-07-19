package panel

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"io/ioutil"
	"log"
	"os"
)

var path string

func LocalInitialize(g *gocui.Gui) {
	p, e := os.Getwd()
	if e != nil {
		return
	}
	path = p
	maxX, maxY := g.Size()
	localview, err := g.SetView("local", 0, 1, maxX/2-1, maxY-3)
	if err != nil {
		return
	}
	localview.Clear()
	localview.Title = " local machine " + path + " "
	localview.Highlight = true
	localview.SelBgColor = gocui.ColorGreen
	localview.SelFgColor = gocui.ColorBlack
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(localview, "..")
	for _, f := range files {
		fmt.Fprintln(localview, f.Name())
	}
}
