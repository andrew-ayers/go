package roomgen

import (
	//"fmt"
	"github.com/andrew-ayers/helpers/dungeon"
)

var r = dungeon.NewRand()

func roomGen() {
	dungeon.FillBox(50, 50, 60, 40, dungeon.TGrss)
	dungeon.Box(50, 50, 60, 40, dungeon.TLava)
}

func Generate() {
	roomGen()
}
