package form

import (
	"github.com/EbitenPot/ebidesigner/ebiwindow"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"os"
)

//::private::
type TEbitenFormFields struct {
}

func (f *TEbitenForm) OnFormClose(sender vcl.IObject, closeAction *types.TCloseAction) {
	os.Exit(0)
}

func (f *TEbitenForm) OnFormCreate(sender vcl.IObject) {
	ebiwindow.InitEbiten()
}
