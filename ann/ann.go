package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*****************************************************************************/

// struct for a singular neurode
type Neurode struct {
	id int

	ntype string // type of neurode ("IP" = input, "OP" = output, "IN" = intermediate)

	x int // +/- 32768 x/y/z coordinates
	y int
	z int

	pthresh int
	pdelay  int

	recepts []int // incoming neurode IDs
	forcept int   // outgoing neurode ID

	flength int // random "length" of forcept contact
}

// struct for a mass of neurodes
type Brane struct {
	id       int
	neurodes []Neurode
}

/*****************************************************************************/

var brane = Brane{} // initialize the brane

/*****************************************************************************/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	buildBrane(1, 2)

	fmt.Printf("Brane: %v\n\n", brane)
	fmt.Printf("Neurodes: %v\n\n", len(brane.neurodes))

	wireBrane(1, 1)
}

func buildBrane(bid int, level int) {
	// number of neurodes to build for this "level"
	ncount := int(math.Pow(2, float64(level)))

	// build the neurodes
	for i := 0; i < ncount; i++ {
		brane.neurodes = append(brane.neurodes, buildNeurode())
	}

	// reduce level and keep building (as needed)
	level--

	if level > 0 {
		buildBrane(bid, level)
	}

	// give an ID to the brane
	brane.id = bid

	return
}

func buildNeurode() Neurode {
	// build a randomized new neurode
	var newNeurode = Neurode{
		id: len(brane.neurodes),

		ntype: "IN",

		x: rand.Intn(512) - 256,
		y: rand.Intn(512) - 256,
		z: rand.Intn(512) - 256,

		pthresh: 0,
		pdelay:  0,

		flength: rand.Intn(1024),
	}

	return newNeurode
}

func wireBrane(min int, max int) {
	// wire the brane in a random fashion

	// for each neurode...
	for i := 0; i < len(brane.neurodes); i++ {
		this := brane.neurodes[i]

		// ...find all neurodes within reach of its forcept
		//var nearNeurodes []Neurode

		for j := 0; j < len(brane.neurodes); j++ {
			// skip self...
			if j != i {
				that := brane.neurodes[j]

				// calc 3d distance
				distance := int(math.Hypot(math.Hypot(float64(this.x)-float64(that.x), float64(this.y)-float64(that.y)), float64(this.z)-float64(that.z)))

				if distance <= this.flength {
					fmt.Println(this.id, that.id, distance)
				}
			}
		}
	}

}
