package game

type GameSetting struct {
	PlayerCount int
	MapType     int
	PieceCount  int
}

type GameContext struct {
	gameSetting GameSetting
	pieces    [][]int
	players   []*Player
	currentAction string

	currentPlayer int
	diceQueue []int
}
