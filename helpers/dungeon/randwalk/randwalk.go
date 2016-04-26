package randwalk

import (
	//"fmt"
	"github.com/andrew-ayers/helpers/dungeon"
)

var r = dungeon.NewRand()

func randWalk(xpos int, ypos int, p int, mine int, checks bool, checkf bool) {
	var dirx = 0
	var diry = 0

	var movetest = 25

	for mine > 0 && movetest > 0 {

		for (dirx == 0 && diry == 0) || r.Intn(99) > p {
			dirx = 0
			diry = 0
			for (dirx == 0 && diry == 0) || (dirx != 0 && diry != 0) {
				diry = 0
				dirx = r.Intn(3) - 1
				if dirx == 0 {
					diry = r.Intn(3) - 1
				}
			}
		}

		var domove = false

		if dirx != 0 && diry == 0 {
			domove = dungeon.IsTile(xpos+dirx, ypos, dungeon.TRock) &&
				(dungeon.IsTile(xpos+dirx, ypos-1, dungeon.TRock) || !checks) &&
				(dungeon.IsTile(xpos+dirx, ypos+1, dungeon.TRock) || !checks) &&
				(dungeon.IsTile(xpos+(dirx*2), ypos, dungeon.TRock) || !checkf) &&
				(dungeon.IsTile(xpos+(dirx*2), ypos-1, dungeon.TRock) || !checkf) &&
				(dungeon.IsTile(xpos+(dirx*2), ypos+1, dungeon.TRock) || !checkf)
		} else {
			domove = dungeon.IsTile(xpos, ypos+diry, dungeon.TRock) &&
				(dungeon.IsTile(xpos-1, ypos+diry, dungeon.TRock) || !checks) &&
				(dungeon.IsTile(xpos+1, ypos+diry, dungeon.TRock) || !checks) &&
				(dungeon.IsTile(xpos, ypos+(diry*2), dungeon.TRock) || !checkf) &&
				(dungeon.IsTile(xpos-1, ypos+(diry*2), dungeon.TRock) || !checkf) &&
				(dungeon.IsTile(xpos+1, ypos+(diry*2), dungeon.TRock) || !checkf)
		}

		if domove == true {
			if xpos+dirx > 0 && xpos+dirx < dungeon.MapX-1 {
				xpos += dirx
			}

			if ypos+diry > 0 && ypos+diry < dungeon.MapY-1 {
				ypos += diry
			}

			dungeon.Set(xpos, ypos, dungeon.TOpen)

			mine--

			movetest = 25

			continue
		} else {
			movetest--
		}
	}
}

func Generate(iterations int, p int, mine int, checksides bool, checkfront bool) {
	for iterations > 0 {
		var randx = r.Intn(dungeon.MapX - 1)
		var randy = r.Intn(dungeon.MapY - 1)

		randWalk(randx, randy, p, mine, checksides, checkfront)

		iterations--
	}
}
