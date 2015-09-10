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

func checkerr(e error) {
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

func Display() {
	for y := 0; y < MapY; y++ {
		var yy = fmt.Sprintf("%2d", y)

		fmt.Println(yy + ": " + fmt.Sprintf("%s", Mapp[y*MapX:y*MapX+MapX]))
	}
}

func Save(filename string) {
	f, err := os.Create(filename)

	checkerr(err)

	defer f.Close()

	for y := 0; y < MapY; y++ {
		_, err := f.WriteString(fmt.Sprintf("%s\n", Mapp[y*MapX:y*MapX+MapX]))

		checkerr(err)
	}

	f.Sync()
}
