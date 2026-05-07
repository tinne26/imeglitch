package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type IMEBridge interface {
	Show(int32, int32)
	Hide()
}

var IME IMEBridge

type Game struct {
	touches  []ebiten.TouchID
	imeShown bool
}

func (g *Game) Layout(w, h int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	return int(math.Round(scale * float64(w))), int(math.Round(scale * float64(h)))
}

func (g *Game) Update() error {
	g.touches = g.touches[:0]
	g.touches = inpututil.AppendJustPressedTouchIDs(g.touches)
	if len(g.touches) > 0 && IME != nil {
		if g.imeShown {
			IME.Hide()
		} else {
			IME.Show(0, 0)
		}
		g.imeShown = !g.imeShown
	}

	return nil
}

func (g *Game) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.White)

	bounds := canvas.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	colors := []color.RGBA{
		{0xFF, 0x59, 0x5E, 0xFF}, {0xFF, 0xCA, 0x3A, 0xFF}, {0x8A, 0xC9, 0x26, 0xFF},
		{0x19, 0x82, 0xC4, 0xFF}, {0x6A, 0x4C, 0x93, 0xFF}, {0x42, 0xE2, 0xB8, 0xFF},
		{0xF4, 0x60, 0x36, 0xFF},
	}

	y, dy := float32(0.0), float32(height/len(colors))+1
	for _, clr := range colors {
		vector.FillRect(canvas, 0, y, float32(width), dy, clr, false)
		y += dy
	}
	vector.FillRect(canvas, 0, 0, 64, float32(height), color.RGBA{48, 64, 128, 255}, false)
	ebitenutil.DebugPrint(canvas, fmt.Sprintf("BOUNDS: %v", bounds))
}
