// Algorithm from http://www.evilscience.co.uk/map-generator/

package cellular

import (
	"github.com/andrew-ayers/helpers/dungeon"
)

var r = dungeon.NewRand()

func open(p int) {
	for i := 0; i < dungeon.MapX*dungeon.MapY; i++ {
		if r.Intn(99) >= p {
			dungeon.Mapp[i] = byte(dungeon.TOpen)
		}
	}
}

func isClosed(xpos int, ypos int) int {
	var count = 0

	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if !(x == 0 && y == 0) {
				if dungeon.IsTile(xpos+x, ypos+y, dungeon.TRock) == true {
					count++
				}
			}
		}
	}

	return count
}

func Generate(p int, h bool, n int, iterations int) {
	open(p)

	var opened, closed = dungeon.TOpen, dungeon.TRock

	if h != true {
		opened, closed = dungeon.TRock, dungeon.TOpen
	}

	for iterations > 0 {
		var randx = r.Intn(dungeon.MapX - 1)
		var randy = r.Intn(dungeon.MapY - 1)

		if isClosed(randx, randy) > n {
			dungeon.Set(randx, randy, closed)
		} else {
			dungeon.Set(randx, randy, opened)
		}

		iterations--
	}
}
