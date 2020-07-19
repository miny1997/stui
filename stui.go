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
	if err := g.SetKeybinding("", gocui.KeyEnd, gocui.ModNone, quit); err != nil {
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
		cmds := []string{"\u001b[47m\u001b[30mend\u001b[0m\u001b[0m quit ",
			"\u001b[47m\u001b[30mfn2\u001b[0m\u001b[0m file info ",
			"\u001b[47m\u001b[30mfn3\u001b[0m\u001b[0m reload dir\n",
			"\u001b[47m\u001b[30marrow keys\u001b[0m\u001b[0m move ",
			"\u001b[47m\u001b[30menter\u001b[0m\u001b[0m options ",
			"\u001b[47m\u001b[30mshift\u001b[0m\u001b[0m change side "}
		for _, cmd := range cmds {
			fmt.Fprint(v, cmd)
		}
		//fmt.Fprintln(v, "end - quit  f2 - file info  f3 - reload directory  "+
		//	"arrow keys - move  enter - open options  shift - change side")
	}
	if localview, err := g.SetView("local", 0, 1, maxX/2-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		panel.LocalInitialize(g, localview)
	}
	if remoteview, err := g.SetView("remote", maxX/2, 1, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		panel.RemoteInitialize(g, remoteview)
	}
	g.SetCurrentView("local")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	os.Exit(0)
	return gocui.ErrQuit
}
