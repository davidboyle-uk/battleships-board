package ai

import (
	"sort"

	"battleships-board/types"
)

/*
// if no hits
	// determine coord with largest surrounding area (based on misses)

// if hits
	// loop over hits determine probability
		// determine possible hits in surrounding coords
			// depends on remaining ships
			// ships could be touching
			// bounds of map prevent
*/

func CalculateMove(b types.Board) types.Coord {
	if !b.HasHits() {
		return calcMoveWhenNoHits(b)
	}
	return calcMoveBasedOnProbability(b)
}

func calcMoveBasedOnProbability(b types.Board) types.Coord {
	hitsOnBoard := b.GetHits()
	sunkOnBoard := b.GetSunk()
	// get the trajectories for the current hits
	hitTrajectories := traceTrajectories(b, hitsOnBoard)
	remainingShips := calcRemainingShips(b, sunkOnBoard)
	// calculate by extending existing trajectores in the relevant direction
	allProbabilities := calcProbabilities(b, hitTrajectories, remainingShips)
	probabilities := combineProbabilities(allProbabilities)
	if len(probabilities) > 0 {
		mostProbable := mostProbable(probabilities)
		return pickCoordFromSlice(mostProbable)
	}

	return calcMoveWhenNoHits(b)
}

func combineProbabilities(allProbabilities types.Probabilities) types.Probabilities {
	var combinedProbabilities = make(types.Probabilities)

	for sourceBaseScore, sourceBaseCoords := range allProbabilities {
		for _, c := range sourceBaseCoords {
			totalScore := sourceBaseScore
			for targetBaseScore, targetBaseCoords := range allProbabilities {
				for _, t := range targetBaseCoords {
					if c == t {
						totalScore += targetBaseScore
					}
				}
			}
			if !contains(combinedProbabilities[totalScore], c) {
				combinedProbabilities[totalScore] = append(combinedProbabilities[totalScore], c)
			}
		}
	}

	return combinedProbabilities
}

func contains(h []types.Coord, t types.Coord) bool {
	for _, v := range h {
		if t == v {
			return true
		}
	}
	return false
}

func calcProbabilities(b types.Board, hitTrajectories types.Trajectories, remainingShips types.Fleet) types.Probabilities {
	var probabilities = make(types.Probabilities)

	for origin, trajectories := range hitTrajectories {
		switch len(trajectories) {
		case 0:
			for direction, move := range types.PossibleMoves {
				next := origin.Add(move)
				if next.IsOutOfBounds(b.Dim) {
					continue
				}
				trajectory := []types.Coord{origin}
				if _, ok := b.Moves[next]; !ok {
					probability := calcProbability(b.Dim, origin, direction, trajectory, remainingShips)
					probabilities[probability] = append(probabilities[probability], next)
					continue
				}
			}
		default:
			for direction, trajectory := range trajectories {
				move := types.PossibleMoves[direction]
				last := trajectory[len(trajectory)-1]
				next := last.Add(move)
				if next.IsOutOfBounds(b.Dim) {
					continue
				}
				if _, ok := b.Moves[next]; !ok {
					probability := calcProbability(b.Dim, origin, direction, trajectory, remainingShips)
					probabilities[probability] = append(probabilities[probability], next)
				}
			}
		}
	}

	return probabilities
}

func mostProbable(probabilities types.Probabilities) []types.Coord {
	keys := make([]int, 0, len(probabilities))
	for k := range probabilities {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return probabilities[keys[len(keys)-1]]
}

func calcProbability(boardSize int, origin types.Coord, direction types.Direction, trajectory []types.Coord, remainingShips types.Fleet) int {
	var probability int

	// for each remaining ship type
	for _, fortilla := range remainingShips {
		// is the trajectory plus the length of the ship minus the existing length of the trajectory still in bounds
		endPoint := origin.AddNum(types.PossibleMoves[direction], int(fortilla.ShipType))
		if !endPoint.IsOutOfBounds(boardSize) {
			probability++
		}
		// is the trajectory plus the move less than the length of the ship
		if len(trajectory)+1 < int(fortilla.ShipType) {
			probability++
		}
	}

	return probability
}

func calcRemainingShips(b types.Board, sunkOnBoard []types.Coord) types.Fleet {
	var remainingShips = types.NewFleet()

	sunkTrajectories := traceTrajectories(b, sunkOnBoard)

	for _, sunkShip := range sunkTrajectories {
		sunkShipType := types.ShipType(len(sunkShip) + 1)
		for _, fortilla := range remainingShips {
			if sunkShipType == fortilla.ShipType {
				remainingShips.RemoveShip(sunkShipType)
			}
		}
	}

	return remainingShips
}

func calcMoveWhenNoHits(b types.Board) types.Coord {
	coordSpaces := calcCoordSpaces(b)
	scores := make(map[int][]types.CoordSpace)
	for _, cs := range coordSpaces {
		span := (cs.MaxX - cs.MinX) * (cs.MaxY - cs.MinY)
		scores[span] = append(scores[span], cs)
	}
	largestSpans := largestSpan(scores)
	return closestToMiddle(largestSpans)
}

func closestToMiddle(set []types.CoordSpace) types.Coord {
	midPoint := calcMidpoint(set[0])
	var closesPoints = make(map[float64][]types.Coord)
	for _, sp := range set {
		d := sp.Coord.Dist(midPoint)
		closesPoints[d] = append(closesPoints[d], sp.Coord)
	}
	finalSet := minDist(closesPoints)

	return pickCoordFromSlice(finalSet)
}

func calcMidpoint(space types.CoordSpace) types.Coord {
	return types.Coord{
		X: (space.MaxX - space.MinX) / 2,
		Y: (space.MaxY - space.MinY) / 2,
	}
}

func pickCoordFromSlice(set []types.Coord) types.Coord {
	l := len(set)
	if l == 1 {
		return set[0]
	}
	return set[types.RandNum(l-1)]
}

func largestSpan(scores map[int][]types.CoordSpace) []types.CoordSpace {
	keys := make([]int, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return scores[keys[len(keys)-1]]
}

func minDist(dists map[float64][]types.Coord) []types.Coord {
	keys := make([]float64, 0, len(dists))
	for k := range dists {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	return dists[keys[0]]
}

func calcCoordSpaces(b types.Board) []types.CoordSpace {
	coordSpaces := []types.CoordSpace{}
	for x := 0; x <= b.Dim-1; x++ {
		for y := 0; y <= b.Dim-1; y++ {
			c := types.Coord{
				X: x,
				Y: y,
			}
			if _, ok := b.Moves[c]; !ok {
				coordSpaces = append(coordSpaces, calcCoordSpace(c, b))
			}
		}
	}
	return coordSpaces
}

func calcCoordSpace(c types.Coord, b types.Board) types.CoordSpace {
	cs := types.CoordSpace{
		Coord: c,
		MinX:  c.X,
		MaxX:  c.X,
		MinY:  c.Y,
		MaxY:  c.Y,
	}
	for _, move := range types.PossibleMoves {
		t := c
		for {
			next := t.Add(move)
			if next.IsOutOfBounds(b.Dim) {
				break
			}
			if _, ok := b.Moves[next]; ok {
				break
			}
			if next.X > cs.MaxX {
				cs.MaxX = next.X
			}
			if next.X < cs.MinX {
				cs.MinX = next.X
			}
			if next.Y > cs.MaxY {
				cs.MaxY = next.Y
			}
			if next.Y < cs.MinY {
				cs.MinY = next.Y
			}
			t = next
		}
	}
	return cs
}
