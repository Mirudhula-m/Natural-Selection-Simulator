# Natural-Selection-Simulator

Follow this [link](https://github.com/Mirudhula-m/Natural-Selection-Simulator/tree/main) to view this project's github repository.

## Overview

We see everyone researchers talking about how a certain trait evolves over time and how certain other populations with undesirable traits gets eliminated. Visualizing such events gives us an idea about what exactly is happening.


In this example demo, we have three different traits that a population could have:
<ol>
	<li>Speed</li>
	<li>Size</li>
	<li>Appetite</li>
</ol><br>
These traits can be combined multiple ways to form the overall character of a population.

We can add multiple such populations with the desired traits, and let the simulation run to see who survives after n generations of time!


## Installation of Packages

This program extensively uses the 2D game library known as Ebiten. The following instructions can be followed in order to make Ebiten run smoothly on your laptop:

<ul>
<li><strong>Version:</strong> Ebiten prefers Go ver. 1.15 or later.</li>
<li><strong>C Compiler:</strong> Ebiten uses C compiler in addition to Go. On macOS, please download clang, if it not already present.</li>
<li>Ebiten unfortunately does not work with GOPATH. Therefore, we will have to turn on the GO111MODULE to on, before we run any program using Ebiten. This can be done on your local bash_profile file or simply on the terminal by using:

	export GO111MODULE=on
</li>
<li>Once this is setup, we will need to get the actual Ebiten v2 packages. Use the following code on your command line:
	
	go get github.com/hajimehoshi/ebiten/v2
</li>
<li>Now, you have successfully installed Ebiten on your system. To test run it, try out an example code as below:
	
	go run -tags=example github.com/hajimehoshi/ebiten/v2/examples/rotate
If this runs, then you are all set for the next step.
</li>
</ul>

This program also uses some external packages that we may have to install before running the program.

<ul>
<li>The font used in this game is obtained from a fonts package on golang. Use the following code in the command prompt to install it:
	
	go get golang.org/x/image/font/sfnt
</li>
<li> Next, we also use a package name <strong> go-chart </strong> to draw the plots. This can be installed as follows:
	
	go get -u github.com/wcharczuk/go-chart
</li>
</ul>

Once you are done with this, you can download the code for the natural selection as well, and essentially start running the program! <br>
Please note that individual systems may need some more dependencies for the codes to run. In the likely event that you need to download any more of them, you will be prompted to do so on your command prompt.


## Running the Simulator

The simulator is fairly easy to run, and fun to observe!

To start the simulator, complete the steps below:

<ul>
<li>cd into your game directory</li>
<li>Initialize go.mod in your command prompt. Below is an example code on the terminal that you can use:
	
	go mod init github.com/yourname/yourgame
</li>
<li>Use:
	
	go build
	./simulator
This will start the simulator.
</li>
</ul>

You should see a window pop up, as below:

 ![Title Page!](/screenshots/title.png)

Pressing on 1 will add more populations to your game and you can define the population charactr traits based on the combination of the three traits mentioned previously. Just press the corresponding key to select the desired trait.

![Add Population!](/screenshots/addpop.png)

Pressing on 2 while in main title page will begin the simulation.

In this simulation, each population has a different colour of sprite that corresponds to it. At the start of each generation (which lasts around 7s), all the sprites will start at boundaries of a platform. These sprites are required to begin their search for food and continue till they fill their appetite (energy capacity). Once they are full, they need to go back to the boundary of the platform before the generation ends. If they are unable to do so (either because they were unable to find food or because the generation ran out of food stock), they die. The sprites that do survive and go back, reproduce asexually on the next run of the generation.

![Game!](/screenshots/game.png)

If you let this run for a few generations, we should be able to see some stability or stagnation in the number of sprites in each population. This can observed as a plot for each game. 

In order to see the plot, press the 0 key to stop the current simulation. This will take you back to the man menu of the game, where you can actually re-start the game with a different set of parameters. There should be a 'population.png' file that is generated locally when you press this 0 key. This shows how each population varied in numbers over generations.

## Possible Future Direction

### Minor Improvements

<ol>
	<li>The different coloured sprites could be tagged with respect to their corresponding traits during that run of the simulation.</li>
	<li>The plots could be drawn more efficiently. Increase the line width of the plots and avoid showing populations that are not present in the game run as a legend in the plot. It might also be better to change the x and y axis ticklabels so that they are easier to read.</li>
</ol>

### Additional Content

<ol>
<li>It would be interesting to randomize the trait levels. For example, the speed of the organism could start at a random value, and after a few generations, we should be able to observe a constant average speed that will try to maximize the number of species observed.</li>
<li>Calculate the fitness function, and try to do some back calculations to get back the population size after n generations. We can try to see if we can get back the observed value, and get the error between observed and expected.</li>
</ol>

## Credits and Acknowledgement

This simulator was inspired by Primer's [video](https://www.youtube.com/watch?v=0ZGbIKd0XrM) on Simulating Natural Selection. Primer simulates this by only using graphics as a tool. I thought it would be interesting to bring this to life by actually coding up the simulation.






