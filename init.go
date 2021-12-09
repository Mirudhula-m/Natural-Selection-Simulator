package main

import (
	"math/rand"
	"image"
	_ "image/png"
	"bytes"
	"log"
	_ "embed"
	// "strings"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"

	
	// go get golang.org/x/image/font/sfnt@v0.0.0-20210628002857-a66eb6448b8d //"golang.org/x/image/font/sfnt"
	// go get -u github.com/wcharczuk/go-chart
)

const (
	screenWidth  = 2000
	screenHeight = 600
	maxAngle     = 256

	tileSize         = 32
	titleFontSize    = fontSize * 1.5
	fontSize         = 24
	smallFontSize    = fontSize / 2

	MinSprites = 10
	MaxSprites = 50000

	MinFood = 1
	MaxFood = 50000

	dayDuration = 7 // in seconds

	InitialSpriteNum = 10

	MaxPopulationNum = 10
)

const (
	ModeTitle = iota
	ModeAddPopulation
	ModeGame
)


// Embedding the image into a slice of bytes
// var spriteByte []byte

//go:embed images/sprite1.png
var spriteByte1 []byte
//go:embed images/sprite2.png
var spriteByte2 []byte
//go:embed images/sprite3.png
var spriteByte3 []byte
//go:embed images/sprite4.png
var spriteByte4 []byte
//go:embed images/sprite5.png
var spriteByte5 []byte
//go:embed images/sprite6.png
var spriteByte6 []byte
//go:embed images/sprite7.png
var spriteByte7 []byte
//go:embed images/sprite8.png
var spriteByte8 []byte
//go:embed images/sprite9.png
var spriteByte9 []byte
//go:embed images/sprite10.png
var spriteByte10 []byte



//go:embed images/platform4.png
var platformByte []byte
//go:embed images/food.png
var foodByte []byte


// Initialize populations with characteristics like number of populations, etc.
func (g *Game) initPopulation() {
	defer func() {
		g.populations.inited = true
	}()
	g.populations.numTypes = 0
	g.populations.population = make([]*Population, MaxPopulationNum)
}


// This function initializes all the expecting png images that need to be decoded 
// ... into byte sized object.
func init() {
	fmt.Println("Initializing...")

	spriteImgs = make([]*ebiten.Image, InitialSpriteNum)
	for i := 0; i < InitialSpriteNum; i++ {

		switch i {

			case 0:
				// Decode an image from the image file's byte slice.
				img, _, err := image.Decode(bytes.NewReader(spriteByte1))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)

			case 1:
				img, _, err := image.Decode(bytes.NewReader(spriteByte2))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 2:
				img, _, err := image.Decode(bytes.NewReader(spriteByte3))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 3:
				img, _, err := image.Decode(bytes.NewReader(spriteByte4))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 4:
				img, _, err := image.Decode(bytes.NewReader(spriteByte5))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 5:
				img, _, err := image.Decode(bytes.NewReader(spriteByte6))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 6:
				img, _, err := image.Decode(bytes.NewReader(spriteByte7))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 7:
				img, _, err := image.Decode(bytes.NewReader(spriteByte8))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 8:
				img, _, err := image.Decode(bytes.NewReader(spriteByte9))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
			case 9:
				img, _, err := image.Decode(bytes.NewReader(spriteByte10))
				if err != nil {
					log.Fatal(err)
				}
				spriteImgs[i] = ebiten.NewImageFromImage(img)
		}
	}	

	img, _, err := image.Decode(bytes.NewReader(platformByte))
	if err != nil {
		log.Fatal(err)
	}
	platformImg = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(foodByte))
	if err != nil {
		log.Fatal(err)
	}
	foodImg = ebiten.NewImageFromImage(img)



	// From this section onwards, the initialization of text fonts and styles takes places

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})


}


