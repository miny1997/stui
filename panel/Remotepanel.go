package panel

import "github.com/jroimartin/gocui"

func RemoteInitialize(g *gocui.Gui) {
	maxX, maxY := g.Size()
	remoteview, err := g.SetView("remote", maxX/2, 1, maxX-1, maxY-3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
		remoteview.Title = " remote server "
		remoteview.Highlight = false
	}
}
