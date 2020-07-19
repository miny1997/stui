package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"os"
	"stui/panel"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)
	g.Cursor = false
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("title", -1, -1, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "stui - a sftp client with a terminal user interface")
		v.Frame = false
	}
	if v, err := g.SetView("cmd", -1, maxY-3, maxX-1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Frame = false
		fmt.Fprintln(v, "esc - quit  f2 - connect  f3 - disconnect  f4 - file info  f5 - reload directory  "+
			"arrow keys - move  enter - open options  shift - change side")
	}
	panel.LocalInitialize(g)
	panel.RemoteInitialize(g)
	g.SetCurrentView("local")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	os.Exit(0)
	return gocui.ErrQuit
}
