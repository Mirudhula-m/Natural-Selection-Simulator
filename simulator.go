/*****************************************************************
 *****************************************************************
Natural Selection Simulator - SURVIVAL OF THE FITTEST!

In this example demo, we have three different traits:
Speed
Size
Appetite

These traits can be combined multiple ways to form the overall character of a population.
We can form the populations with desired traits and run the simulation to see who survives!

 ***************************************************************** 
******************************************************************/

package main

import (
	"log"
	"math/rand"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
)


// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Set window size - width and height
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Natural Selection Simulator")
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowResizable(true)

	// Game is looped till user quits
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}