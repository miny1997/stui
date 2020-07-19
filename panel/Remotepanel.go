package panel

import "github.com/jroimartin/gocui"

func RemoteInitialize(g *gocui.Gui, remoteview *gocui.View) {
	remoteview.Title = " remote server "
	remoteview.Highlight = false
}
