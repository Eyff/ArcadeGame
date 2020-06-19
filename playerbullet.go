package invad

import(
	"github.com/hajimehoshi/ebiten"
)

//Playerbullet is the bullets player release
type Playerbullet struct{
	Position *Point2D
	Area *Rectangle2D
	Sprites *SpriteHolder
	Spritenum uint64
	Alive bool
}

//NewPlayerBullet creates a PlayerBullet
func NewPlayerBullet(x,y float64, sH *SpriteHolder, spritenum uint64) *Playerbullet{
	pb := &Playerbullet{}
	pb.Position = NewPoint(x,y)
	pb.Area = sH.GetSpriteRectangle2D(spritenum)
	pb.Area.Place(pb.Position)
	pb.Sprites = sH
	pb.Spritenum = spritenum
	pb.Alive = true
	return pb
}

//Move moves the player bullet
func (pb *Playerbullet)Move(x,y float64){
	pb.Position.X +=x
	pb.Position.Y +=y
	pb.Area.Move(x,y)
}

//IsAlive gets alive status of playerbullet
func (pb Playerbullet)IsAlive() bool{
	return pb.Alive
}

//Colided tells bullet that it collided and died
func (pb *Playerbullet)Collided(){
	pb.Alive = false
}

//Draw draws playerbullet on screen
func (pb Playerbullet)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	options.GeoM.Translate(pb.Position.X, pb.Position.Y)
	screen.DrawImage(pb.Sprites.DisplaySprite(pb.Spritenum),options)
	options.GeoM.Translate(-pb.Position.X, -pb.Position.Y)
}
		
