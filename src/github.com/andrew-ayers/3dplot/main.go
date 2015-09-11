// 3dplot - see http://www.atariarchives.org/basicgames/showpage.php?page=167
package main

import (
	"fmt"
	"math"
	"strings"
)

var mapp = []byte{}

func fx(z float32) float32 {
	return 30 * float32(math.Exp(float64(-z*z/100)))
	//return float32(math.Sqrt(float64(900.01-z*z)))*0.9 - 2
	//return 30 * float32(math.Cos(float64(z/16))) * 2
	//return (30 * float32(math.Exp(math.Cos(float64(z/16))))) - 30
	//return 30 * float32(math.Sin(float64(z/10)))
}

func replaceAtIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func orig() {
	const offset float32 = 0.9

	fmt.Println("           3D PLOT BY MARK BRAMHALL (DEC)")
	fmt.Println("    CREATIVE COMPUTING -- MORRISTOWN, NEW JERSEY")
	fmt.Println("ADAPTED FROM BASIC COMPUTER GAMES BY DAVID AHL (1978)")
	fmt.Println("      GOLANG VERSION BY ANDREW L. AYERS (2015)")

	fmt.Println()

	var x int = 0
	var y int = 0
	var z int = 0
	var Lim int = 0
	var y1 int = 0

	for x = -30; x < 30; x += 1 {
		Lim = 0
		y1 = 5 * int(math.Sqrt(float64(900-x*x))/5)
		line := strings.Repeat(" ", 512)
		for y = y1; y > -y1; y -= 5 {
			f := float32(math.Sqrt(float64(x*x + y*y)))

			z = 25 + int(fx(f)) - int(offset*float32(y))

			if z <= Lim {
				continue
			}

			Lim = z

			line = replaceAtIndex(line, '.', z)
		}

		fmt.Println(line)
	}
}

func newplot() {
	spaces := strings.Repeat(" ", 30)

	fmt.Println(spaces + "           3D PLOT BY MARK BRAMHALL (DEC)")
	fmt.Println(spaces + "    CREATIVE COMPUTING -- MORRISTOWN, NEW JERSEY")
	fmt.Println(spaces + "ADAPTED FROM BASIC COMPUTER GAMES BY DAVID AHL (1978)")
	fmt.Println(spaces + "      GOLANG VERSION BY ANDREW L. AYERS (2015)")

	fmt.Println()

	mapp = make([]byte, 100*100)

	for i := 0; i < 100*100; i++ {
		mapp[i] = byte(' ')
	}

	const offset float32 = 0.9

	var x int = 0
	var y int = 0
	var z int = 0
	var Lim int = 0
	var y1 int = 0

	for x = -24; x < 25; x += 1 {
		Lim = 0
		y1 = 5 * int(math.Sqrt(float64(900-x*x))/5)
		line := strings.Repeat(" ", 100)
		for y = y1; y > -y1; y -= 3 {
			f := float32(math.Sqrt(float64(x*x + y*y)))

			z = 25 + int(fx(f)) - int(offset*float32(y))

			if z <= Lim {
				continue
			}

			Lim = z

			line = replaceAtIndex(line, '.', z)
		}

		var bline [100]byte
		copy(bline[:], line)
		for i := 0; i < 100; i++ {
			mapp[(99-i)*100+(50-x*2)] = bline[i]
		}
	}

	for y = 0; y < 100; y++ {
		//var yy = fmt.Sprintf("%2d", y)

		//fmt.Println(yy + ": " + fmt.Sprintf("%s", mapp[y*100:y*100+100]))
		fmt.Println(fmt.Sprintf("%s", mapp[y*100:y*100+100]))
	}
}

func main() {
	orig()

	fmt.Println()
	fmt.Println()

	newplot()
}
