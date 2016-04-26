// dungeon1 project main.go

package main

import (
	"github.com/andrew-ayers/helpers/dungeon"
	cdungeon "github.com/andrew-ayers/helpers/dungeon/cellular"
	rwdungeon "github.com/andrew-ayers/helpers/dungeon/randwalk"
	rgdungeon "github.com/andrew-ayers/helpers/dungeon/roomgen"
)

func main() {
	dungeon.Init(150, 100)

	dungeon.Clear()

	//demoCellular()

	//demoRandWalk()

	demoRoomGen()

	dungeon.Border(dungeon.TRock)

	//dungeon.FillBox(50, 50, 60, 40, dungeon.TGrss)
	//dungeon.Box(50, 50, 60, 40, dungeon.TGrss)

	dungeon.Display()

	//dungeon.Save("dungeon.dat")
}

func demoCellular() {
	//cdungeon.Generate(45, true, 4, 50000) // island
	//cdungeon.Generate(45, false, 4, 50000) // labyrinth
	cdungeon.Generate(65, true, 4, 80000) // dense cave
	//cdungeon.Generate(45, true, 4, 85000) // sparse island
}

func demoRandWalk() {
	//rwdungeon.Generate(200, 55, 75, true, true) // short and long corridors
	//rwdungeon.Generate(350, 30, 50, true, true) // short corridors
	//rwdungeon.Generate(350, 80, 100, true, true) // long corridors
	//rwdungeon.Generate(200, 55, 75, false, false) // eroded caves
	rwdungeon.Generate(350, 80, 100, false, false) // chambers and halls
}

func demoRoomGen() {
	rgdungeon.Generate(15)
}
