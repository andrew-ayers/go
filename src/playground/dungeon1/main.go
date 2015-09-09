// dungeon1 project main.go
// Algorithm from http://www.evilscience.co.uk/map-generator/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const mapx = 150
const mapy = 100

var mapp [mapy][mapx]byte

//var p, h, n, i = 45, true, 4, 50000 // island map
//var p, h, n, i = 45, false, 4, 50000 // labyrinth map
var p, h, n, i = 65, true, 4, 80000 // dense cave map
//var p, h, n, i = 45, true, 4, 85000 // sparse island map
//var p, h, n, i = 45, true, 4, 85000 // sparse island map
//var p, h, n, i = 65, false, 5, 50000
//var p, h, n, i = 48, true, 4, 50000

func newrand() *rand.Rand {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func clear() {
	for y := 0; y < mapy; y++ {
		for x := 0; x < mapx; x++ {
			mapp[y][x] = byte('.')
		}
	}
}

func open() {
	r := newrand()

	for y := 0; y < mapy; y++ {
		for x := 0; x < mapx; x++ {
			if r.Intn(99) < p {
				mapp[y][x] = byte('#')
			}
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
	if xpos >= 0 && xpos < mapx && ypos >= 0 && ypos < mapy {
		return (mapp[ypos][xpos] == '#')
	} else {
		return false
	}
}

func generate(iterations int) {
	r := newrand()

	clear()
	open()

	for iterations > 0 {
		var randx = r.Intn(mapx - 1)
		var randy = r.Intn(mapy - 1)

		var isclosed = closed(randx, randy)

		if h == true {
			if isclosed > n {
				// close cell
				mapp[randy][randx] = '#'
			} else {
				// open cell
				mapp[randy][randx] = '.'
			}
		} else {
			if isclosed > n {
				// open cell
				mapp[randy][randx] = '.'
			} else {
				// close cell
				mapp[randy][randx] = '#'
			}
		}

		iterations--
	}
}

func display() {
	for y := 0; y < mapy; y++ {
		var yy = fmt.Sprintf("%2d", y)

		fmt.Println(yy + ": " + fmt.Sprintf("%s", mapp[y]))
	}
}

func save() {
	f, err := os.Create("dungeon.dat")

	checkerr(err)

	defer f.Close()

	for y := 0; y < mapy; y++ {
		_, err := f.WriteString(fmt.Sprintf("%s\n", mapp[y]))

		checkerr(err)
	}

	f.Sync()
}

func main() {
	generate(i)
	display()
	save()
}
