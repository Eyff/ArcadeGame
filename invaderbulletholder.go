package invad

import(
	"github.com/hajimehoshi/ebiten"
)

//InvaderBulletHolder holds the invaderbullets
type InvaderBulletHolder struct{
	bullets []*Invaderbullet
	spriteholder *SpriteHolder
	Spritenum uint64
}

//NewInvaderBulletHoldler creates an InvaderBulletHolder
func NewInvaderBulletHolder(sH *SpriteHolder, spritenum uint64) *InvaderBulletHolder{
	ibh := &InvaderBulletHolder{}
	ibh.spriteholder = sH
	ibh.Spritenum = spritenum
	return ibh
}

//SpawnBullet spawns a bullet
func (IBH *InvaderBulletHolder)SpawnBullet(x,y float64){
	IBH.bullets =append(IBH.bullets,NewInvaderBullet(x,y,IBH.spriteholder, IBH.Spritenum))
}

//Move moves the invaderbullets
func (IBH *InvaderBulletHolder)Move(){
	for i := 0; i < len(IBH.bullets); i++{ 
		IBH.bullets[i].Move(0,0.4)
	}
}

//checks for collision
func (IBH *InvaderBulletHolder)CheckCollision(p *Player){
	for i :=0; i< len(IBH.bullets); i++{
		if IBH.bullets[i].Area.Within(IBH.bullets[i].Position.X+1, 475){
			IBH.bullets[i].Collided()
		}
		if IBH.bullets[i].Area.Intersects(p.Area){
			IBH.bullets[i].Collided()
			p.GotHit()
		}
	}
}

//Checks if bullets are alive removes dead bullets
func (IBH *InvaderBulletHolder)CheckAlive() {
	for i := 0; i < len(IBH.bullets); i++{
		if !IBH.bullets[i].IsAlive(){
			a := IBH.bullets[0:i]
			b := IBH.bullets[i+1:]
			IBH.bullets = a
			IBH.bullets = append(IBH.bullets,b...)
		}
	}
}

//Draws bullets in invaderbulletholder
func (IBH *InvaderBulletHolder)DrawBullets(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	for i := 0; i< len(IBH.bullets);i++{
		IBH.bullets[i].Draw(screen, options)
	}
}
