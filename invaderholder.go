package invad

import(
	"github.com/hajimehoshi/ebiten"
	"time"
	"math/rand"
	)

//Invaderholder holds the fleet of Invaders
type InvaderHolder struct{
	ibh *InvaderBulletHolder
	pbh *PlayerbulletHolder
	exh *ExplosionHolder
	invaders []*Invader
	timeNow time.Time
	timeElapsed time.Duration
	moving bool
	rightmove bool
	firemax int
	currentfire int
}

func NewInvaderHolder(ibh  *InvaderBulletHolder,pbh *PlayerbulletHolder, exh *ExplosionHolder) *InvaderHolder{
	iH := &InvaderHolder{}
	iH.ibh = ibh
	iH.pbh = pbh
	iH.exh = exh
	return  iH
}

//Spawn creates a fleet of invaders
func (iH *InvaderHolder)Spawn(level int, sH *SpriteHolder){
	for i := 0; i < 8*level; i++{ 
		x := float64(i*30+200 -i/8*240)
		y:= float64(i/8*30 + 40)
		iH.invaders = append(iH.invaders, NewInvader(x,y,sH,0))
	}
	iH.firemax = level
	iH.timeNow = time.Now()
	iH.moving = true
}

//Draw draws the Invader fleet
func (iH InvaderHolder)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	for i := 0; i < len(iH.invaders); i++{
		iH.invaders[i].Draw(screen, options)
	}
}

//IsAlive checks if any Invader in the fleet is left
func (iH InvaderHolder)IsAlive() bool{
	for i := 0; i < len(iH.invaders); i++{
		if iH.invaders[i].Alive{
			return true
		}
	}
	return false
}

//Move moves the fleet of invaders
func (iH *InvaderHolder)Move(){
	for i := 0; i < len(iH.invaders); i++{
		if iH.pbh.CheckCollision(iH.invaders[i]){
			iH.invaders[i].GotKilled()
		}
	}
	iH.CheckAlive()
	if len(iH.invaders) > 0 {
	blockLeft:= iH.invaders[0].Area.Within(2,iH.invaders[0].Position.Y+5)
	var j int
	if len(iH.invaders)+1 / 8 > 0{
	for i := 0; i < len(iH.invaders); i++{
		if i+1 %8 == 0{
			j = i
		}
	}
	} else {
		j = len(iH.invaders)-1
	}
	blockRight:= iH.invaders[j].Area.R < 580 && iH.invaders[j].Area.R > 560
	iH.timeElapsed = time.Now().Sub(iH.timeNow)
	if iH.timeElapsed.Seconds() > 3{
		iH.timeNow = time.Now()
		iH.moving = !iH.moving
	}
	if iH.moving {
		iH.currentfire = 0
		if blockLeft {
			iH.rightmove = true
		}
		if blockRight { 
			iH.rightmove = false
		}
		if iH.rightmove{
			for i := 0; i < len(iH.invaders);i++{
				iH.invaders[i].Move(0.3,0)
			}
		}
		if !iH.rightmove {
			for i := 0;i < len(iH.invaders);i++{
				iH.invaders[i].Move(-0.3,0)
			}
		}
		
	}
	if !iH.moving{
		rand.Seed(time.Now().UnixNano())
		for i:= 0; i < len(iH.invaders); i++{
			a:= rand.Intn(100)
			if a < 10 && iH.currentfire < iH.firemax{
				iH.ibh.SpawnBullet(iH.invaders[i].Position.X+5, iH.invaders[i].Position.Y+10)
				iH.currentfire++
			}
		}
	}
}
	
}
//Wipe cleans Invaders array and respawns with lvl 1
func (iH *InvaderHolder)Wipe(sH *SpriteHolder){
	iH.invaders = nil
	iH.Spawn(1,sH)
}

//CheckAlive checks if invaders are alive and if none is left returns true
func (iH *InvaderHolder)CheckAlive() bool{
	if len(iH.invaders)== 0{
		return true
	}
	if len(iH.invaders)!= 0{
	for i := 0; i < len(iH.invaders); i++{
		if !iH.invaders[i].IsAlive(){
			iH.exh.SpawnExplosion(iH.invaders[i].Position.X-7, iH.invaders[i].Position.Y-3)
			a:= iH.invaders[0:i]
			if i+1 <= len(iH.invaders){
				b:= iH.invaders[i+1:]
				iH.invaders = append(a, b...)
				return false
			}
			iH.invaders = a
		}
	}
	}
	return false
}
