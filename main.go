package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH = 640;
	HEIGHT = 480;
	SCREEN_WIDTH = 640;
	SCREEN_HEIGHT = 480;
	TITLE = "Ebiten Go Workshop | Chonlatee"
)

type Game struct {}

func (g *Game) Update() error {
	return nil;
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT;
}

func main() {
	game := &Game{};
	ebiten.SetWindowSize(WIDTH, HEIGHT);
	ebiten.SetWindowTitle(TITLE);

	err := ebiten.RunGame(game);
	if err != nil {
		panic(err);
	}
}