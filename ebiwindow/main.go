// Created by EldersJavas(EldersJavas&gmail.com)

package ebiwindow

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var EbitenGame Game

type Game struct {
	Title        string
	ScreenWidth  int
	ScreenHeight int
	Objects      []*Object
	OK           int
	Px           float64
	Py           float64
}

//func NewGame() *Game {
//	g := new(Game)
//	return
//}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))

	// screen.Fill(color.RGBA{R: 124, G: 5, B: 233,A: 255})

	for _, o := range g.Objects {
		switch o.Otype {
		case EBITEN_IMAGE:
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(1, 1)
			op.GeoM.Translate(g.Px, g.Py)
			screen.DrawImage(o.U.(*ebiten.Image), op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
func init() {
	EbitenGame.Title = "Ebiten"
	EbitenGame.OK = 0
	if EbitenGame.OK == 0 {
		EbitenGame.Objects = append(EbitenGame.Objects, &Object{})
		for i, o := range EbitenGame.Objects {
			o.Ttest()
			fmt.Println(o, "ok", i)
		}
		EbitenGame.OK = 1
	}
}

func InitEbiten() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(EbitenGame.Title)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetMaxTPS(3)
	if err := ebiten.RunGame(&EbitenGame); err != nil {
		log.Println(err)
	}
}
