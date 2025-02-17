package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	paddingpx      = 2
	cellSizepx     = 8
	screenWidth    = 1000
	screenHeight   = screenWidth
	frameThreshold = 12
)

type World struct {
	grid    []bool
	width   int
	height  int
	running bool
}

func (w *World) randomGrid() {
	for i := range w.grid {
		w.grid[i] = rand.UintN(2) == 0
	}
}

func (w *World) clearGrid() {
	for i := range w.grid {
		w.grid[i] = false
	}
}

func (w *World) fillGrid() {
	for i := range w.grid {
		w.grid[i] = true
	}
}

func (w *World) step() {
	newGrid := make([]bool, len(w.grid))
	neighbourCount := func(grid []bool, width, height, x, y int) int {
		sum := 0
		for yy := y - 1; yy <= y+1; yy++ {
			for xx := x - 1; xx <= x+1; xx++ {
				if yy >= height || xx >= width || yy < 0 || xx < 0 || (yy == y && xx == x) {
					continue
				}
				if grid[xx+yy*height] {
					sum++
				} else {
					continue
				}
			}
		}
		return sum
	}
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			neighbours := neighbourCount(w.grid, w.width, w.height, x, y)

			currentCell := &w.grid[x+y*w.height]
			newCell := &newGrid[x+y*w.height]

			if *currentCell {
				if neighbours < 2 {
					*newCell = false
				} else if neighbours > 3 {
					*newCell = false
				} else {
					*newCell = true
				}
			} else {
				if neighbours == 3 {
					*newCell = true
				}
			}
		}
	}
	w.grid = newGrid

}

type Game struct {
	world              *World
	countFramesPassed  int
	drawFrameThreshold int
	running            bool
}

func NewGame(width, height int) *Game {
	world := &World{
		grid:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	world.randomGrid()
	game := &Game{
		world:              world,
		countFramesPassed:  0,
		drawFrameThreshold: frameThreshold,
		running:            true,
	}
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return game
}

func (g *Game) Update() error {
	shouldStep := func() bool {
		if g.countFramesPassed >= g.drawFrameThreshold {
			g.countFramesPassed = 0
			return true
		}
		g.countFramesPassed++
		return false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.running = !g.running
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.world.fillGrid()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.world.clearGrid()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.world.randomGrid()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		g.drawFrameThreshold++
		fmt.Println(g.drawFrameThreshold)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		if g.drawFrameThreshold > 0 {
			g.drawFrameThreshold = g.drawFrameThreshold - 1
		}
		fmt.Println(g.drawFrameThreshold)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		yPos, xPos = (xPos / (paddingpx + cellSizepx)), (yPos / (paddingpx + cellSizepx))
		g.world.grid[xPos+yPos*g.world.height] = !g.world.grid[xPos+yPos*g.world.height]
		fmt.Println(xPos, yPos)
	}
	if shouldStep() && g.running {
		g.world.step()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x22, 0x22, 0x22, 0xff})
	for y := 0; y < g.world.height; y++ {
		for x := 0; x < g.world.width; x++ {
			if g.world.grid[x+y*g.world.height] {
				vector.DrawFilledRect(
					screen,
					float32(y*(cellSizepx+paddingpx)+paddingpx/2),
					float32(x*(cellSizepx+paddingpx)+paddingpx/2),
					float32(cellSizepx),
					float32(cellSizepx),
					color.RGBA{
						0xff,
						0xff,
						0xff,
						0xff,
					},
					false,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.world.width * (cellSizepx + paddingpx), g.world.height * (cellSizepx + paddingpx)
}

func main() {
	game := NewGame(50, 50)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
