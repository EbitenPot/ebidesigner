package main

import (
	"github.com/EbitenPot/ebidesigner/form"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.SetTitle("ebidesigner")
	vcl.Application.SetScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form.Mainform)
	//go ebiwindow.InitEbiten()
	vcl.Application.Run()

}
