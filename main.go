package main

import (
	"log"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WIDTH = 640;
	HEIGHT = 480;
	SCREEN_WIDTH = 640;
	SCREEN_HEIGHT = 480;
	TITLE = "Ebiten Go Workshop | Chonlatee"
)

type Game struct {
	player *ebiten.Image
	playerPostX float64
	playerPostY float64
}

func (g *Game) Update() error {
	g.playerPostX += 1;
	return nil;
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{};
	op.GeoM.Translate(g.playerPostX, g.playerPostY);
	screen.DrawImage(g.player, op);
} 

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT;
}

func main() {
	game := &Game{};
	ebiten.SetWindowSize(WIDTH, HEIGHT);
	ebiten.SetWindowTitle(TITLE);

	img, _, err := ebitenutil.NewImageFromFile("ship.png");
	if err != nil {
		log.Fatal(err);
	}

	game.player = img;
	game.playerPostX = float64(SCREEN_WIDTH / 2) - 25;
	game.playerPostY = float64(SCREEN_HEIGHT / 2);

	err = ebiten.RunGame(game);
	if err != nil {
		panic(err);
	}
}