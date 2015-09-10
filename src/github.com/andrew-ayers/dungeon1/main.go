// dungeon1 project main.go

package main

import (
	"github.com/andrew-ayers/helpers/dungeon"
	cdungeon "github.com/andrew-ayers/helpers/dungeon/cellular"
)

func main() {
	dungeon.Init(150, 100)
	dungeon.Clear()

	//cdungeon.Generate(45, true, 4, 50000) // island map
	//cdungeon.Generate(45, false, 4, 50000) // labyrinth map
	//cdungeon.Generate(65, true, 4, 80000) // dense cave map
	cdungeon.Generate(65, true, 4, 80000) // dense cave map
	//cdungeon.Generate(45, true, 4, 85000) // sparse island map

	dungeon.Display()

	dungeon.Save("dungeon.dat")
}
