// Created by EldersJavas(EldersJavas&gmail.com)

package ebiwindow

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 124, G: 5, B: 233,A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func InitEbiten() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebiten")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowFloating(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Println(err)
	}
}
