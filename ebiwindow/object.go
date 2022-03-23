// Created by EldersJavas(EldersJavas&gmail.com)

package ebiwindow

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	EBITEN_IMAGE = iota
)

type Object struct {
	Otype     int
	Oresource []byte
	U         interface{}
}

func (o *Object) Ttest() {
	o.U = ebiten.NewImage(100, 100)
	fmt.Println("ebiten.NewImage(100, 100) OK ", o.U)
	o.Otype = EBITEN_IMAGE
	o.U.(*ebiten.Image).Fill(color.RGBA{R: 124, G: 5, B: 233, A: 255})
}
