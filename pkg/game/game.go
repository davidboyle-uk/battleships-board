package game

import (
	"github.com/dbx123/battleships-board/pkg/ships"
	"github.com/dbx123/battleships-board/types"
)

func Initialise(boardSize int) types.Game {
	return types.Game{
		Players: []*types.Player{
			{
				Name:  "Player 1",
				Board: generateBoard(boardSize),
				Moves: make(types.Moves),
			},
			{
				Name:  "Player 2",
				Board: generateBoard(boardSize),
				Moves: make(types.Moves),
			},
		},
	}
}

func generateBoard(boardSize int) types.Board {
	board := types.Board{
		Dim:   boardSize,
		Moves: make(types.Moves),
	}
	s := ships.GenerateShips(boardSize)
	for _, ship := range s {
		for _, c := range ship.Coords {
			board.Moves[c.String()] = types.CoordState{
				Ship:  &ship,
				State: types.SEA,
			}
		}
	}
	board.ShipTot = ships.GetVolume(s)
	return board
}

func TakeShot(from, to *types.Player, target types.Coord) string {
	if move, ok := to.Board.Moves[target.String()]; ok {
		if move.State != types.HIT && move.State != types.SUNK {
			move.Ship.Hits++
			to.Board.Moves[target.String()] = types.CoordState{
				State: types.HIT,
				Ship:  move.Ship,
			}
			// save against the player that shot
			from.Moves[target.String()] = types.CoordState{
				State: types.HIT,
			}
			from.Hits++
		}
		if move.Ship.Hits == len(move.Ship.Coords) {
			for _, coord := range move.Ship.Coords {
				to.Board.Moves[coord.String()] = types.CoordState{
					State: types.SUNK,
					Ship:  move.Ship,
				}
				// save against the player that shot
				from.Moves[target.String()] = types.CoordState{
					State: types.SUNK,
				}
			}
		}
	} else {
		// save against the player that shot
		from.Moves[target.String()] = types.CoordState{
			State: types.MISS,
		}
	}

	// check for winning shot
	if from.Hits == from.Board.ShipTot {
		return from.Name
	}

	return ""
}
