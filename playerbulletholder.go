package invad

import(
	"github.com/hajimehoshi/ebiten"
		
)

//Playerbulletholder holds all bullets shot by the player
type PlayerbulletHolder struct{
	bullets []*Playerbullet
	sprites *SpriteHolder
	spritenum uint64
}

//NewPlayerbulletHolder creates a PlayerbulletHolder
func NewPlayerbulletHolder(sH *SpriteHolder, spritenum uint64) *PlayerbulletHolder{
	pbh := &PlayerbulletHolder{}
	pbh.sprites = sH
	pbh.spritenum = spritenum
	return pbh
}

//SpawnBullet spawns a bullet
func (pbh *PlayerbulletHolder)SpawnBullet(x,y float64){
	pbh.bullets = append(pbh.bullets, NewPlayerBullet(x,y, pbh.sprites, pbh.spritenum))
}

//Move moves the bullets
func (pbh *PlayerbulletHolder)Move() {
	for i:= 0; i < len(pbh.bullets); i++{
		pbh.bullets[i].Move(0,-0.6)
	}
}

//CheckCollision checks for collision with upper screenborder and invader
func (pbh *PlayerbulletHolder)CheckCollision(inv *Invader) bool{
	for i := 0; i< len(pbh.bullets); i++{
		if pbh.bullets[i].Area.Within(pbh.bullets[i].Position.X+1, 5){
			pbh.bullets[i].Collided()
		}
		if pbh.bullets[i].Area.Intersects(inv.Area){
			pbh.bullets[i].Collided()
			pbh.CheckAlive()
			return true
		}
	}	
	return false
}

//CheckAlive checks live status of bullets and removes dead bullets from bullets array
func (pbh *PlayerbulletHolder)CheckAlive(){
	for i := 0; i< len(pbh.bullets); i++{
		if !pbh.bullets[i].IsAlive(){
			a := pbh.bullets[0:i]
			b := pbh.bullets[i+1:]
			pbh.bullets = append(a,b...)
		}
	}
}

//DrawBullets draws bullest on screen
func (pbh PlayerbulletHolder)DrawBullets(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	for i:= 0; i< len(pbh.bullets); i++{
		pbh.bullets[i].Draw(screen, options)
	}
}




