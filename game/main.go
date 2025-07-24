package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"strconv"
	"strings"

	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var sprite_car = load_image("./assets/car.png")
var sprite_bg = load_image("./assets/bg.png")
var Terminated = errors.New("terminated")

type Game struct {
	car_pos   Vector
	init_pos  Vector
	speed     float64
	last_time time.Time
	serial    Serial
	cooldown  time.Duration
}
type Vector struct {
	x   float64
	y   float64
	rot float64
}

type Info struct {
	dir    string
	speed  float64
	factor int
}

func load_image(name string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		panic(err)
	}

	return img
}

func process_info(i string) Info {
	if i == "" {
		return Info{dir: ""}
	}
	var out Info
	var err error
	params := strings.Split(i, " ")

	out.dir = strings.TrimPrefix(params[0], "D:")
	out.speed, err = strconv.ParseFloat(strings.TrimPrefix(params[1], "S:"), 64)
	if err != nil {
		log.Print("error in parsing data from serial:", i+": ", err)
		return Info{dir: ""}
	}
	out.factor, err = strconv.Atoi(strings.TrimPrefix(params[2], "F:"))
	if err != nil {
		log.Print("error in parsing data from serial:", i+": ", err)
		return Info{dir: ""}
	}
	return out
}

func (g *Game) update_movement(dir string) error {
	if dir == "L" {
		g.car_pos.x -= g.speed
	} else if dir == "R" {
		g.car_pos.x += g.speed
	}

	g.car_pos.y -= g.speed

	return nil

}

func (g *Game) Update() error {

	var info Info
	if time.Now().Sub(g.last_time) > g.cooldown {
		g.last_time = time.Now()
		info = process_info(receive_from_serial(g.serial))
		log.Print("RECIEVED!")
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return Terminated
	}
	g.update_movement(info.dir)

	if info.dir != "" {
		g.speed = info.speed
	}

	return nil
}

func Draw_background(screen *ebiten.Image, g *Game) {

	/* prallax bg window to make it infinte */

	if sprite_bg == nil {
		screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
		return
	}

	bg_width := float64(sprite_bg.Bounds().Dx())
	bg_height := float64(sprite_bg.Bounds().Dy())
	screen_width, screen_height := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	parallax_factor := 0.5

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

	Draw_background(screen, g)

	str := fmt.Sprintf("CONNECTED TO SERIAL: %s(%d)\nSPEED: %f\n", g.serial.port, g.serial.rate, g.speed)
	ebitenutil.DebugPrint(screen, str)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.init_pos.x, g.init_pos.y)
	opts.GeoM.Scale(0.2, 0.2)

	screen.DrawImage(sprite_car, opts)

}

func (g *Game) Layout(width, height int) (x, y int) {
	return width, height
}

func main() {

	g := &Game{
		car_pos:   Vector{x: 1800, y: 4000, rot: 0},
		init_pos:  Vector{x: 1800, y: 3000, rot: 0},
		speed:     0.0,
		serial:    open_serial(),
		cooldown:  time.Duration(10) * time.Millisecond,
		last_time: time.Now(),
	}

	defer g.serial.socket.Close()
	ebiten.SetWindowSize(1000, 1000)
	err := ebiten.RunGame(g)

	if err != nil {
		if err == Terminated {
			return
		}
		panic(err)
	}
}
