package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lidldev/GopherJump/assets"
)

type Game struct {
	player Player
	c      camera
}

func NewGame() *Game {
	g := &Game{}
	g.c.init()

	return g
}

func (g *Game) Update() error {
	g.player.Update()
	g.c.setPos(g.player.c.x/unit-275, 0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.c.clear()

	s := assets.MainChar

	if g.player.c.vx > 0 {
		s = assets.LeftChar
	} else if g.player.c.vx < 0 {
		s = assets.RightChar
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(g.player.c.x)/unit, float64(g.player.c.y)/unit)

	g.c.draw(s, op)

	g.c.render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
