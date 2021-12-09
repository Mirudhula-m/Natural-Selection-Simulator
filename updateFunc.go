package main


import (
	// "github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"fmt"
	"os"

	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)


// Default population parameters
func (g *Game) SetDefaultPop(key int) {

	g.populations.population[key] = &Population {
		energyCap: 		4,
		scaleFactor:  	1.0,
		speed:			2.0,
	}
	g.populations.numTypes++
}


// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	// If populations is already initialized, don't re-initialize
	if !g.populations.inited {
		g.initPopulation()
		g.SetDefaultPop(0)
	}
	switch g.mode {
		// This mode displays the title page
		case ModeTitle:
			for _, key := range keys {
				if inpututil.IsKeyJustPressed(key) {
					if key == ebiten.Key1 {
						g.mode = ModeAddPopulation
						g.SetDefaultPop(g.populations.numTypes)

					} else {
						g.mode = ModeGame
					}
				}
			}

			// This mode helps in choosing population traits
		case ModeAddPopulation:
			for _, key := range keys {
				// Based on the key that is pressed, the specific trait is introduced
				// ... in the population
				if inpututil.IsKeyJustPressed(key) {
					if key == ebiten.Key1 {
						g.populations.population[g.populations.numTypes-1] = &Population{
							speed: 			3.0,
							energyCap: 		g.populations.population[g.populations.numTypes-1].energyCap,
							scaleFactor:  	g.populations.population[g.populations.numTypes-1].scaleFactor,
						}

					} else  if key == ebiten.Key2 {
						g.populations.population[g.populations.numTypes-1] = &Population{
							scaleFactor: 1.5,
							speed: 		 g.populations.population[g.populations.numTypes-1].speed,
							energyCap:   g.populations.population[g.populations.numTypes-1].energyCap,
						}
					
					} else if key == ebiten.Key3 {
						g.populations.population[g.populations.numTypes-1] = &Population{
							energyCap:     10,
							speed: 		   g.populations.population[g.populations.numTypes-1].speed,
							scaleFactor:   g.populations.population[g.populations.numTypes-1].scaleFactor,
						}
					
					} else {
						g.mode = ModeTitle
					}
				}
			}

			// This mode starts the game
		case ModeGame:
			
		    if !g.inited {
				g.init(numFoods)
			}

			for i := 0; i < g.populations.numTypes; i++ {
				g.UpdateGame()
				g.populations.population[i].sprites.Update(g.populations.population[i].speed, g.populations.population[i].energyCap)
			}
			

			end := time.Since(start).Seconds()
			// fmt.Println(end)

			if end >= dayDuration {
				g.EndDay()
				g.gens++
				g.inited = false
			}

			// This section helps in exiting the game and to go back to the title page
			// Once the zero key is pressed, the game is stopped, parameters are re-initialized,
			// ... and and the final number in population is calculated. Then, the plots are generated
			for _, key := range keys {
				if inpututil.IsKeyJustPressed(key) {
					if key == ebiten.Key0 {
						g.mode = ModeTitle
						g.inited = false
						g.populations.inited = false

						XValues := make([]float64, g.gens)
						popNums = make([][]float64, MaxPopulationNum)
						for k := 0; k < g.populations.numTypes; k++ {
							popNums[k] = make([]float64, g.gens)
							for i := 0; i < g.gens; i++ {
								XValues[i] = float64(i + 1)
								popNums[k][i] = float64(g.populations.population[k].sprites.numArr[i])
							}

						}
						// fmt.Println(popNums)
						len := g.populations.numTypes
						left := MaxPopulationNum - len
						fmt.Println(len, left)
						for j := 0; j < left; j++ {

							popNums[len+j] = make([]float64,g.gens)
							for l := 0; l < g.gens; l++ {
								popNums[len+j][l] = 0
							}
						}
						g.plotPopNum(XValues, popNums,)
					}					
				}
			}
	}
	
	return nil
}


func (g *Game) plotPopNum(X []float64, Y [][]float64) {
	fmt.Println(X, Y)
	graph := chart.Chart{
		Title: "Population numbers over generations",
		XAxis: chart.XAxis{
			Name: "Number of generations",
		},
		YAxis: chart.YAxis{
			Name: "Number of specimens in the population",
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:    50,
				Left:   25,
				Right:  25,
				Bottom: 10,
			},
			FillColor: drawing.ColorFromHex("efefef"),
		},
	    Series: []chart.Series{
	    		chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					},
			    	Name: "Control",
			        XValues: X,
			        YValues: Y[0],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(1).WithAlpha(64),
					},
			    	Name: "Light Red Population",
			        XValues: X,
			        YValues: Y[1],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(2).WithAlpha(64),
					},
			    	Name: "Light Pink",
			        XValues: X,
			        YValues: Y[2],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(3).WithAlpha(64),
					},
			    	Name: "Light Blue",
			        XValues: X,
			        YValues: Y[3],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(4).WithAlpha(64),
					},
			    	Name: "Light Green",
			        XValues: X,
			        YValues: Y[4],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(5).WithAlpha(64),
					},
			    	Name: "Dark Purple",
			        XValues: X,
			        YValues: Y[5],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(6).WithAlpha(64),
					},
			    	Name: "Dark Pink",
			        XValues: X,
			        YValues: Y[6],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(7).WithAlpha(64),
					},
			    	Name: "Light Blue",
			        XValues: X,
			        YValues: Y[7],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(8).WithAlpha(64),
					},
			    	Name: "Orange",
			        XValues: X,
			        YValues: Y[8],
			    },	
			    chart.ContinuousSeries{
			    	Style: chart.Style{
						StrokeColor: chart.GetDefaultColor(9).WithAlpha(64),
					},
			    	Name: "Bright Green",
			        XValues: X,
			        YValues: Y[9],
			    },
	    },
	}
	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

