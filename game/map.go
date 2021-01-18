package game

type Map interface {
	GetPorts() []int
	GetTileCount() int
	GetInitialPosition(side int) (int, error)
	GetNextPosition(side int, currentPlace int, dice int) (int, error)
}