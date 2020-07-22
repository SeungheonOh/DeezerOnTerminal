package layout

import (
	"github.com/GoDeezer/lib/deezer"
	ui "github.com/gizak/termui/v3"
)

type ModuleShare struct {
	DeezerClient  *deezer.Client
	MusicQueue    []string // use []song instead of []string
	MusicCurrent  string
	MusicProgress float64 // 0 to 1
}

type Module interface {
	Render()
	Resize(int, int)
}

// interface for module with event handle
type Layout interface {
	Module
	HandleEvent(ui.Event)
}

type LayoutList struct {
	CurrentLayout int
	Layout        []Layout
}

func (self *LayoutList) Next() {
	self.CurrentLayout++
	if self.CurrentLayout >= len(self.Layout) {
		self.CurrentLayout = 0
	}
}

func (self *LayoutList) Render() {
	self.Layout[self.CurrentLayout].Render()
}