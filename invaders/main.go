package main

import(
	"github.com/hajimehoshi/ebiten"
	"log"
	"github.com/Eyff/ArcadeGame"
)

// Game implements ebiten.Game interface.
type Game struct{
	sH *invad.SpriteHolder
	pl *invad.Player
	pbh *invad.PlayerbulletHolder
	exh *invad.ExplosionHolder
	op *ebiten.DrawImageOptions
	iH *invad.InvaderHolder
	ibh *invad.InvaderBulletHolder
	td *invad.TextDisplay
	lvl int
	loaded bool
}

// Initialize sets the values of variables in game structure
func (g *Game)Initialize(){
	g.op = &ebiten.DrawImageOptions{}
	g.sH = invad.NewSpriteHolder()
	g.sH.AddImage(0,"sprites.png")
	g.sH.AddSubpicture(0,0,invad.NewSubpicture(1,2,16,14))
	g.sH.AddSubpicture(1,0,invad.NewSubpicture(28,3,44,20))
	g.sH.AddSubpicture(2,0,invad.NewSubpicture(54,5,57,15))
	g.sH.AddSubpicture(3,0,invad.NewSubpicture(65,6,76,16))
	g.sH.AddSubpicture(4,0,invad.NewSubpicture(0,41,639,45))
	g.sH.AddSubpicture(5,0,invad.NewSubpicture(91,1,118,19))
	g.pbh = invad.NewPlayerbulletHolder(g.sH,3)
	g.exh = invad.NewExplosionHolder(g.sH,5)
	g.pl = invad.NewPlayer(312,400,g.sH, 1, 3,g.pbh,g.exh)
	g.lvl= 1
	g.ibh = invad.NewInvaderBulletHolder(g.sH,2)
	g.iH = invad.NewInvaderHolder(g.ibh,g.pbh,g.exh)
	g.iH.Spawn(g.lvl, g.sH)
	g.td = &invad.TextDisplay{}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	if !g.loaded{
			g.td.Init()
			g.loaded = true
		}
	g.pl.Steer(ebiten.IsKeyPressed(ebiten.KeyLeft), ebiten.IsKeyPressed(ebiten.KeyRight), ebiten.IsKeyPressed(ebiten.KeySpace))
	g.iH.Move()
	g.ibh.Move()
	g.pbh.Move()
	g.pbh.CheckAlive()
	g.ibh.CheckCollision(g.pl)
	g.ibh.CheckAlive()
	if g.iH.CheckAlive(){
		g.lvl++
		g.iH.Spawn(g.lvl, g.sH)
	}
	g.exh.UpdateExplosions()
    // Write your game's logical update.
    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.pl.Draw(screen, g.op)
	g.iH.Draw(screen, g.op)
	g.pbh.DrawBullets(screen, g.op)
	g.ibh.DrawBullets(screen,g.op)
	g.exh.DrawExplosions(screen,g.op)
	g.td.DrawLives(g.pl.Lives,screen)
	g.td.DrawLevel(g.lvl, screen)
	if g.pl.Lives <= 0 {
		g.td.DrawGameOver(screen)
	}
    // Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}

func main() {
    game := &Game{}
    // Sepcify the window size as you like. Here, a doulbed size is specified.
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Invaders must die")
    game.Initialize()
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
