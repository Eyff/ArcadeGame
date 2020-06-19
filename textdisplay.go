package invad

import (
	"image/color"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
	"io/ioutil"
	"log"
	"strconv"
)

//TextDisplay cares about displaying text
type TextDisplay struct{
	text string
	font *truetype.Font
	mplusNormalFont font.Face
	color color.Color
}

//Init sets variables of textdisplay
func (td *TextDisplay)Init(){
	fontBytes, err := ioutil.ReadFile("Blockstepped.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	var err2 error
	td.font, err2 = truetype.Parse(fontBytes)
	if err2 != nil {
		log.Fatal(err)
	}
	td.mplusNormalFont = truetype.NewFace(td.font, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	td.text = "Game Over!!!!!!!!!!!!!/n/n/nPress Enter to restart!"
}

//DrawLives draws the current lives
func (td TextDisplay)DrawLives(lives int, screen *ebiten.Image){
	life := strconv.Itoa(lives)
	text.Draw(screen, life, td.mplusNormalFont,10,30,color.White)
}

func (td TextDisplay)DrawLevel(level int, screen *ebiten.Image){
	lvl := strconv.Itoa(level)
	text.Draw(screen, lvl, td.mplusNormalFont, 550, 30, color.White)
}
//DrawGameOver draws gameover message
func (td TextDisplay)DrawGameOver(screen *ebiten.Image){
	text.Draw(screen, td.text, td.mplusNormalFont, 200,200,color.White)
}
