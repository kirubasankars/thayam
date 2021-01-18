package maps

import "errors"

type Map9 struct {
	ports        []int
	player0Route [][2]int
	player1Route [][2]int
	player2Route [][2]int
	player3Route [][2]int
}

func (map9 *Map9) Build() {
	map9.ports = []int{3, 11, 19, 27, 32, 38, 44, 50, 56, 60, 64, 68, 80}

	map9.player0Route = [][2]int{
		{31, 0},
		{2, 32},
	}
	map9.player1Route = [][2]int{
		{31, 0},
		{10, 50},
		{55, 32},
		{49, 68},
		{71, 56},
		{67, 78},
		{79, 72},
		{77, 80},
	}
	map9.player2Route = [][2]int{
		{31, 0},
		{18, 44},
		{55, 32},
		{43, 64},
		{71, 56},
		{63, 76},
		{79, 72},
		{75, 80},
	}
	map9.player3Route = [][2]int{
		{31, 0},
		{26, 38},
		{55, 32},
		{37, 60},
		{71, 56},
		{59, 74},
		{79, 72},
		{73, 80},
	}
}

func (map9 Map9) GetTileCount() int {
	return 80
}

func (map9 Map9) GetPorts() []int {
	return map9.ports
}

func (map9 Map9) GetInitialPosition(side int) (int, error) {
	if side < 0 || side > 3 {
		return -1, errors.New("invalid operation")
	}
	return map9.ports[side], nil
}

func (map9 Map9) GetNextPosition(side int, currentPlace int, dice int) (int, error) {
	if side < 0 || side > 3 {
		return -1, errors.New("invalid operation")
	}

	var paths [][2]int
	if side == 0 {
		paths = map9.player0Route
	}
	if side == 1 {
		paths = map9.player1Route
	}
	if side == 2 {
		paths = map9.player2Route
	}
	if side == 3 {
		paths = map9.player3Route
	}

	var nextPlace = currentPlace

	jumpNextPath := func(n int) int {
		for _, v := range paths {
			if v[0] == n {
				return v[1]
			}
		}
		return -1
	}

	for i := 0; i < dice; i++ {
		if n := jumpNextPath(nextPlace); n != -1 {
			nextPlace = n
		} else {
			nextPlace = nextPlace + 1
		}
		if nextPlace > 48 {
			return -1, errors.New("invalid operation")
		}
	}
	return nextPlace, nil
}

func NewMap9() *Map9 {
	board := new(Map9)
	return board
}