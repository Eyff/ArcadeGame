package invad

import(
	"github.com/hajimehoshi/ebiten"
	"time"
)

//Player is the playersprite of the game
type Player struct{
	pbh *PlayerbulletHolder
	exh *ExplosionHolder
	Position *Point2D
	Area *Rectangle2D
	SpriteNum uint64
	Sprites *SpriteHolder
	Lives int
	timenow time.Time
}

//NewPlayer creates a new playersprite of the game
func NewPlayer(x,y float64, sh *SpriteHolder, spritenum uint64, lives int, pbh *PlayerbulletHolder, exh *ExplosionHolder) *Player{
	p := &Player{}
	p.Sprites = sh
	p.Area = sh.GetSpriteRectangle2D(spritenum)
	p.Position = NewPoint(x,y)
	p.Area.Place(p.Position)
	p.SpriteNum = spritenum
	p.Lives = lives
	p.pbh = pbh
	p.timenow = time.Now()
	p.exh = exh
	return p
}

//Moves the playersprite rigth or left
func (p *Player)Steer(leftkey bool, rightkey bool,spacekey bool){
	if leftkey && p.Position.X > 0{
		p.Position.X -= 0.6
		p.Area.Move(-0.6,0)
	}
	if rightkey && p.Position.X < 584{
		p.Position.X += 0.6
		p.Area.Move(0.6,0)
	}
	if spacekey && (time.Since(p.timenow).Seconds() > 2){
		p.timenow = time.Now()
		p.pbh.SpawnBullet(p.Position.X+7,p.Position.Y)
	}
}

func (p *Player)GotHit(){
	p.exh.SpawnExplosion(p.Position.X-5,p.Position.Y-3)
	p.Lives--
}

//Draws the playersprite on the screen
func (p Player)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	if p.Lives > 0{
		options.GeoM.Translate(p.Position.X, p.Position.Y)
		screen.DrawImage(p.Sprites.DisplaySprite(p.SpriteNum), options)
		options.GeoM.Translate(-p.Position.X, -p.Position.Y)
	}
}
