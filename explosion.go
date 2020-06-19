package invad

import(
	"github.com/hajimehoshi/ebiten"
	"time"
)


//Explosion happens when bullets collide with player or opponent
type Explosion struct{
	Position *Point2D
	Area *Rectangle2D
	Sprites *SpriteHolder
	Spritenum uint64
	duration time.Duration
	timestart time.Time
}

//NewExplosion creates a new Explosion
func NewExplosion(x,y float64, sH *SpriteHolder, spritenum uint64) *Explosion{
	exp := &Explosion{}
	exp.Position = NewPoint(x,y)
	exp.Area = sH.GetSpriteRectangle2D(spritenum)
	exp.Area.Place(exp.Position)
	exp.Sprites = sH
	exp.Spritenum = spritenum
	exp.timestart = time.Now()
	return exp
}

//IsAlive checks if explosion is alive
func (exp Explosion)IsAlive() bool{
	return exp.duration.Seconds() < 1.5
}

//SetDuration sets explosion duration based on timestart
func (exp *Explosion)SetDuration(){
	exp.duration = time.Since(exp.timestart)
}

//Draws explosions on screen
func (exp Explosion)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	options.GeoM.Translate(exp.Position.X, exp.Position.Y)
	screen.DrawImage(exp.Sprites.DisplaySprite(exp.Spritenum), options)
	options.GeoM.Translate(-exp.Position.X, -exp.Position.Y)
}
