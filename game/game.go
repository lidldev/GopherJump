package game

import (
	"image/color"

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
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.c.clear()

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(1, 0.4)
	op2.GeoM.Translate(-3000, 0)
	g.c.draw(assets.Background, op2)

	s := assets.MainChar

	if g.player.c.vx > 0 {
		s = assets.RightChar
	} else if g.player.c.vx < 0 {
		s = assets.LeftChar
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(g.player.c.x)/unit, float64(g.player.c.y)/unit-25)

	g.c.draw(s, op)

	g.c.render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
