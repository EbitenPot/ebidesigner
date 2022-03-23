package form

import (
	"github.com/EbitenPot/ebidesigner/ebiwindow"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainformFields struct {
}

func (f *TMainform) OnFormCreate(sender vcl.IObject) {

}

func (f *TMainform) OnStartButtonClick(sender vcl.IObject) {
	vcl.Application.CreateForm(&EbitenForm)
	EbitenForm.Show()
}

func (f *TMainform) OnXxxxxxChange(sender vcl.IObject) {
	ebiwindow.EbitenGame.Px = float64(f.Xxxxxx.Position())
}

func (f *TMainform) OnYyyyyyyChange(sender vcl.IObject) {
	ebiwindow.EbitenGame.Py = float64(f.Yyyyyyy.Position())
}
