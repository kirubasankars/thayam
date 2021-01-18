package game

import (
	"errors"
	"fmt"
	"math/rand"
	"thayam/maps"
	"thayam/utils"
	"time"
)

type Game struct {
	gameContext *GameContext

	actionTimeout time.Time

	board      Map
	ticker     *time.Ticker
	tickerDone chan bool
	phase      string
}

func (game *Game) UpdateSetting(gs GameSetting) error {
	if game.phase != "SETUP" {
		return errors.New("invalid operation")
	}
	game.initialGame(gs)
	return nil
}

func (game Game) Start() {
	game.actionTimeout = time.Now()
	game.gameContext.currentAction = "ROLL_DICE"
	game.phase = "STARTED"
	go func() {
		for {
			select {
			case <-game.tickerDone:
				return
			case t := <-game.ticker.C:
				_ = t
				game.actionLoop()
			}
		}
	}()
}

func (game Game) Stop() {
	game.tickerDone <- false
	game.ticker.Stop()
}

func (game Game) RollDice() {
	game.gameContext.diceQueue = append(game.gameContext.diceQueue, rollDice())
}

func (game Game) PlaceAPiece(playerID int, piecePosition int, diceIndex int) error {
	var diceQueue = game.gameContext.diceQueue
	var dice = diceQueue[diceIndex]
	var pieces = game.gameContext.pieces

	if piecePosition == -1 {

		placeInPort := func() error {
			n, err := game.board.GetInitialPosition(playerID)
			if err != nil {
				return err
			}
			pieces[n][playerID]++
			game.gameContext.diceQueue = utils.Remove(diceQueue, diceIndex)
			return nil
		}

		if game.gameContext.players[playerID].Initialized {
			if dice == 1 || dice == 5 {
				return placeInPort()
			}
		} else {
			if dice == 1 {
				game.gameContext.players[playerID].Initialized = true
				return placeInPort()
			}
		}

	} else {

		if pieces[piecePosition][playerID] > 0 {
			nextPosition, err := game.board.GetNextPosition(playerID, piecePosition, dice)
			if err != nil {
				return err
			}
			pieces[piecePosition][playerID]--
			pieces[nextPosition][playerID]++
			game.gameContext.diceQueue = utils.Remove(diceQueue, diceIndex)
		} else {
			return errors.New("no piece there")
		}

	}

	if len(game.gameContext.diceQueue) == 0 {
		game.EndAction()
	}

	return nil
}

func (game Game) GetValidPositions(playerID int, piecePosition int) [][2]int {
	var diceQueue = game.gameContext.diceQueue
	var output = make([][2]int, len(diceQueue))
	for idx, dice := range diceQueue {
		p, err := game.board.GetNextPosition(playerID, piecePosition, dice)
		if err != nil {
			output[idx] = [2]int{dice, p}
		}
	}
	return output
}

func (game *Game) EndAction() {
	//fmt.Println(game.gameContext.currentPlayer, game.gameContext.currentAction)

	gameContext := game.gameContext
	var endRoll = false
	if gameContext.currentAction == "ROLL_DICE" {
		queueLen := len(gameContext.diceQueue)
		if queueLen > 0 {
			lastRoll := gameContext.diceQueue[queueLen-1]
			if lastRoll == 2 || lastRoll == 3 || lastRoll == 4 {
				endRoll = true
			}
		}
	}

	if gameContext.currentAction == "ROLL_DICE" {
		if endRoll {
			gameContext.currentAction = "MOVE_PIECE"
		}
	} else {
		if gameContext.currentPlayer >= gameContext.gameSetting.PlayerCount-1 {
			gameContext.currentPlayer = 0
		} else {
			gameContext.currentPlayer++
		}
		gameContext.diceQueue = []int{}
		gameContext.currentAction = "ROLL_DICE"
	}

	game.actionTimeout = time.Now()
}

func (game *Game) countBoardPieces(playerID int) int {
	var pieces = game.gameContext.pieces
	var count = 0
	for _, v := range pieces {
		count += v[playerID]
	}
	return count
}

func (game *Game) getBoardPieces(playerID int) []int {
	var pieces = game.gameContext.pieces
	var output []int
	for idx, v := range pieces {
		if v[playerID] > 0 {
			for x := 0; x < v[playerID]; x++ {
				output = append(output, idx)
			}
		}
	}
	return output
}

func (game *Game) movePiece() {
	fmt.Println("Player", game.gameContext.currentPlayer, game.gameContext.diceQueue)
	currentPlayerID := game.gameContext.currentPlayer
	currentPlayer := game.gameContext.players[game.gameContext.currentPlayer]

	if !currentPlayer.Initialized {
		for idx, v := range game.gameContext.diceQueue {
			if v == 1 {
				game.PlaceAPiece(currentPlayerID, -1, idx)
				currentPlayer.Initialized = true
				break
			}
		}
	}

	if currentPlayer.Initialized {
		c := game.countBoardPieces(currentPlayerID)
		d := game.gameContext.gameSetting.PieceCount - c

		if d > 0 {
			for p := 0; p < d; p ++ {
				var onesAndFixes []int
				for idx, v := range game.gameContext.diceQueue {
					if v == 1 || v == 5 {
						onesAndFixes = append(onesAndFixes, idx)
					}
				}
				for _, v := range onesAndFixes {
					game.PlaceAPiece(currentPlayerID, -1, v)
				}

				if len(onesAndFixes) == 0 {
					break
				}
			}
		}

		c = game.countBoardPieces(currentPlayerID)

		for len(game.gameContext.diceQueue) > 0 && c > 0 {
			inBoardPieces := game.getBoardPieces(currentPlayerID)
			p := inBoardPieces[rand.Intn(len(inBoardPieces))]
			game.PlaceAPiece(currentPlayerID, p, 0)
		}
	}

	var pieces []int
	for _ , p := range game.gameContext.pieces {
		pieces = append(pieces, p[currentPlayerID])
	}

	fmt.Println(pieces)

}

func (game *Game) actionLoop() {
	var timeout int
	if game.gameContext.currentAction == "ROLL_DICE" {
		timeout = 1
	}
	if game.gameContext.currentAction == "MOVE_PIECE" {
		timeout = 1
	}
	durationSeconds := time.Now().Sub(game.actionTimeout).Seconds()
	if int(durationSeconds) > timeout {
		if game.gameContext.currentAction == "ROLL_DICE" {
			game.RollDice()
		}
		if game.gameContext.currentAction == "MOVE_PIECE" {
			game.movePiece()
		}
		game.EndAction()
	}
}

func (game *Game) initialGame(gs GameSetting) {
	game.gameContext.gameSetting = gs
	if gs.MapType == 0 {
		game.board = maps.NewMap7()
	}
	if gs.MapType == 1 {
		game.board = maps.NewMap9()
	}
	game.gameContext.players = make([]*Player, gs.PlayerCount)
	for idx, _ := range game.gameContext.players {
		p := new(Player)
		p.Initialized = true
		p.ID = idx
		game.gameContext.players[idx] = p
	}
	game.gameContext.pieces = make([][]int, game.board.GetTileCount())

	for idx, _ := range game.gameContext.pieces {
		game.gameContext.pieces[idx] = make([]int, gs.PlayerCount)
	}
}

func NewGame(gs GameSetting) *Game {
	game := new(Game)
	game.gameContext = new(GameContext)
	game.gameContext.gameSetting = gs
	game.phase = "SETUP"
	game.initialGame(gs)
	game.ticker = time.NewTicker(500 * time.Millisecond)
	game.tickerDone = make(chan bool)
	return game
}
