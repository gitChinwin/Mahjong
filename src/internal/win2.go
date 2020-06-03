package internal

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

func (pl *Player) Win2() bool {
	pl.SortTiles()

	// 牌数量 3n+2
	if (pl.HoldTiles.Len()-2)%3 != 0 {
		return false
	}

	return true
}
