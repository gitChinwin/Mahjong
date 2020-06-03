package internal

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

// win 回溯算法
func (pl *Player) WinWin() bool {
	pl.SortTiles()

	// 牌数量 3n+2
	if (pl.HoldTiles.Len()-2)%3 != 0 {
		return false
	}

	var tmp [5][]*Tile

	for _, each := range pl.HoldTiles {
		tmp[each.Type-1] = append(tmp[each.Type-1], each)
	}

	for i := 0; i < 5; i++ {
		if len(tmp[i]) == 0 {
			continue
		}

		if len(tmp[i]) == 1 {
			return false
		}

		if len(tmp[i]) == 2 {
			if !isPair(tmp[i][0], tmp[i][1]) {
				return false
			}
		}

		for j := 0; j < len(tmp[i]); j++ {
			//if tmp[i][j]
		}
	}
	return true
}

// 将牌(雀头、对子)
func isPair(a, b *Tile) bool {
	return a.Type == b.Type && a.Rank == b.Rank
}

// 顺子	Sequence
func isSequence(a, b, c *Tile) bool {
	if !(a.Type == b.Type && a.Type == c.Type) {
		return false
	}
	if !(a.Rank == b.Rank+1 && b.Rank == c.Rank+1) {
		return false
	}
	return true
}

// 刻子	Triplet
func isTriplet(a, b, c *Tile) bool {
	return (a.Type == b.Type && a.Type == c.Type) && (a.Rank == b.Rank && a.Rank == c.Rank)
}

func removePair() {

}

func removeSequence() {

}

func removeTriplet() {

}
