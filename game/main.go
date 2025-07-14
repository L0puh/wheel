package main

import (
	"image/color"
	_ "image/png"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var sprite_car = load_image("./assets/car.png")
var sprite_bg = load_image("./assets/bg.png")

type Game struct {
	car_pos  Vector
	init_pos Vector
	speed    float64
}
type Vector struct {
	x   float64
	y   float64
	rot float64
}

func load_image(name string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		panic(err)
	}

	return img
}

func (g *Game) Update() error {
	g.speed = float64(400 / ebiten.TPS())
	//TODO: change it to values from arduino
	g.car_pos.rot = float64(time.Now().Second())

	var delta Vector
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.y = g.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.y = -g.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.x = g.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.x = -g.speed
	}

	if delta.x != 0 && delta.y != 0 {
		factor := g.speed / math.Sqrt(delta.x*delta.x+delta.y*delta.y)
		delta.x *= factor
		delta.y *= factor
	}

	g.car_pos.x += delta.x
	g.car_pos.y += delta.y

	return nil
}

func Draw_background(screen *ebiten.Image, g *Game) {

	/* prallax bg window to make in infinte */

	if sprite_bg == nil {
		screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
		return
	}

	bg_width := float64(sprite_bg.Bounds().Dx())
	bg_height := float64(sprite_bg.Bounds().Dy())
	screen_width, screen_height := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	parallax_factor := g.speed

	scroll_x := -g.car_pos.x * parallax_factor
	scroll_y := -g.car_pos.y * parallax_factor

	offset_x := math.Mod(scroll_x, screen_width)
	if offset_x < 0 {
		offset_x += screen_width
	}

	offset_y := math.Mod(scroll_y, screen_height)
	if offset_y < 0 {
		offset_y += screen_height
	}

	scale_x := screen_width / bg_width
	scale_y := screen_height / bg_height

	for i := -1; i <= 2; i++ {
		for j := -1; j <= 2; j++ {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Scale(scale_x, scale_y)

			x := offset_x + float64(i)*screen_width
			y := offset_y + float64(j)*screen_height

			op.GeoM.Translate(x, y)

			screen.DrawImage(sprite_bg, op)
		}
	}
}
func (g *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.init_pos.x, g.init_pos.y)
	opts.GeoM.Scale(0.2, 0.2)

	Draw_background(screen, g)
	screen.DrawImage(sprite_car, opts)

}

func (g *Game) Layout(width, height int) (x, y int) {
	return width, height
}

func main() {
	g := &Game{
		car_pos:  Vector{x: 1800, y: 4000, rot: 0},
		init_pos: Vector{x: 1800, y: 3000, rot: 0},
		speed:    10.0,
	}

	ebiten.SetWindowSize(1000, 1000)
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
