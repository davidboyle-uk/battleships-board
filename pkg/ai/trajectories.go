package ai

import (
	"github.com/dbx123/battleships-board/types"
)

func traceTrajectories(b types.Board, pointsOnBoard []types.Coord) types.Trajectories {
	t := make(types.Trajectories)
	for _, hit := range pointsOnBoard {
		t[hit] = make(types.Trajectory)
		for direction := range types.PossibleMoves {
			traceTrajectory(b, hit, t, direction)
		}
	}
	return t
}

func traceTrajectory(b types.Board, hit types.Coord, t types.Trajectories, direction types.Direction) {
	move := types.PossibleMoves[direction]
	n := hit
	for {
		mv := n.Add(move)
		if _, ok := b.Moves[mv.String()]; !ok {
			break
		}
		next := b.Moves[mv.String()]
		if next.State == types.HIT {
			t[hit][direction] = append(t[hit][direction], mv)
		}
		n = mv
	}
}

func calcMaxTrajectoryLength(boardSize int, remainingShips types.Fleet) int {
	var remainingShipsLength int
	for shipType, flotilla := range remainingShips {
		remainingShipsLength += int(shipType) * flotilla.Num
	}
	if remainingShipsLength < boardSize {
		return remainingShipsLength
	}
	return boardSize
}
