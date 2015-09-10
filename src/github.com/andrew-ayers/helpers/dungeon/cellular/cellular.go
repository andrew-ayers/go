// Algorithm from http://www.evilscience.co.uk/map-generator/

package cellular

import (
	"github.com/andrew-ayers/helpers/dungeon"
)

func open(p int) {
	r := dungeon.NewRand()

	for i := 0; i < dungeon.MapX*dungeon.MapY; i++ {
		if r.Intn(99) >= p {
			dungeon.Mapp[i] = byte(dungeon.TOpen)
		}
	}
}

func closed(xpos int, ypos int) int {
	var count = 0

	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if !(x == 0 && y == 0) {
				if check(xpos+x, ypos+y) == true {
					count++
				}
			}
		}
	}

	return count
}

func check(xpos int, ypos int) bool {
	if xpos >= 0 && xpos < dungeon.MapX && ypos >= 0 && ypos < dungeon.MapY {
		return (dungeon.Mapp[ypos*dungeon.MapX+xpos] == byte(dungeon.TRock))
	} else {
		return false
	}
}

func Generate(p int, h bool, n int, iterations int) {
	r := dungeon.NewRand()

	open(p)

	for iterations > 0 {
		var randx = r.Intn(dungeon.MapX - 1)
		var randy = r.Intn(dungeon.MapY - 1)

		var isclosed = closed(randx, randy)

		if h == true {
			if isclosed > n {
				// close cell
				dungeon.Mapp[randy*dungeon.MapX+randx] = byte(dungeon.TRock)
			} else {
				// open cell
				dungeon.Mapp[randy*dungeon.MapX+randx] = byte(dungeon.TOpen)
			}
		} else {
			if isclosed > n {
				// open cell
				dungeon.Mapp[randy*dungeon.MapX+randx] = byte(dungeon.TOpen)
			} else {
				// close cell
				dungeon.Mapp[randy*dungeon.MapX+randx] = byte(dungeon.TRock)
			}
		}

		iterations--
	}
}
