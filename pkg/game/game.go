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
	b := types.Board{
		Dim:   boardSize,
		Moves: make(types.Moves),
	}
	s := ships.GenerateShips(boardSize)
	b = addShipsToBoard(b, s)
	return b
}

func addShipsToBoard(b types.Board, s types.Ships) types.Board {
	for _, ship := range s {
		for _, c := range ship.Coords {
			b.Moves[c.String()] = types.CoordState{
				Ship:  &ship,
				State: types.SEA,
			}
		}
	}
	b.ShipTot = ships.GetVolume(s)
	return b
}

func TakeShot(from, to *types.Player, target types.Coord) string {
	if move, ok := to.Board.Moves[target.String()]; ok {
		if move.State != types.HIT && move.State != types.SUNK {

			s := move.Ship

			switch {
			case s.Hits == len(s.Coords)-1:
				// update all coords as SUNK
				for _, c := range s.Coords {
					to.Board.Moves[c.String()] = types.CoordState{
						State: types.SUNK,
						Ship: &types.Ship{
							Coords: s.Coords,
							Hits:   s.Hits + 1,
						},
					}
					// save against the player that shot
					from.Moves[c.String()] = types.CoordState{
						State: types.SUNK,
					}
				}
			default:
				// update all coords with new Hits, preserve State
				for _, c := range s.Coords {
					existingMove := to.Board.Moves[c.String()]
					state := existingMove.State
					if c == target {
						state = types.HIT
					}
					to.Board.Moves[c.String()] = types.CoordState{
						State: state,
						Ship: &types.Ship{
							Coords: existingMove.Ship.Coords,
							Hits:   existingMove.Ship.Hits + 1,
						},
					}
				}
				// save against the player that shot
				from.Moves[target.String()] = types.CoordState{
					State: types.HIT,
				}
			}

			// count the hit
			from.Hits++
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
