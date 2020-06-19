package invad

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"image"
)

//Subpicture is data on a picture within a image
type Subpicture struct{
	L,U,R,D int
}

//NewSubpicture creates a subpicture
func NewSubpicture(l,u,r,d int) *Subpicture{
	return &Subpicture{l,u,r,d}
}

//GetRectangle returns the image.Rectangle of the Subpicture
func (sp Subpicture)GetRectangle() image.Rectangle{
	return image.Rect(sp.L, sp.U, sp.R, sp.D)
}

//GetRectangle2D returns the Rectangle2D a Sprite covers
func (sp Subpicture)GetRectangle2D() *Rectangle2D{
	return NewRectangle2D(float64(sp.L), float64(sp.U), float64(sp.R), float64(sp.D))
}

//SpriteHolder holds all images of the game and subpicture on these images
type SpriteHolder struct{
	subpictures map[uint64]*Subpicture
	images map[uint64]*ebiten.Image
	tilesheet map[uint64]uint64
}

//NewSpriteHolder creates a new SpriteHolder
func NewSpriteHolder() *SpriteHolder{
	sH := &SpriteHolder{}
	sH.subpictures = make(map[uint64]*Subpicture)
	sH.images = make(map[uint64]*ebiten.Image)
	sH.tilesheet = make(map[uint64]uint64)
	return sH
}

//AddImage adds an image to the spriteholder
func (sH *SpriteHolder)AddImage(imagenum uint64, filename string){
	img,_,err:= ebitenutil.NewImageFromFile(filename, ebiten.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}
	sH.images[imagenum]= img
}
//AddSubpicture adds a subpicture at subpicturenumber on imagenumber to the spriteholder
func (sH *SpriteHolder)AddSubpicture(subpicturenumber uint64, imagenumber uint64, subpicture *Subpicture){
	sH.tilesheet[subpicturenumber] = imagenumber
	sH.subpictures[subpicturenumber] = subpicture
}

//DisplaySprite returns the ebiten Image of the Sprite
func (sH SpriteHolder)DisplaySprite(subpicturenumber uint64) *ebiten.Image{
	return sH.images[sH.tilesheet[subpicturenumber]].SubImage(sH.subpictures[subpicturenumber].GetRectangle()).(*ebiten.Image)
}

func (sH SpriteHolder)GetSpriteRectangle2D(subpicturenumber uint64) *Rectangle2D{
	return sH.subpictures[subpicturenumber].GetRectangle2D()
}
