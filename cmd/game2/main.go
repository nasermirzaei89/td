package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2"
	"github.com/pkg/errors"
)

type drawer struct {
	game2.BaseObject
	bound  image.Rectangle
	active bool
}

func main() {
	g := game2.NewGame()

	d := new(drawer)

	g.AddObject(d)

	frame := 0

	g.On(game2.EventNameGameStart, func(ev game2.Event) {
		log.Println(ev.Name())
	})

	g.On(game2.EventNameGameEnd, func(ev game2.Event) {
		log.Println(ev.Name())
	})

	g.On(game2.EventNameMouseLeftJustPress, func(ev game2.Event) {
		eventGlobalClick, _ := ev.(*game2.EventMouseLeftJustPress)

		d.bound.Min.X = eventGlobalClick.MouseX
		d.bound.Min.Y = eventGlobalClick.MouseY

		d.active = true
	})

	g.On(game2.EventNameMouseLeftPress, func(ev game2.Event) {
		eventGlobalClick, _ := ev.(*game2.EventMouseLeftPress)

		d.bound.Max.X = eventGlobalClick.MouseX
		d.bound.Max.Y = eventGlobalClick.MouseY
	})

	g.On(game2.EventNameMouseLeftJustRelease, func(ev game2.Event) {
		d.active = false
	})

	g.OnUpdate(func() {
		frame++

		if frame == 100 {
			g.RemoveObject(d)
		}
	})

	g.OnDraw(func(ev *game2.EventDraw) {
		ebitenutil.DebugPrint(ev.Screen, fmt.Sprintf("Frame %d", frame))
	})

	d.OnDraw(func(ev *game2.EventDraw) {
		if !d.active {
			return
		}

		ebitenutil.DrawRect(ev.Screen, float64(d.bound.Min.X), float64(d.bound.Min.Y), float64(d.bound.Size().X), float64(d.bound.Size().Y), color.White)
	})

	d.OnCreate(func(ev *game2.EventCreate) {
		log.Println(ev.Name())
	})

	d.OnDestroy(func(ev *game2.EventDestroy) {
		log.Println(ev.Name())
	})

	if err := g.Run(); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
