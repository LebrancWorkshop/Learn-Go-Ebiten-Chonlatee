package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WIDTH = 640;
	HEIGHT = 480;
	SCREEN_WIDTH = 640;
	SCREEN_HEIGHT = 480;
	PLAYER_WIDTH = 50;
	PLAYER_HEIGHT = 80;
	PLAYER_SPEED = 3;
	TITLE = "Ebiten Go Workshop | Chonlatee"
)

type Game struct {
	player *ebiten.Image
	playerPostX float64
	playerPostY float64
}

func (g *Game) Update() error {

	// Keyboard Function - COMING SOON. 

	// g.playerPostX += 1;
	// if g.playerPostX + float64(PLAYER_WIDTH) >= float64(SCREEN_WIDTH) {
	// 	g.playerPostX = float64(SCREEN_WIDTH) - float64(PLAYER_WIDTH); 
	// }
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
	game.playerPostX = float64(SCREEN_WIDTH / 2) - float64(PLAYER_WIDTH / 2);
	game.playerPostY = float64(SCREEN_HEIGHT / 2);

	err = ebiten.RunGame(game);
	if err != nil {
		panic(err);
	}
}