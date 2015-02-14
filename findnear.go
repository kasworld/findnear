// 2d tile space find functions
package findnear

import (
	"math"
	"sort"

	"kasworld/goguelike/lib/direction"
)

type XYLen struct {
	X, Y int
	L    float64
}
type XYLenList []XYLen

func (s XYLenList) Len() int {
	return len(s)
}
func (s XYLenList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s XYLenList) Less(i, j int) bool {
	return s[i].L < s[j].L
}

func NewXYLenList(xmax, ymax int) XYLenList {
	rtn := make(XYLenList, 0)
	for x := -xmax / 2; x < xmax/2; x++ {
		for y := -ymax / 2; y < ymax/2; y++ {
			rtn = append(rtn, XYLen{
				x, y,
				math.Sqrt(float64(x*x + y*y)),
			})
		}
	}
	sort.Sort(rtn)
	return rtn
}

// search from center
type DoFn func(int, int) bool

func (pll XYLenList) FindAll(x, y int, fn DoFn) bool {
	return pll.Find(x, y, 0, len(pll), fn)
}

func (pll XYLenList) Find(x, y int, start, end int, fn DoFn) bool {
	if start > end || start < 0 || end > len(pll) {
		return false
	}
	for _, v := range pll[start:end] {
		if fn(x+v.X, y+v.Y) {
			return true
		}
	}
	return false
}

func Call8WayTile(ox, oy int, fn DoFn) []uint8 {
	TileDirs := []uint8{}
	for i := uint8(1); i <= 8; i++ {
		x, y := ox+direction.Dir2Info[i].Vt[0], oy+direction.Dir2Info[i].Vt[1]
		if fn(x, y) {
			TileDirs = append(TileDirs, i)
		}
	}
	return TileDirs
}
func Call4WayTile(ox, oy int, fn DoFn) []uint8 {
	TileDirs := []uint8{}
	for i := uint8(1); i <= 8; i += 2 {
		x, y := ox+direction.Dir2Info[i].Vt[0], oy+direction.Dir2Info[i].Vt[1]
		if fn(x, y) {
			TileDirs = append(TileDirs, i)
		}
	}
	return TileDirs
}

func GetSignAbs(n int) (int, int) {
	if n > 0 {
		return 1, n
	} else if n < 0 {
		return -1, -n
	} else {
		return 0, 0
	}
}

func CallPath(srcx, srcy, dstx, dsty int, fn DoFn) {
	diffx := dstx - srcx
	diffy := dsty - srcy
	xs, xa := GetSignAbs(dstx - srcx)
	ys, ya := GetSignAbs(dsty - srcy)
	if xa > ya { // horizontal
		for x := srcx + xs; x != dstx; x += xs {
			y := srcy + (x-srcx)*diffy/diffx
			if fn(x, y) {
				return
			}
		}
	} else {
		for y := srcy + ys; y != dsty; y += ys {
			x := srcx + (y-srcy)*diffx/diffy
			if fn(x, y) {
				return
			}
		}
	}
}