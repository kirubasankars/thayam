package game

type GameLobby struct {
	games map[string]*Game
}

func (lobby *GameLobby) CreateGame(setting GameSetting) string {
	game := NewGame(setting)
	lobby.games["game"] = game
	return "game"
}

func (lobby *GameLobby) GetGame(gameID string) *Game {
	return lobby.games[gameID]
}

func NewGameLobby() GameLobby {
	for i := 0; i < 25; i ++ {
		rollDice()
	}
	ge := new (GameLobby)
	ge.games = make(map[string]*Game)
	return *ge
}