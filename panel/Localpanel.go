package panel

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"io/ioutil"
	"log"
	"os"
)

var path string
var files []string
var pos = 0

func LocalInitialize(g *gocui.Gui, localview *gocui.View) {
	p, e := os.Getwd()
	if e != nil {
		return
	}
	path = p
	localview.Title = " local machine " + path + " "
	localview.Highlight = true
	localview.SelBgColor = gocui.ColorGreen
	localview.SelFgColor = gocui.ColorBlack
	fileiter, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	files = append(files, "../")
	fmt.Fprintln(localview)
	for _, f := range fileiter {
		file := f.Name()
		if f.IsDir() {
			files = append(files, file+"/")
		}

	}
	for _, f := range fileiter {
		file := f.Name()
		if !f.IsDir() {
			files = append(files, file)
		}
	}
	for _, f := range files {
		fmt.Fprintln(localview, f)
	}
	if err := g.SetKeybinding("local", gocui.KeyArrowDown, gocui.ModNone, Local_arrow_down); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("local", gocui.KeyArrowUp, gocui.ModNone, Local_arrow_up); err != nil {
		log.Panicln(err)
	}
}

func Local_arrow_down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if len(files)-1 == cy {
			if err := v.SetCursor(cx, 0); err != nil {
				ox, _ := v.Origin()
				if err := v.SetOrigin(ox, 0); err != nil {
					return err
				}
				pos = 0
			}
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			pos = pos + 1
			fmt.Println(pos)
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}
func Local_arrow_up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if cy == 0 {
			if err := v.SetCursor(cx, len(files)-1); err != nil {
				if err := v.SetOrigin(ox, len(files)-1); err != nil {
					return err
				}
			}
			return nil
		}
		if err := v.SetCursor(cx, cy-1); err != nil {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
