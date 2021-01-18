package main

import (
	"thayam/game"
	"time"
)

func main()  {

	lobby := game.NewGameLobby()

	gs := game.GameSetting{MapType: 0, PieceCount: 8, PlayerCount: 2}
	gameID := lobby.CreateGame(gs)
	g := lobby.GetGame(gameID)
	g.Start()
	//g.Stop()
	time.Sleep(60 * time.Minute)
}
