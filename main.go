package main

import (
	"github.com/ying32/govcl/vcl"
	"govcl-template/form"
)

func main() {
    vcl.Application.SetTitle("project")
    vcl.Application.SetScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&form.Mainform)
	vcl.Application.Run()
}
