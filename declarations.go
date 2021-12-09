package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ebitenSprite *ebiten.Image
	ebitenPlatform *ebiten.Image
)


type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          float64
	vy          float64
	energy		float64 // this is the food collected
	reached		bool
	speed		float64
}

type Sprites struct {
	sprites   []*Sprite
	num       int
	newNum	  int
	numArr    []int
}

type Population struct {
	energyCap	float64
	scaleFactor	float64
	speed		float64
	sprites		Sprites
	initNum		int

}

type Populations struct {
	inited		bool
	numTypes    int
	population  []*Population
}

type Food struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	present		bool
}

type Foods struct {
	foods		[]*Food
	num			int
}

type Game struct {
	populations  Populations
	foods	 	 Foods
	inited    	 bool	
	op       	 ebiten.DrawImageOptions
	keys 	 	 []ebiten.Key
	mode	 	 int
	gens         int
}



