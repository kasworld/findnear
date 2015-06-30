// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// 2d tile space find functions
package findnear

import (
	"math"
	"sort"

	"github.com/kasworld/direction"
	// "github.com/kasworld/go-abs"
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

func Call8WayTile(ox, oy int, fn DoFn) []direction.Dir_Type {
	TileDirs := []direction.Dir_Type{}
	for i := direction.Dir_Type(1); i <= 8; i++ {
		x, y := ox+i.Vt()[0], oy+i.Vt()[1]
		if fn(x, y) {
			TileDirs = append(TileDirs, i)
		}
	}
	return TileDirs
}
func Call4WayTile(ox, oy int, fn DoFn) []direction.Dir_Type {
	TileDirs := []direction.Dir_Type{}
	for i := direction.Dir_Type(1); i <= 8; i += 2 {
		x, y := ox+i.Vt()[0], oy+i.Vt()[1]
		if fn(x, y) {
			TileDirs = append(TileDirs, i)
		}
	}
	return TileDirs
}
