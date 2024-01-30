package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var Assets embed.FS

var MainChar = GetSingleImage("mainchar.png")
var LeftChar = GetSingleImage("left.png")
var RightChar = GetSingleImage("right.png")
var Background = GetSingleImage("Sprite-0001.png")

func GetSingleImage(name string) *ebiten.Image {
	file, err := Assets.Open(name)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
