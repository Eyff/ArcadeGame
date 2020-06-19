package invad

import (
	"github.com/hajimehoshi/ebiten"
	)

//Invaderbullet is a shot from an invader
type Invaderbullet struct{
	Position *Point2D
	Area *Rectangle2D
	Sprites *SpriteHolder
	Spritenum uint64
	Alive bool
}

//NewInvaderBullet creates a new Invaderbullet
func NewInvaderBullet(x,y float64, sH *SpriteHolder, spritenum uint64) *Invaderbullet{
	ab := &Invaderbullet{}
	ab.Position = NewPoint(x,y)
	ab.Area = sH.GetSpriteRectangle2D(spritenum)
	ab.Area.Place(ab.Position)
	ab.Sprites = sH
	ab.Spritenum = spritenum
	ab.Alive = true
	return ab
}

//Move mioves the invaderbullet
func (ib *Invaderbullet)Move(x,y float64){
	ib.Position.X += x
	ib.Position.Y += y
	ib.Area.Move(x,y)
}


//IsAlive gets if a bullet is still alive
func (ib Invaderbullet)IsAlive() bool{
	return ib.Alive
}

//Collided tells bullet that it collided and died
func (ib *Invaderbullet)Collided(){
	ib.Alive = false
}

//Draw draws the invaderbullet on the screen
func (ib Invaderbullet)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	options.GeoM.Translate(ib.Position.X, ib.Position.Y)
	screen.DrawImage(ib.Sprites.DisplaySprite(ib.Spritenum), options)
	options.GeoM.Translate(-ib.Position.X, -ib.Position.Y)
}
