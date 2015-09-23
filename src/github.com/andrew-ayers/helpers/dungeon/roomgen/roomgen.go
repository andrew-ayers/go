package roomgen

import (
	//"fmt"
	"github.com/andrew-ayers/helpers/dungeon"
)

var r = dungeon.NewRand()

func roomGen() {
	dungeon.FillBox(50, 50, 60, 40, dungeon.TOpen)
	dungeon.Box(60, 60, 41, 21, dungeon.TRock, 1)
	//dungeon.Flood(65, 65, dungeon.TLava, dungeon.TRock, true)
	dungeon.Flood(60, 60, dungeon.TLava, dungeon.TOpen, true)
}

func genRoom(xpos int, ypos int, size int) (int, int, int) {
	dungeon.FillBox(xpos, ypos, size, size, dungeon.TOpen)

	var i = 0

	var numDais = r.Intn(4)

	for i = 0; i < numDais; i++ {
		genFiller(xpos, ypos, size, dungeon.TDais)
	}

	var numPool = r.Intn(2)

	if numDais < 2 {
		for i = 0; i < numPool; i++ {
			var x = 0
			var y = 0
			var s = 0

			x, y, s = genFiller(xpos, ypos, size/2, dungeon.TWatr)

			if r.Intn(99) > 50 {
				dungeon.Box(x, y, s, s, dungeon.TDais, 1)
			}
		}
	}

	return xpos, ypos, size
}

func genFiller(xpos int, ypos int, size int, fill rune) (int, int, int) {
	var x = xpos + r.Intn(size)
	var y = ypos + r.Intn(size)

	if dungeon.Get(x, y) != dungeon.TDais {
		var s = 5 + r.Intn(size/2)

		x, y = dungeon.Constrain(x, y, s, s, xpos, ypos, xpos+size, ypos+size)

		dungeon.FillBox(x, y, s, s, fill)

		return x, y, s
	}

	return -1, -1, -1
}

/*
func genPool() {

}

func genPeristyle() {
}

func genColumn() {

}

func genVoid() {

}
*/

func Generate(maxsize int) (int, int, int) {
	//roomGen()
	//return 0, 0, 0
	var randx = r.Intn(dungeon.MapX - 1)
	var randy = r.Intn(dungeon.MapY - 1)

	var size = 10 + r.Intn(maxsize)

	randx, randy = dungeon.ConstrainToMap(randx, randy, size, size)

	return genRoom(randx, randy, size)
}
