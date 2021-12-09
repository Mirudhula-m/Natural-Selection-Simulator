package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"time"
)


// Image initialization related
var (
	spriteImgs	[]*ebiten.Image
	platformImg	*ebiten.Image
	foodImg		*ebiten.Image
)

// timing
var start	time.Time

// Initial parameters
var (
	numSprites = 10
	numFoods = 100
)

// Font-related
var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

// Scaling-related
var (
	scaleXplatform = 0.98
	scaleYplatform = 0.98

	thresholdEnergy = 3
)


var (
	keys = []ebiten.Key{
		ebiten.Key1,
		ebiten.Key2,
		ebiten.Key3,
		ebiten.Key4,
		ebiten.Key0,
	}
)


var popNums [][]float64

