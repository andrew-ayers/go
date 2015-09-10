package dungeon

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Common dungeon map tile characters
var TVoid = ' ' // void
var TOpen = '.' // open ground
var TRock = 'X' // rock
var TSand = ':' // sand
var TGrss = '+' // grass
var TTree = 'Y' // tree
var TWatr = '~' // water
var TWhrl = '@' // whirlpool
var TLava = '^' // lava
var TGate = '#' // gate
var TDoor = 'D' // door
var TStup = '>' // stairs up
var TStdn = '<' // stairs down
var TBrdg = '=' // bridge

var Mapp, MapX, MapY = []byte{}, 0, 0

func NewRand() *rand.Rand {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func Init(width int, height int) {
	MapX = width
	MapY = height

	Mapp = make([]byte, width*height)
}

func Clear() {
	for i := 0; i < MapX*MapY; i++ {
		Mapp[i] = byte(TRock)
	}
}

func Border(tile rune) {
	Box(0, 0, MapX, MapY, tile)
}

func HLine(xpos int, ypos int, width int, tile rune) {
	for i := xpos; i < xpos+width; i++ {
		Set(i, ypos, tile)
	}
}

func VLine(xpos int, ypos int, height int, tile rune) {
	for i := ypos; i < ypos+height; i++ {
		Set(xpos, i, tile)
	}
}

func Box(xpos int, ypos int, width int, height int, tile rune) {
	HLine(xpos, ypos, width, tile)
	HLine(xpos, ypos+height-1, width, tile)

	VLine(xpos, ypos, height, tile)
	VLine((xpos + width - 1), ypos, height, tile)
}

func FillBox(xpos int, ypos int, width int, height int, tile rune) {
	for i := ypos; i < ypos+height; i++ {
		HLine(xpos, i, width, tile)
	}
}

func Set(xpos int, ypos int, tile rune) {
	if xpos >= 0 && xpos < MapX && ypos >= 0 && ypos < MapY {
		Mapp[ypos*MapX+xpos] = byte(tile)
	}
}

func Get(xpos int, ypos int) rune {
	if xpos >= 0 && xpos < MapX && ypos >= 0 && ypos < MapY {
		return rune(Mapp[ypos*MapX+xpos])
	}

	return rune(0)
}

func IsTile(xpos int, ypos int, tile rune) bool {
	if xpos >= 0 && xpos < MapX && ypos >= 0 && ypos < MapY {
		return (Get(xpos, ypos) == tile)
	}

	return false
}

func Display() {
	for y := 0; y < MapY; y++ {
		var yy = fmt.Sprintf("%2d", y)

		fmt.Println(yy + ": " + fmt.Sprintf("%s", Mapp[y*MapX:y*MapX+MapX]))
	}
}

func Save(filename string) {
	f, err := os.Create(filename)

	panicIfError(err)

	defer f.Close()

	for y := 0; y < MapY; y++ {
		_, err := f.WriteString(fmt.Sprintf("%s\n", Mapp[y*MapX:y*MapX+MapX]))

		panicIfError(err)
	}

	f.Sync()
}
