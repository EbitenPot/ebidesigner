package form

import (
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

func (f *TMainform) OnPageControl1Change(sender vcl.IObject) {

}

