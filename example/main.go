package main

import (
	"ebby"
	"ebby/control"
	"snake/test/textscen"
	"snake/test/textscen2"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	sm := control.ScenarioMap{
		"test":  textscen.Scenario(),
		"test2": textscen2.Scenario(),
	}

	ebby.New(sm).Run()
}
