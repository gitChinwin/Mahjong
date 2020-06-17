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

		//|四筒|五筒|六筒|五条|  |六条|七条|三万|三万|  |四万|五万|六万|六万|  |七万|五万|
		Convey("test 3", func(t C) {
			p.HoldTiles = append(p.HoldTiles,
				&Tile{TileDot, 4},
				&Tile{TileDot, 5},
				&Tile{TileDot, 6},

				&Tile{TileBamboo, 5},
				&Tile{TileBamboo, 6},
				&Tile{TileBamboo, 7},

				&Tile{TileCharacter, 3},
				&Tile{TileCharacter, 3},
				&Tile{TileCharacter, 4},
				&Tile{TileCharacter, 5},
				&Tile{TileCharacter, 5},
				&Tile{TileCharacter, 6},
				&Tile{TileCharacter, 6},

				&Tile{TileCharacter, 7},
			)

			So(p.Win(), ShouldBeTrue)
		})

		Convey("test 4", func(t C) {
			p.HoldTiles = append(p.HoldTiles,
				&Tile{TileDot, 4},
				&Tile{TileDot, 5},
				&Tile{TileDot, 6},

				&Tile{TileBamboo, 5},
				&Tile{TileBamboo, 6},
				&Tile{TileBamboo, 7},

				&Tile{TileCharacter, 3},
				&Tile{TileCharacter, 3},
				&Tile{TileCharacter, 4},
				&Tile{TileCharacter, 4},

				&Tile{TileCharacter, 5},
				&Tile{TileCharacter, 5},
				&Tile{TileCharacter, 6},
				&Tile{TileCharacter, 6},
			)

			So(p.Win(), ShouldBeTrue)
		})
	})
}