// Initialize game parameters
func (g *Game) init(N_food int) {
	// This function is only entered if the intial game parameters have not
	// ... been  initialized.

	// Once first initialization is done, mark done
	defer func() {
		g.inited = true
	}()

	// ----------------------------------
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Initializing population parameters and other one time parameters
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// ----------------------------------
	for k := 0; k < g.populations.numTypes; k++ {
		if g.populations.population[k].initNum == 0 {
			g.populations.population[k].sprites.num = numSprites
			g.populations.population[k].sprites.newNum = numSprites
			g.populations.population[k].initNum = 1
			g.populations.population[k].sprites.numArr = make([]int, 0)
			g.gens = 0
		}

		// Forming the number of sprites as a per population basis per generation so that
		// ... they can be used later
		g.populations.population[k].sprites.numArr = append(g.populations.population[k].sprites.numArr, g.populations.population[k].sprites.num)
		fmt.Println(g.populations.population[k].sprites.numArr)

		// var defaultParam int = 1
		// var sumSpeed float64
		// if g.populations.population[k].speed != 2.0  {
		// 	defaultParam = 0
		// 	avgSpeed = append(avgSpeed, g.populations.population[k].speed)
		// }

		g.populations.population[k].sprites.num = g.populations.population[k].sprites.newNum
		g.populations.population[k].sprites.sprites = make([]*Sprite, g.populations.population[k].sprites.num)



		// ----------------------------------
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// Initializing sprite parameters
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// ----------------------------------
		for i := range g.populations.population[k].sprites.sprites {
			w, h := spriteImgs[k].Size()
			areaL := 700
			areaB := 370
			c := 73

			var x, y, vx, vy int
			// random number from 1 to 4 to choose which border the sprite starts from
			rNum := rand.Intn(4) + 1

			if rNum == 1 {
				// left border
				x, y =  int(screenWidth/2 - areaL/2) - c, int((rand.Float64() * float64(areaB)) + float64(screenHeight/2 - areaB/2))
				vx = 1.0
				vy = 2*rand.Intn(2)-1
			} else if rNum == 2 {
				// top border
				x, y =  int((rand.Float64() * float64(areaL)) + float64(screenWidth/2 - areaL/2)), int(screenHeight/2 - areaB/2) - c
				vy = 1
				vx = 2*rand.Intn(2)-1
			} else if rNum == 3 {
				// right border
				x, y =  int(screenWidth/2 + areaL/2) + c, int((rand.Float64() * float64(areaB)) + float64(screenHeight/2 - areaB/2))
				vx = -1
				vy = 2*rand.Intn(2)-1
			} else if rNum == 4 {
				// top border
				vy = -1
				vx = 2*rand.Intn(2)-1
				x, y =  int((rand.Float64() * float64(areaL)) + float64(screenWidth/2 - areaL/2)), int(screenHeight/2 + areaB/2) + c
			} else {
				panic("rNum incorrect!")
			}

			g.populations.population[k].sprites.sprites[i] = &Sprite{
				imageWidth:  w,
				imageHeight: h,
				x:           x,
				y:           y,
				vx:          float64(vx),
				vy:          float64(vy),
				reached:	 false,
			}
		// 	if defaultParam == 0 {
		// 		g.populations.population[k].speed = avgSpeed[len(avgSpeed)-1]
		// 		speedDiff := weightedRandom(1.5, 10)
		// 		sign := 2*rand.Intn(2)-1
		// 		g.populations.population[k].sprites.sprites[i].speed += float64(sign) * speedDiff
		// 		sumSpeed += g.populations.population[k].sprites.sprites[i].speed
		// 	}
		// }
		// if defaultParam == 0 {
		// 	avgSpeed = append(avgSpeed, avgSpeed[len(avgSpeed)-1] - (sumSpeed / float64(g.populations.population[k].sprites.num)))
		// 	fmt.Println(avgSpeed[len(avgSpeed)-1] - (sumSpeed / float64(g.populations.population[k].sprites.num)))
		}
	}

	// ----------------------------------
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Initializing food parameters
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// ----------------------------------
	g.foods.num = N_food
	g.foods.foods = make([]*Food, g.foods.num)


	for i := range g.foods.foods {
		w, h := foodImg.Size()
		areaL := 700
		areaB := 370
		x, y := int((rand.Float64() * float64(areaL)) + float64(screenWidth/2 - areaL/2)), int((rand.Float64() * float64(areaB)) + float64(screenHeight/2 - areaB/2))
		g.foods.foods[i] = &Food{
			imageWidth:  w,
			imageHeight: h,
			x:           x,
			y:           y,
			present:	 true,
		}
		
	}

	start = time.Now()

}