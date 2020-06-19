package invad

import (
	"github.com/hajimehoshi/ebiten"
)

//Invader is the opponent of the player
type Invader struct{
	Position *Point2D
	Area *Rectangle2D
	Sprites *SpriteHolder
	Spritenum uint64
	Alive bool
}

//NewInvader creates an Invader
func NewInvader(x,y float64, sH *SpriteHolder, spritenum uint64) *Invader{
	i := &Invader{}
	i.Sprites = sH
	i.Area = i.Sprites.GetSpriteRectangle2D(spritenum)
	
	i.Position = NewPoint(x,y)
	i.Area.Place(i.Position)
	i.Spritenum = spritenum
	i.Alive = true
	return i
}

//Move moves an invader 
func (i *Invader)Move(x,y float64){
	i.Position.X += x
	i.Position.Y += y
	i.Area.Move(x,y)
}
func (i *Invader)GotKilled(){
	i.Alive = false
}

func (i Invader)IsAlive() bool{
	return i.Alive
}

//Draws the Invader on the screen
func (i Invader)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	if i.Alive{
		options.GeoM.Translate(i.Position.X, i.Position.Y)
		screen.DrawImage(i.Sprites.DisplaySprite(i.Spritenum), options)
		options.GeoM.Translate(-i.Position.X, -i.Position.Y)
	}
}


