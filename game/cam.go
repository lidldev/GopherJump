package game

import "github.com/hajimehoshi/ebiten/v2"

type camera struct {
	x int
	y int

	img *ebiten.Image
}

func (c *camera) setPos(x, y int) {
	c.x = x
	c.y = y
}

func (c *camera) init() {
	c.img = ebiten.NewImage(800, 600)
}

func (c *camera) render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(c.img, op)
}

func (c *camera) draw(img *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(-c.x), float64(-c.y))

	c.img.DrawImage(img, op)
}

func (c *camera) clear() {
	c.img.Clear()
}
