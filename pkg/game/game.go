package game

import (
	"battleships-board/pkg/ships"
	"battleships-board/types"
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
			board.Moves[c] = types.CoordState{
				Ship:  ship,
				State: types.SEA,
			}
		}
	}
	board.ShipTot = ships.GetVolume(s)
	return board
}

func TakeShot(from, to *types.Player, target types.Coord) string {
	if _, ok := to.Board.Moves[target]; ok {
		move := to.Board.Moves[target]
		if move.State != types.HIT && move.State != types.SUNK {
			ship := move.Ship
			ship.Hits++
			to.Board.Moves[target] = types.CoordState{
				State: types.HIT,
				Ship:  ship,
			}
			// save against the player that shot
			from.Moves[target] = types.CoordState{
				State: move.State,
				Ship:  move.Ship,
			}
			from.Hits++
		}
		move = to.Board.Moves[target]
		if move.Ship.Hits == len(move.Ship.Coords) {
			for _, coord := range move.Ship.Coords {
				to.Board.Moves[coord] = types.CoordState{
					State: types.SUNK,
					Ship:  move.Ship,
				}
				// save against the player that shot
				from.Moves[target] = types.CoordState{
					State: types.SUNK,
					Ship:  move.Ship,
				}
			}
		}
	} else {
		to.Board.Moves[target] = types.CoordState{
			State: types.MISS,
		}
		// save against the player that shot
		from.Moves[target] = types.CoordState{
			State: types.MISS,
		}
	}

	// check for winning shot
	if from.Hits == from.Board.ShipTot {
		return from.Name
	}

	return ""
}
