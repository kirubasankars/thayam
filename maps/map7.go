package maps

import "errors"

type Map7 struct {
	ports        []int
	player0Route [][2]int
	player1Route [][2]int
	player2Route [][2]int
	player3Route [][2]int
}

func (map7 *Map7) Build() {
	map7.ports = []int{2, 8, 14, 20, 24, 28, 32, 36, 48}
	map7.player0Route = [][2]int{
		{23, 0},
		{1, 24},
	}
	map7.player1Route = [][2]int{
		{23, 0},
		{7, 36},
		{39, 24},
		{35, 46},
		{47, 40},
		{45, 48},
	}
	map7.player2Route = [][2]int{
		{23, 0},
		{13, 32},
		{39, 24},
		{31, 44},
		{47, 40},
		{43, 48},
	}
	map7.player3Route = [][2]int{
		{23, 0},
		{19, 28},
		{39, 24},
		{27, 42},
		{47, 40},
		{41, 48},
	}
}

func (map7 Map7) GetTileCount() int {
	return 49
}

func (map7 Map7) GetPorts() []int {
	return map7.ports
}

func (map7 Map7) GetInitialPosition(playerID int) (int, error) {
	if playerID < 0 || playerID > 3 {
		return -1, errors.New("invalid operation")
	}
	return map7.ports[playerID], nil
}

func (map7 Map7) GetNextPosition(playerID int, currentPlace int, dice int) (int, error) {
	if playerID < 0 || playerID > 3 {
		return -1, errors.New("invalid operation")
	}

	var paths [][2]int
	if playerID == 0 {
		paths = map7.player0Route
	}
	if playerID == 1 {
		paths = map7.player1Route
	}
	if playerID == 2 {
		paths = map7.player2Route
	}
	if playerID == 3 {
		paths = map7.player3Route
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

func NewMap7() *Map7 {
	board := new(Map7)
	board.Build()
	return board
}