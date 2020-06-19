package invad

import(
	"github.com/hajimehoshi/ebiten"
)

type ExplosionHolder struct{
	explosions []*Explosion
	sH *SpriteHolder
	spritenum uint64
}

func NewExplosionHolder(sH *SpriteHolder, spritenum uint64) *ExplosionHolder{
	eH := &ExplosionHolder{}
	eH.sH = sH
	eH.spritenum = spritenum
	return eH
}

func (eH *ExplosionHolder)SpawnExplosion(x,y float64){
	eH.explosions = append(eH.explosions, NewExplosion(x,y, eH.sH, eH.spritenum))
}


func (eH *ExplosionHolder)UpdateExplosions(){
	for i := 0 ;i < len(eH.explosions); i++{
		if !eH.explosions[i].IsAlive(){
			a:= eH.explosions[0:i]
			b:= eH.explosions[i+1:]
			eH.explosions = append(a, b...)
		}
	}
	for i:= 0 ; i < len(eH.explosions); i++{
		eH.explosions[i].SetDuration()
	}
}
	
func (eH ExplosionHolder)DrawExplosions(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	for i:= 0; i<len(eH.explosions); i++{
		eH.explosions[i].Draw(screen, options)
	}
}
