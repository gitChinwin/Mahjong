package internal

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/17
 */

func Test_HaveConcealedKong(t *testing.T) {
	Convey("Test HaveConcealedKong", t, func() {
		obj := make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.HaveConcealedKong(), ShouldBeTrue)

		obj = make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.HaveConcealedKong(), ShouldBeFalse)

		obj = make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
			&Tile{TileBamboo, 1},
			&Tile{TileBamboo, 2},
			&Tile{TileBamboo, 3},
			&Tile{TileBamboo, 4},
			&Tile{TileBamboo, 5},
			&Tile{TileBamboo, 6},
			&Tile{TileBamboo, 7},
			&Tile{TileDragon, 1},
		)
		So(obj.HaveConcealedKong(), ShouldBeTrue)
	})

}

func Test_CanKong(t *testing.T) {
	Convey("Test CanKong", t, func() {
		obj := make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.CanKong(&Tile{TileDragon, 2}), ShouldBeFalse)
		So(obj.CanKong(&Tile{TileDragon, 1}), ShouldBeTrue)
		So(obj.CanKong(&Tile{TileDragon, 3}), ShouldBeFalse)

		obj = make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.CanKong(&Tile{TileDragon, 2}), ShouldBeFalse)
		So(obj.CanKong(&Tile{TileDragon, 3}), ShouldBeFalse)
	})

}

func Test_CanPong(t *testing.T) {
	Convey("Test CanPong", t, func() {
		obj := make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.CanPong(&Tile{TileDragon, 2}), ShouldBeFalse)
		So(obj.CanPong(&Tile{TileDragon, 1}), ShouldBeTrue)
		So(obj.CanPong(&Tile{TileDragon, 3}), ShouldBeFalse)

		obj = make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
		)
		So(obj.CanPong(&Tile{TileDragon, 2}), ShouldBeFalse)
		So(obj.CanPong(&Tile{TileDragon, 1}), ShouldBeTrue)
		So(obj.CanPong(&Tile{TileDragon, 3}), ShouldBeFalse)

		obj = make(ts, 0)
		obj = append(obj,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
			&Tile{TileDragon, 2},
		)
		So(obj.CanPong(&Tile{TileDragon, 1}), ShouldBeTrue)
		So(obj.CanPong(&Tile{TileDragon, 2}), ShouldBeTrue)
		So(obj.CanPong(&Tile{TileDragon, 3}), ShouldBeFalse)
	})

}
