package internal

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/4
 */

func TestPlayer_Win(t *testing.T) {
	Convey("test win tile", t, func() {
		p := Player{}
		p.HoldTiles = make([]*Tile, 0)

		Convey("test 1", func(t C) {
			p.HoldTiles = append(p.HoldTiles,
				&Tile{TileDot, 1},
				&Tile{TileDot, 2},
				&Tile{TileDot, 3},

				&Tile{TileBamboo, 4},
				&Tile{TileBamboo, 5},
				&Tile{TileBamboo, 6},

				&Tile{TileBamboo, 7},
				&Tile{TileBamboo, 8},
				&Tile{TileBamboo, 9},

				&Tile{TileCharacter, 2},
				&Tile{TileCharacter, 3},
				&Tile{TileCharacter, 4},

				&Tile{TileDragon, 1},
				&Tile{TileDragon, 1},
			)

			So(p.Win(), ShouldBeTrue)
		})

		Convey("test 2", func(t C) {
			p.HoldTiles = append(p.HoldTiles,
				&Tile{TileDot, 1},
				&Tile{TileDot, 1},
				&Tile{TileDot, 1},

				&Tile{TileBamboo, 4},
				&Tile{TileBamboo, 4},
				&Tile{TileBamboo, 4},

				&Tile{TileBamboo, 6},
				&Tile{TileBamboo, 6},
				&Tile{TileBamboo, 6},

				&Tile{TileBamboo, 7},
				&Tile{TileBamboo, 7},
				&Tile{TileBamboo, 7},

				&Tile{TileDot, 2},
				&Tile{TileDot, 2},
			)

			So(p.Win(), ShouldBeTrue)
		})
	})
}

//func BenchmarkPlayer_Win(b *testing.B) {
//
//}