f, _ := os.Create("population.png")
defer f.Close()
// buffer := bytes.NewBuffer([]byte{})
graph.Render(chart.PNG, f)
}



// Function to determin who dies at the end of the generation
func (g *Game) EndDay() {

	for k := 0; k < g.populations.numTypes; k++ {
		var deadSprites int
		for i := 0; i < g.populations.population[k].sprites.num; i++ {
			if g.populations.population[k].sprites.sprites[i].reached == false {
				deadSprites++
			}
		}
		// includes reproduction
		g.populations.population[k].sprites.newNum = (g.populations.population[k].sprites.num - deadSprites) * 2
	}
	

}


// Random number from a gaussian
func weightedRandom(max float64, numDice int) float64 {
    num := 0.0
    for i := 0; i < numDice; i++ {
        num += rand.Float64() * max/float64(numDice);
    }    
    return num
}

// Function to see if the sprite happened across a food
func (s *Sprite) IfAteFood (f *Food, scaleFactor float64, k int) bool {
	ws, hs := spriteImgs[k].Size() 
	ws = int(float64(ws)*0.2*scaleFactor)
	hs = int(float64(hs)*0.2*scaleFactor)

	if f.x >= s.x - ws/2 && f.y >= s.y - hs/2 && f.x <= s.x + ws/2 && f.y <= s.y +ws/2 {
		return true
	} else {
		return false
	}

}

// Update parameters to determine whether the food should be drawn in the next frame
func (g *Game) UpdateGame() {
	for k := 0; k < g.populations.numTypes; k++ {
		for i := 0; i < g.populations.population[k].sprites.num; i++ {
			for j := 0; j < g.foods.num; j++ {
				if g.populations.population[k].sprites.sprites[i].IfAteFood(g.foods.foods[j], g.populations.population[k].scaleFactor, k) == true && g.foods.foods[j].present == true && g.populations.population[k].sprites.sprites[i].energy < g.populations.population[k].energyCap{
					g.populations.population[k].sprites.sprites[i].energy++
					// fmt.Println(g.sprites.sprites[i].energy)
					g.foods.foods[j].present = false
				}
			}
		}
	}
}

// Update the direction of each sprite
func (s *Sprite) Update(rate, energyCap float64) {

	areaL := 700
	areaB := 370
	c := 73

	// rate := 2.0

	// If the sprite still needs to eat food
	if s.energy < energyCap {

		boundX1 := int(screenWidth/2 - areaL/2) 
		boundX2 := int(screenWidth/2 + areaL/2) 
		boundY1 := int(screenHeight/2 - areaB/2)
		boundY2 := int(screenHeight/2 + areaB/2)

		s.x += int(rate*s.vx)
		s.y += int(rate*s.vy)

		// boundary conditions
		if s.x < boundX1 {
			s.vx = 1.0
		}
		if s.x > boundX2 {
			s.vx = -1.0
		} 
		if s.y < boundY1 {
			s.vy = 1.0
		}
		if s.y > boundY2 {
			s.vy = -1.0
		}

		// Velocity is randomized using a gaussian
		wr := weightedRandom(100.0, 10)
		if wr < 25.0 || wr > 75.0 {
			vx, vy := 2*rand.Intn(2)-1, 2*rand.Intn(2)-1
			s.vx = float64(vx)
			s.vy = float64(vy)	
		}

		// When the sprite has reached appetite capacity
	} else {

		if s.x <= int(screenWidth/2 - areaL/2) - c || s.y <= int(screenHeight/2 - areaB/2) - c || s.x >= int(screenWidth/2 + areaL/2) + c || s.y >= int(screenHeight/2 + areaB/2) + c {
			s.x = s.x
			s.y = s.y
			s.reached = true
		} else {
			s.x += int(rate*s.vx)
			s.y += int(rate*s.vy)
		}

	}

}

// Sends each sprite to sprite updation
func (s *Sprites) Update(rate, energyCap float64) {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Update(rate, energyCap)
	}
}