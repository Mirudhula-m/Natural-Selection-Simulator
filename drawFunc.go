package main

import (
	// "math"
	// "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)


// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	
	if g.mode == ModeGame {

		w, h := platformImg.Size()

		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(1, 1, 1, 0.5)
		// Geometric adjustments of images
		g.op.GeoM.Reset()
		g.op.GeoM.Scale(scaleXplatform,scaleYplatform)
		// Translation of image from current position
		g.op.GeoM.Translate(-scaleXplatform*float64(w)/2, - scaleYplatform*float64(h)/2)
		g.op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)
		screen.DrawImage(platformImg, &g.op)		


		for k := 0; k < g.populations.numTypes; k++ {
			// ----------------------------------
			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			// Draw Sprite - for each population 
			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			// ----------------------------------
			for i := 0; i < g.populations.population[k].sprites.num; i++ {
				w, h := spriteImgs[k].Size()

				op := &ebiten.DrawImageOptions{}
				op.ColorM.Scale(0, 0, 0, 0.5)

				s := g.populations.population[k].sprites.sprites[i]
				scaleFactor := g.populations.population[k].scaleFactor

				g.op.GeoM.Reset()
				g.op.GeoM.Scale(0.2*scaleFactor, 0.2*scaleFactor)
				g.op.GeoM.Translate(-0.2*float64(w)/2, -0.2*float64(h)/2)
				g.op.GeoM.Translate(float64(s.x), float64(s.y))
				screen.DrawImage(spriteImgs[k], &g.op)		
			}
		}
		

		// ----------------------------------
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// Draw Food
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// ----------------------------------
		for j := 0; j < g.foods.num; j++ {
			w, h := foodImg.Size()

			op := &ebiten.DrawImageOptions{}
			op.ColorM.Scale(1, 1, 1, 0.35)

			s := g.foods.foods[j]

			if s.present == true {
				g.op.GeoM.Reset()
				g.op.GeoM.Scale(0.05, 0.05)
				g.op.GeoM.Translate(-0.05*float64(w)/2, -0.05*float64(h)/2)
				g.op.GeoM.Translate(float64(s.x), float64(s.y))
				screen.DrawImage(foodImg, &g.op)
			}
					
		}
	}
	
	// ----------------------------------
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Draw Text
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// ----------------------------------

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"NATURAL SELECTION SIMULATOR"}
		texts = []string{"", "", "", "", "", "", "", "1.ADD POPULATION", "", "2.START SIMULATION"}
	
	case ModeAddPopulation:
		titleTexts = []string{"Type the index of traits to add to population.."}
		texts = []string{"", "", "", "", "", "", "", "1.Increased speed", "", "2.Increased Size", "", "3.Increased appetite", "", "4.Return to main menu"}
	}
	for i, l := range titleTexts {
		x := (screenWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*titleFontSize, color.White)
	}
	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.White)
	}




}
